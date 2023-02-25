package main

import (
	controller "go-product/controllers"
	"go-product/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/api/product", controller.CreateProduct)
	r.GET("/api/product", controller.ReadProducts)
	r.GET("/api/product/:id", controller.ReadProductByID)
	r.PUT("/api/product/:id", controller.UpdateProduct)
	r.DELETE("/api/product/:id", controller.RemoveProduct)

	r.GET("/api/product/filter", controller.MerchantFilter)
	r.GET("/api/product/page/:pageNum", controller.ReadProductByPage)

	r.Run()

}
