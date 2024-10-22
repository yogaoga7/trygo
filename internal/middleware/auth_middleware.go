package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware mengecek apakah request memiliki header autentikasi yang valid
func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "valid-token" { // Misalnya cek token valid
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort() // Hentikan eksekusi dan balas respon
		return
	}

	c.Next() // Lanjutkan ke handler berikutnya jika valid
}
