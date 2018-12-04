build: 
	@go get github.com/labstack/echo
	@go get github.com/dgrijalva/jwt-go

run: 
	@go run main/main.go