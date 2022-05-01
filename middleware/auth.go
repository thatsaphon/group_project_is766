package middleware

import (
	"context"
	"fmt"
	"group-project/firebase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware : to verify all authorized operations
func AuthMiddleware(c *gin.Context) {
	firebaseAuth := firebase.ClientAuth
	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
		c.Abort()
		return
	}
	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
	c.Set("UUID", token.UID)
	email := fmt.Sprintf("%s", token.Firebase.Identities["email"])
	c.Set("email", email[1:len(email)-1])
	c.Next()
}
