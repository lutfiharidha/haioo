package app

import (
	"github.com/gin-gonic/gin"
)

func CartsRoute(route *gin.Engine) {
	m := NewCartModels()
	m.Init()
	c := NewCartController()
	c.Init(m)
	cartRoutes := route.Group("api/v1/cart")
	{
		cartRoutes.POST("/add", c.AddProduct)
		cartRoutes.DELETE("/delete", c.DeleteProduct)
		cartRoutes.GET("/show", c.ShowCart)
	}
}

func Router() {
	r := gin.New()
	r.Use(CORS()) //handle cors
	CartsRoute(r) //call registered route

	r.Run(":8081")

}
