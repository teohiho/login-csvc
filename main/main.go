package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/dgrijalva/jwt-go"
	"time"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"os"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

	
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	IDDonvi  string `json:"id_donvi"`
	IDRole   string `json:"id_role"`
}



func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	resp, _ := http.Get("http://localhost:5500/user")

	defer resp.Body.Close()
	
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
	}

	fmt.Printf("%s\n", string(contents))
	var data = []User{}
	_ = json.Unmarshal(contents, &data)
	


	for _, value := range data {
		fmt.Println("username: ", username)
		fmt.Println("username: ", value.Username)
		fmt.Println("username: ", password)
		fmt.Println("username: ", value.Password)
		fmt.Println("============================>")
		if username == value.Username && password == value.Password {
			fmt.Println("username: ", username)
			fmt.Println("username: ", value.Username)
			fmt.Println("username: ", password)
			fmt.Println("username: ", value.Password)
			claims := &JwtCustomClaims{
				username,
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
		AllowCredentials: true,
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