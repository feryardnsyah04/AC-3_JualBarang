package main

import (
	"shopping-cart-api/controllers"
	"shopping-cart-api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "shopping-cart-api Koneksi Berhasil",
		})
	})

	router.GET("/cart", controllers.GetCartItems)
	router.POST("/cart", controllers.AddCartItem)
	router.PUT("/cart/:id", controllers.UpdateCartItemQuantity)
	router.DELETE("/cart/:id", controllers.DeleteCartItem)

	router.Run(":8080")
}
