package controllers

import (
	"database/sql"
	"gosightapi/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context, db *sql.DB) {
	// Implement logic to retrieve items from the database
	items, err := GetItemsFromDB(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

func UpdateItem(c *gin.Context, db *sql.DB) {
	// Parse o JSON recebido para o struct
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Atualize o banco de dados com base nos campos recebidos
	_, err := db.Exec(
		"UPDATE itens SET itm_descricao = ? WHERE itm_codigo = ?",
		item.Nome,   // Nome corresponde a "descricao" no JSON
		item.Codigo, // Codigo corresponde a "codigo" no JSON
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully"})
}

func GetItemsFromDB(db *sql.DB) ([]models.Item, error) {
	rows, err := db.Query("SELECT itm_codigo, itm_descricao FROM itens limit 100")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Codigo, &item.Nome); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
