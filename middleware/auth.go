package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. रिक्वेस्ट के Headers से Authorization की वैल्यू निकालना
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "टोकन गायब है! कृपया लॉगिन करें।"})
			c.Abort() // रिक्वेस्ट को आगे बढ़ने से यहीं रोक दें
			return
		}

		// 2. चेक करना कि क्या टोकन "Bearer <token>" फॉर्मेट में है (इंडस्ट्री स्टैंडर्ड)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "गलत टोकन फॉर्मेट (Bearer Token आवश्यक है)"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. JWT टोकन को वैलिडेट और डिकोड (Parse) करना
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// पक्का करना कि साइनिंग मेथड सही है
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// 4. अगर टोकन एक्सपायर हो गया है या अमान्य है
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "अमान्य या एक्सपायर्ड टोकन!"})
			c.Abort()
			return
		}

		// 5. टोकन से यूजर की ID (sub) निकालना और उसे रिक्वेस्ट कॉन्टेक्स्ट में सेव करना
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["sub"]
			c.Set("userID", userID) // इसे आगे कंट्रोलर्स में यूज़ किया जा सकता है
		}

		c.Next() // सब कुछ सही है, अब रिक्वेस्ट को आगे कंट्रोलर के पास जाने दें
	}
}
