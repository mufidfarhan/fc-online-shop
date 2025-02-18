package handler

import (
	"database/sql"
	"errors"
	"log"
	"tokoonline/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListProducts(db *sql.DB) gin.HandlerFunc { // **teknik dependency injection
	return func(c *gin.Context) {
		// TODO: ambil dari database
		products, err := model.SelectProduct(db)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// TODO: berikan response
		c.JSON(200, products)
		// **untuk di akhir fungsi handler bisa tanpa menuliskan return
	}
}

func GetProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: baca id dari url
		id := c.Param("id")

		// TODO: ambil dari database dengan id
		product, err := model.SelectProductByID(db, id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) { // **penambahan handler error untuk db yang kosong sehingga tidak mengembalikan error 500
				log.Printf("Terjadi kesalahan saat mengambil data produk: %v\n", err)
				c.JSON(404, gin.H{"error": "Produk tidak ditemukan"})
				return
			}

			log.Printf("Terjadi kesalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		// TODO: berikan response
		c.JSON(200, product)
	}
}

func CreateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product model.Product
		if err := c.Bind(&product); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data produk tidak valid"})
			return
		}

		product.ID = uuid.New().String()

		if err := model.InsertProduct(db, product); err != nil {
			log.Printf("Terjadi kesalahan saat menambahkan produk: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		c.JSON(201, product)
	}
}

func UpdateProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var productReq model.Product
		if err := c.Bind(&productReq); err != nil {
			log.Printf("Terjadi kesalahan saat membaca request body: %v\n", err)
			c.JSON(400, gin.H{"error": "Data produk tidak valid"})
			return
		}

		product, err := model.SelectProductByID(db, id)
		if err != nil {
			log.Printf("Terjadi kesalahan saat mengambil data produk: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		if productReq.Name != "" {
			product.Name = productReq.Name
		}

		if productReq.Price != 0 {
			product.Price = productReq.Price
		}

		if err := model.UpdateProduct(db, product); err != nil {
			log.Printf("Terjadi kesalahan saat memperbarui data produk: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		c.JSON(201, product)
	}
}

func DeleteProduct(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := model.DeleteProduct(db, id); err != nil {
			log.Printf("Terjadi kesalahan saat menghapus data produk: %v\n", err)
			c.JSON(500, gin.H{"error": "Terjadi kesalahan pada server"})
			return
		}

		c.JSON(201, gin.H{"message": "Produk berhasil dihapus"})
	}
}
