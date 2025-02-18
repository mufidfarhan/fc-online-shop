package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := os.Getenv("ADMIN_SECRET")

		// TODO: ambil header authorization
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{"error": "Akses tidak diizinkan"})
			c.Abort() // **untuk menghentikan proses request
			return
		}

		// TODO: validasi header sesuai dengan kata sandi admin
		if auth != key {
			c.JSON(401, gin.H{"error": "Akses tidak diizinkan"})
			c.Abort()
			return
		}

		// TODO: lanjutkan request ke header
		c.Next()
	}
}
