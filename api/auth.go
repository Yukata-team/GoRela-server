package api

import (
	"fmt"
	"net/http"
	"time"

	"../model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

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

	fmt.Println(user)

	//railsでいう presence：true
	if user.Email == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	model.CreateUser(user)
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	//パスワードが一致するかどうか
	user := model.FindUser(&model.User{ID: u.ID})
	if user.ID == 0 || user.Password != u.Password {
		return &echo.HTTPError {
			Code: http.StatusUnauthorized,
			Message: "invalid name or password",
		}
	}

	//多分jwtを適用してる
	claims := &jwtCustomClaims{
		user.ID,
		user.Email,
		jwt.StandardClaims {
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
