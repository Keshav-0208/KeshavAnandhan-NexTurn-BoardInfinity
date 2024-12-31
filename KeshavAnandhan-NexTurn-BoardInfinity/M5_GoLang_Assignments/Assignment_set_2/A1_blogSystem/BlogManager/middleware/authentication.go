package middleware

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			fmt.Println("Invalid or missing Authorization header")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			fmt.Println("Error decoding Authorization header:", err)
			c.JSON(401, gin.H{"error": "Invalid Authorization Header"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			fmt.Println("Invalid credentials format")
			c.JSON(401, gin.H{"error": "Invalid Credentials"})
			c.Abort()
			return
		}

		username, password := credentials[0], credentials[1]

		var storedPassword string
		err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
		if err != nil || storedPassword != password {
			fmt.Println("Authentication failed")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		fmt.Println("Authentication successful")
		c.Next()
	}
}
