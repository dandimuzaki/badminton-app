package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/dandimuzaki/badminton-server/initializers"
	"github.com/dandimuzaki/badminton-server/model"
	"github.com/dandimuzaki/badminton-server/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Hide password before sending response
	safeUser := struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    safeUser,
	})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := initializers.DB.First(&user, "email = ?", input.Email).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, _ := utils.GenerateToken(user.ID)

	safeUser := struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	isProd := os.Getenv("ENV") == "production"
	secure := isProd
	serverHost := os.Getenv("SERVER_HOST")

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "badminton-app-production.up.railway.app",
		MaxAge:   3600 * 24,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
  }

  http.SetCookie(c.Writer, &cookie)
	
	log.Printf(serverHost, secure)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": safeUser, "token": token})
}
