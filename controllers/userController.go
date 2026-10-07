package controllers

import (
	"context"
	"fmt"
	"go-jwt-auth/helpers"
	helper "go-jwt-auth/helpers"
	"go-jwt-auth/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/karanrejuuu/go-jwt-auth/database"
	"github.com/karanrejuuu/go-jwt-auth/models"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(databse.Client, "user")
var validate = validator.New()

func HashPassword()

func VerifyPassword()

func Signup() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user): err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return

		}	
		validationErr := validate.Struct(user)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for email"})
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"phone":user.Phone})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the email"})
		}

		if count>0{
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email or phone number already exists"})
		}
	}
}

func Login()

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId): err!= nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user modles.user
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOk, user)
	}
}
