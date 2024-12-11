package controllers

import (
	"shopping-cart-api/database"
	"shopping-cart-api/models"
	"shopping-cart-api/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

type DeleteRequest struct {
	ID int `json:"id"`
}

type UpdateQuantityRequest struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

func GetCartItems(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, product, category, price, quantity FROM cart_items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.Product, &item.Category, &item.Price, &item.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

func AddCartItem(c *gin.Context) {
	var item models.CartItem
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO cart_items (product, category, price, quantity) VALUES (?, ?, ?, ?)"
	result, err := database.DB.Exec(query, item.Product, item.Category, item.Price, item.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	item.ID = int(id)
	c.JSON(http.StatusCreated, item)
}

func DeleteCartItem(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format permintaan tidak valid"})
		return
	}

	query := "DELETE FROM cart_items WHERE id = ?"
	_, err := database.DB.Exec(query, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM cart_items").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if count == 0 {
		_, err := database.DB.Exec("ALTER TABLE cart_items AUTO_INCREMENT = 1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}

func UpdateCartItemQuantity(c *gin.Context) {
	var req UpdateQuantityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format permintaan tidak valid"})
		return
	}

	query := "UPDATE cart_items SET quantity = ? WHERE id = ?"
	_, err := database.DB.Exec(query, req.Quantity, req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quantity produk berhasil diupdate"})
}

func Checkout(c *gin.Context) {
	var checkoutQueue utils.Queue
    for !checkoutQueue.IsEmpty() {
      item, _ := checkoutQueue.Dequeue()
      fmt.Printf("Memproses item: %v\n", item)
    }
    c.JSON(200, gin.H{"message": "Checkout Berhasil"})
}
