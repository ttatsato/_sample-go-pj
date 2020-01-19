package interfaces

import (
	"app/config"
	"app/domain"
	"app/usecase"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

// Run start server
func Run(e *echo.Echo, port string) {
	log.Printf("Server running at http://localhost:%s/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}
	return nil
}

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	log.Printf("Request Body: %v\n", string(reqBody))
	log.Printf("Response Body: %v\n", string(resBody))
}

// Routes returns the initialized router
func Routes(e *echo.Echo) {
	e.Use(middleware.BodyDump(bodyDumpHandler))
	e.Validator = &Validator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/v1/article", func(c echo.Context) error {
		articles, err := usecase.FetchAllArticles()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, convertMapToJsonString(articles))
	})

	e.POST("/api/v1/article", func(c echo.Context) error {
		article := new(domain.Article)
		if err := BindValidate(c, article); err != nil {
			return err
		}
		if err := usecase.CreateNewArticle(article); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusCreated, "OK")
	})

	e.PATCH("/api/v1/article", func(c echo.Context) error {
		article := new(domain.Article)
		if err := BindValidate(c, article); err != nil {
			return err
		}
		if err := usecase.UpdateArticle(article); err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "OK")
	})

	e.DELETE("/api/v1/article/:articleId", func(c echo.Context) error {
		articleId, err := strconv.Atoi(c.Param("articleId"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if err := usecase.RemoveArticle(articleId); err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.String(http.StatusNoContent, "OK")
	})

	// Migration Route
	e.GET("/api/v1/migrate", migrate)
}

// =============================
//    MIGRATE
// =============================
func migrate(c echo.Context) error {
	_, err := config.DBMigrate()
	if err != nil {
		return c.String(http.StatusNotFound, "failed migrate")
	} else {
		return c.String(http.StatusOK, "success migrate")
	}
}
