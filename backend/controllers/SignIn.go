package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"

	"github.com/KUNSALISA/Meeting-Room-Booking-System/config"
	"github.com/KUNSALISA/Meeting-Room-Booking-System/entity"
	"github.com/KUNSALISA/Meeting-Room-Booking-System/services"
)

type (
	Authen struct {
		Codename string
		Password string
	}

	UpdatePassword struct {
		CodeName    string `binding:"required"`
		NewPassword string `binding:"required"`
	}
)

func SignInUser(c *gin.Context) {
	var payload Authen
	var user entity.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB().Preload("Role").
		Where("code_name = ?", payload.Codename).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่พบ Codename นี้"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านไม่ถูกต้อง"})
		return
	}

	jwtWrapper := services.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.CodeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้าง token ได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token_type":   "Bearer",
		"token":        signedToken,
		"user_id":      user.ID,
		"codename":     user.CodeName,
		"first_name":   user.Firstname,
		"last_name":    user.Lastname,
		"image":        user.Image,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"role":         user.Role.RoleName,
	})
}

func ChangePassword(c *gin.Context) {
	var req UpdatePassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user entity.User
	if err := config.DB().Where("code_name = ?", req.CodeName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	hashedPassword, err := config.HashPassword(req.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = hashedPassword

	if err := config.DB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func SignUpUser(c *gin.Context) {
	var payload entity.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userRole entity.Role
	if err := config.DB().Where("role_name = ?", "User").First(&userRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถกำหนด Role เป็น User ได้"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเข้ารหัสรหัสผ่านได้"})
		return
	}
	payload.Password = string(hashedPassword)

	payload.RoleID = userRole.ID

	if err := config.DB().Create(&payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถเพิ่มผู้ใช้ใหม่ได้"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "สมัครสมาชิกสำเร็จ",
		"user_id":  payload.ID,
		"codename": payload.CodeName,
		"email":    payload.Email,
	})
}
