package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ruiborda/go-jwt/src/application/ports/input"
	"github.com/ruiborda/go-jwt/src/domain/entity"
	inputadapter "github.com/ruiborda/go-jwt/src/infrastructure/adapters/input"
	"github.com/ruiborda/go-service-common/dto"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

func RequireJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}
		token := tokenParts[1]

		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			slog.Error("JWT_SECRET environment variable is not set")
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		inputPort := input.NewJWTHS256InputPort[*dto.JwtPrivateClaims]([]byte(jwtSecret))
		jwtInputAdapter := inputadapter.NewJwtInputAdapter[*dto.JwtPrivateClaims](inputPort)

		err := jwtInputAdapter.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		jwt := entity.NewJwtFromToken[*dto.JwtPrivateClaims](token)
		if jwt == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		c.Set("jwtClaims", jwt.Claims)
		c.Next()
	}
}

func RequirePermission(permissionId int) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsValue, exists := c.Get("jwtClaims")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No JWT claims found"})
			return
		}

		claims, ok := claimsValue.(*entity.JWTClaims[*dto.JwtPrivateClaims])
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid JWT claims format"})
			return
		}

		hasPermission := false
		if claims.PrivateClaims != nil && claims.PrivateClaims.PermissionIds != nil {
			for _, id := range claims.PrivateClaims.PermissionIds {
				if id == permissionId {
					hasPermission = true
					break
				}
			}
		}

		if !hasPermission {
			slog.Info("Access denied: missing required permission", "requiredPermission", permissionId, "email", claims.PrivateClaims.Email)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			return
		}

		c.Next()
	}
}
