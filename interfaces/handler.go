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
func Run(e *echo.Echo, port int) {
	log.Printf("Server running at http://localhost:%d/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// Routes returns the initialized router
func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//e.GET("/users/:id", getUser)
	//e.GET("/shop/:shopId", func(context echo.Context) error {
	//	shopId, _ := strconv.Atoi(context.Param("shopId"))
	//	return context.String(http.StatusOK, convertMapToJsonString(usecase.FetchShopInfo(shopId)))
	//})
	//e.GET("/shop/order/:shopId", func(context echo.Context) error {
	//	shopId, _ := strconv.Atoi(context.Param("shopId"))
	//	return context.String(http.StatusOK, convertMapToJsonString(usecase.ListShopOrder(shopId)))
	//})
	//e.GET("/order/:customerId", func(context echo.Context) error {
	//	customerId, _ := strconv.Atoi(context.Param("customerId"))
	//	return context.String(http.StatusOK, convertMapToJsonString(usecase.ListCustomerOrderHistory(customerId)))
	//})
	e.POST("/api/v1/article", func(context echo.Context) error {
		param := new(domain.Article)
		printDebugf("test")
		if err := context.Bind(param); err != nil {
			log.Println("bad request")
			return context.String(http.StatusBadRequest, "bad request")
		}
		return context.String(http.StatusCreated, "")
	})
	//e.GET("/product/:shopId", func(context echo.Context) error {
	//	shopId, _ := strconv.Atoi(context.Param("shopId"))
	//	return context.String(http.StatusOK, convertMapToJsonString(usecase.ListShopProduct(shopId)))
	//})
	//e.POST("/product/new", usecase.CreateNewProduct)

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