package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sanket9162/hotel-reservation/db"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-Api-Token")
		if token == "" {
			return fmt.Errorf("unauthorized")
		}
		fmt.Println(token)
		claims, err := validateToke(token)
		if err != nil {
			return err
		}
		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		if time.Now().Unix() > expires {
			return fmt.Errorf("token expired")
		}
		userID := claims["id"].(string)
		user, err := userStore.GetUserById(c.Context(), userID)
		if err != nil {
			return fmt.Errorf("unaithorized")
		}
		c.Context().SetUserValue("user", user)
		return c.Next()
	}
}

func validateToke(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid string method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse JWT token:", err)
		return nil, fmt.Errorf("unauthorized")
	}
	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")

	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}
	return claims, nil
}
