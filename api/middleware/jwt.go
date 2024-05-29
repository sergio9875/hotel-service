package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func JWTAuthentication(c *fiber.Ctx) error {
	fmt.Println("auth_____")
	_, ok := c.GetReqHeaders()["X-Api-Token"]

	if !ok {
		return fmt.Errorf("unauthorized")
	}
	//if err := parseToken(token); err != nil {
	//	return err
	//}
	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("NEVER DO THIS:::", secret)
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("Failed to parse JWT token:", err)
		return fmt.Errorf("unauthorized")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	}
	return fmt.Errorf("unauthorized")
}
