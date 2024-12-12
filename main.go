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

	router.POST("/cart", controllers.AddToCart)
	router.GET("/cart", controllers.GetCartItems)
	router.PUT("/cart", controllers.UpdateCartItem)
	router.DELETE("/cart", controllers.RemoveFromCart)

	router.Run(":8080")
}
