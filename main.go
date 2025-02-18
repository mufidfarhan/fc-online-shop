package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"tokoonline/handler"
	"tokoonline/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DB_URI"))
	if err != nil {
		fmt.Printf("Gagal membuat koneksi database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Gagal memverifikasi koneksi database: %v\n", err)
		os.Exit(1)
	}

	if _, err = migrate(db); err != nil {
		fmt.Printf("Gagal melakukan migrasi database: %v\n", err)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/api/v1/products", handler.ListProducts(db))
	r.GET("/api/v1/products/:id", handler.GetProduct(db))
	r.POST("/api/v1/checkout", handler.CheckoutOrder(db))

	r.POST("/api/v1/orders/:id/confirm", handler.ConfirmOrder(db))
	r.GET("/api/v1/orders/:id", handler.GetOrder(db))

	r.POST("/admin/products", middleware.AdminOnly(), handler.CreateProduct(db))
	r.PUT("/admin/products/:id", middleware.AdminOnly(), handler.UpdateProduct(db))
	r.DELETE("/admin/products/:id", middleware.AdminOnly(), handler.DeleteProduct(db))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if err = server.ListenAndServe(); err != nil {
		fmt.Printf("Gagal menjalankan server: %v\n", err)
		os.Exit(1)
	}
}
