package main

import (
	// "fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {

	// Set custom claims
		claims := &JwtCustomClaims{
			"hongxuan",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		cookie := new(http.Cookie)
		cookie.Name = "_token_jwt"
		cookie.Value = t
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)


		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
} 
func testlogin(c echo.Context) error {
	// Set custom claims
	claims := &JwtCustomClaims{
		"hongxuan",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "_token_jwt"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)


	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func check(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	name := claims.Name
	if name == "hongxuan" {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ok",
		})
	}

	return c.JSON(http.StatusBadRequest, echo.Map{
			"status": "error",
		})
	
}


func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  }))
	// Login route
	e.POST("/login", login)
	e.GET("/testlogin", testlogin)
	// Unauthenticated route
	// e.GET("/", accessible)

	// Restricted group
	adm := e.Group("/admin")
	// obj := e.Group("/obj")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		TokenLookup: "cookie:_token_jwt",
		SigningKey: []byte("secret"),
	}

	adm.Use(middleware.JWTWithConfig(config))
	// adm.GET("", admin.Restricted())
	adm.GET("/hello", check)



	// obj.GET("/about/:lang", about.GetAbout())

	e.Logger.Fatal(e.Start(":5000"))
}