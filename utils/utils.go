package utils

import (
    "errors"
    "github.com/dgrijalva/jwt-go"
)

// VerifyAuthToken verifies the JWT token and returns the user ID and role
func VerifyAuthToken(tokenString string) (string, string, error) {
    // Your JWT secret key
    secretKey := []byte("your-secret-key")
    
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil || !token.Valid {
        return "", "", errors.New("invalid or expired token")
    }

    // Extract the user ID and role from the token's claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", "", errors.New("invalid token claims")
    }

    userID, ok := claims["user_id"].(string)
    if !ok {
        return "", "", errors.New("user_id not found in token")
    }

    role, ok := claims["role"].(string)
    if !ok {
        return "", "", errors.New("role not found in token")
    }

    return userID, role, nil
}
