package interfaces

import (
	"ibp/config"
	"fmt"
	"github.com/labstack/echo"
	"ibp/domain"
	"log"
	"net/http"
)

// Run start server
func Run(e *echo.Echo, port string) {
	log.Printf("Server running at http://localhost:%s/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

// Routes returns the initialized router
func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/api/v1/article", func(context echo.Context) error {
		param := new(domain.Article)
		printDebugf("test")
		if err := context.Bind(param); err != nil {
			log.Println("bad request")
			return context.String(http.StatusBadRequest, "bad request")
		}
		return context.String(http.StatusCreated, "")
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