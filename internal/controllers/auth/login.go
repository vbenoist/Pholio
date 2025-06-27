package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/vbenoist/pholio/internal/helpers/cfg"
	apimodels "github.com/vbenoist/pholio/internal/models/api"
	"github.com/vbenoist/pholio/internal/services/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var authInput apimodels.AuthInput
	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error::auth": err.Error()})
		return
	}

	registeredAdmin, err := auth.GetAdminFromUsername(authInput.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusBadRequest, gin.H{"error::auth": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error::auth": "Unhandled error on database communication"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(registeredAdmin.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  registeredAdmin.Id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.SetSameSite(http.SameSiteStrictMode)

	config := cfg.GetServerConfig()
	secureCookie := config.Env.Production
	c.SetCookie("auth_token", token, 3600, "/", "localhost", secureCookie, true)

	c.JSON(200, gin.H{
		"token": token,
	})
}
