package middlewares

import (
	"strconv"
	"go-jwt/database"
	"go-jwt/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

func ProductAuthorization() gin.HandlerFunc{
	return func(c *gin.Context){
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil{
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "invalid parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Product.UserID != userID{
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}