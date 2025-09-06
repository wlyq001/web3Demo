package service

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BaseApi struct{}

var BaseApiApp = new(BaseApi)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserId  uint
	User    User `gorm:"foreignkey:UserId"`
}

type Comment struct {
	gorm.Model
	Content string
	UserId  uint
	PostId  uint
	Post    Post `gorm:"foreignkey:PostId"`
	User    User `gorm:"foreignkey:UserId"`
}

func (b *BaseApi) Login(c *gin.Context) {
	var user User
	err := c.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dbUser User
	DB.Model(&user).Where("username = ?", user.Username).First(&dbUser)
	if dbUser.ID == 0 {
		c.JSON(http.StatusOK, "用户名错误")
		return
	}
	err = comparePasswords(dbUser.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusOK, "密码错误")
		return
	}

	token, err := GenerateToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
	return

}

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(8 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (b *BaseApi) Register(c *gin.Context) {
	var user User
	err := c.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	password, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = password
	DB.Create(&user)
	c.JSON(http.StatusOK, "success")
}

func hashPassword(password string) (string, error) {
	// 使用bcrypt库的GenerateFromPassword函数进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword, inputPassword string) error {
	// 使用bcrypt库的CompareHashAndPassword函数比较密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err
}

func (b *BaseApi) JwtTest(c *gin.Context) {
	value, _ := c.Get("userID")
	c.JSON(200, value)
}
