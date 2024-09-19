package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	BasePrice float64 `db:"base_price"`
}

type RequestBody struct {
	ID       int    `json:"id"`
	Discount string `json:"discount"`
	Country  string `json:"country"`
}

var db *sqlx.DB

func initDB() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}

func calculatePrice(basePrice float64, discount string, taxRate float64) float64 {
	finalPrice := basePrice

	if discount != "" {
		if discount[0] == 'p' {
			percentage, err := strconv.Atoi(discount[1:])
			if err == nil {
				finalPrice = finalPrice * (1 - float64(percentage)/100)
			}
		} else {
			fixedDiscount, err := strconv.ParseFloat(discount, 64)
			if err == nil {
				finalPrice = finalPrice - fixedDiscount
			}
		}
	}

	finalPrice = finalPrice * (1 + taxRate)
	if finalPrice < basePrice*0.5 {
		finalPrice = basePrice * 0.5
	}
	return finalPrice
}

func getProductPrice(c *gin.Context) {
	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product Product
	err := db.Get(&product, "SELECT id, name, base_price FROM products WHERE id=$1", requestBody.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	taxRate := 0.0 // This should be replaced by your actual tax fetching logic based on country.
	finalPrice := calculatePrice(product.BasePrice, requestBody.Discount, taxRate)

	c.JSON(http.StatusOK, gin.H{
		"id":    product.ID,
		"name":  product.Name,
		"price": finalPrice,
	})
}

func addProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO products (name, base_price) VALUES ($1, $2)", product.Name, product.BasePrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product added successfully"})
}

func main() {
	initDB()
	router := gin.Default()
	router.POST("/get-price", getProductPrice)
	router.POST("/add-product", addProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
