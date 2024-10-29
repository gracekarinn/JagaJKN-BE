package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func validateBearerToken(c *gin.Context) (string, error) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        return "", fmt.Errorf("authorization header is required")
    }

    bearerToken := strings.Split(authHeader, " ")
    if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
        return "", fmt.Errorf("invalid authorization format")
    }

    return bearerToken[1], nil
}

func UserAuthMiddleware(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, err := validateBearerToken(c)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(jwtSecret), nil
        })

        if err != nil {
            log.Printf("Error parsing token: %v", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            if nik, exists := claims["nik"].(string); exists {
                c.Set("userNIK", nik)
                c.Next()
                return
            }
            log.Printf("NIK not found in claims: %+v", claims)
        }
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        c.Abort()
    }
}

func AdminAuthMiddleware(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, err := validateBearerToken(c)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(jwtSecret), nil
        })

        if err != nil {
            log.Printf("Error parsing token: %v", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            adminID, exists := claims["admin_id"].(float64)
            if !exists {
                log.Printf("Admin ID not found in claims: %+v", claims)
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid admin token"})
                c.Abort()
                return
            }
            c.Set("adminID", uint(adminID))
            if email, exists := claims["email"].(string); exists {
                c.Set("adminEmail", email)
            }
            c.Next()
            return
        }
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        c.Abort()
    }
}

func FaskesAuthMiddleware(jwtSecret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString, err := validateBearerToken(c)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(jwtSecret), nil
        })

        if err != nil {
            log.Printf("Error parsing token: %v", err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            if kodeFaskes, exists := claims["kode_faskes"].(string); exists {
                c.Set("kodeFaskes", kodeFaskes)
                c.Next()
                return
            }
            log.Printf("Kode Faskes not found in claims: %+v", claims)
        }
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
        c.Abort()
    }
}