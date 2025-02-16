package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize the database connection
func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./crud.db")
	if err != nil {
		log.Fatal(err)
	}
	// Create table if not exists
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

// Item structure
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create a new item
func createItem(c *gin.Context) {
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO items (name, description) VALUES (?, ?)`
	_, err := db.Exec(query, item.Name, item.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Item created"})
}

// Read all items
func getItems(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, description FROM items")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read item"})
			return
		}
		items = append(items, item)
	}

	c.JSON(http.StatusOK, items)
}

// Update an item by ID
func updateItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `UPDATE items SET name = ?, description = ? WHERE id = ?`
	_, err := db.Exec(query, item.Name, item.Description, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Item updated"})
}

// Delete an item by ID
func deleteItem(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM items WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Item deleted"})
}

func main() {
	// Initialize the database
	initDB()

	// Set up Gin router
	r := gin.Default()

	// Define CRUD routes
	r.POST("/items", createItem)
	r.GET("/items", getItems)
	r.PUT("/items/:id", updateItem)
	r.DELETE("/items/:id", deleteItem)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
