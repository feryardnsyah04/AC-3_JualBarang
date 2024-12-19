package controllers

import (
	"database/sql"
	"net/http"
	"fmt"
	"shopping-cart-api/database"
	"shopping-cart-api/models"
	"shopping-cart-api/utils"
	"github.com/gin-gonic/gin"
)

var cartStack utils.Stack

func AddToCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.BindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Masukan tidak valid!"})
		return
	}

	id, err := cartStack.Push(cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan item ke database!"})
		return
	}
	cartItem.ID = int(id)
	c.JSON(http.StatusCreated, gin.H{"message": "Barang ditambahkan ke keranjang", "item": cartItem})
}

func GetCartItems(c *gin.Context) {
	query := "SELECT id, product, variant, price, quantity FROM cart_items"
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil item dari database!"})
		return
	}
	defer rows.Close()

	var items []models.CartItem
	for rows.Next() {
		var item models.CartItem
		if err := rows.Scan(&item.ID, &item.Product, &item.Variant, &item.Price, &item.Quantity); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses item dari database!"})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

func UpdateCartItem(c *gin.Context) {
	var updateData models.CartItem
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format permintaan tidak valid!"})
		return
	}

	query := "SELECT id, product, variant, price, quantity FROM cart_items WHERE id = ?"
	row := database.DB.QueryRow(query, updateData.ID)
	var existingItem models.CartItem
	if err := row.Scan(&existingItem.ID, &existingItem.Product, &existingItem.Variant, &existingItem.Price, &existingItem.Quantity); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan di keranjang!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memeriksa item di database!"})
		return
	}

	updateQuery := "UPDATE cart_items SET variant = ?, quantity = ? WHERE id = ?"
	_, err := database.DB.Exec(updateQuery, updateData.Variant, updateData.Quantity, updateData.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate item di database!"})
		return
	}

	for i, item := range cartStack.GetAll() {
		if item.ID == updateData.ID {
			cartStack.GetAll()[i].Variant = updateData.Variant
			cartStack.GetAll()[i].Quantity = updateData.Quantity
			updateData = cartStack.GetAll()[i]
			break
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item berhasil diperbarui", "item": updateData})
}


func RemoveFromCart(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.BindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Masukan tidak valid!"})
		return
	}

	query := "SELECT id, product, variant, price, quantity FROM cart_items WHERE id = ?"
	row := database.DB.QueryRow(query, cartItem.ID)
	var existingItem models.CartItem
	if err := row.Scan(&existingItem.ID, &existingItem.Product, &existingItem.Variant, &existingItem.Price, &existingItem.Quantity); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item tidak ditemukan di database!"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memeriksa item di database!"})
		return
	}

	_, err := database.DB.Exec("DELETE FROM cart_items WHERE id = ?", cartItem.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus item dari database!"})
		return
	}

	if !cartStack.RemoveByID(cartItem.ID) {
		fmt.Println("Item tidak ditemukan di stack, tetapi dihapus dari database")
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item dihapus dari keranjang", "item": existingItem})
}
