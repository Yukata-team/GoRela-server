package handler

import (
	"fmt"
	"github.com/Yukata-team/GoRela-server/model"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims: &jwtCustomClaims{},
	SigningKey: signingKey,
}

func SignupPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"email":  "email",
			"password": "password",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	//パスワードのハッシュを生成
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("パスワード:", user.Password)
	fmt.Println("ハッシュ化されたパスワード", hash)

	user.Password = string(hash)
	fmt.Println("コンバート後のパスワード:", user.Password)

	//railsでいう presence：true
	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	model.CreateUser(user)

	//DBに登録できたらパスワードをからにしておく
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	//パスワードが一致するかどうか
	user := model.FindUser(&model.User{Email: u.Email})
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(u.Password))
	if user.ID == 0 || err != nil {
		return &echo.HTTPError {
			Code: http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}

	//claimを生成
	claims := &jwtCustomClaims{
		user.ID,
		user.Email,
		jwt.StandardClaims {
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	//tokenを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

//投稿時などに認証トークンからidを持ってくる
func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	id := claims.ID
	return id
}