package utils

// import (
// 	"fmt"

// 	"github.com/golang-jwt/jwt/v5"
// )

// type SessionClaims struct {
//     UserID   string `json:"userID"`
//     UserPhoto string `json:"userPhoto"`
//     jwt.StandardClaims
// }

// // DecodeSessionToken decodifica un token de sesión y devuelve los datos del usuario
// func DecodeSessionToken(tokenString string) (string, string, string, error) {
// 	// Inicializar un nuevo objeto Claims para almacenar los datos decodificados
// 	jwt.Claims = &Claims{}
// 	claims := jwt.Claims.(*Claims)

// 	// Utilizar la función ParseWithClaims para validar y decodificar el token
// 	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Verificar si el token es válido
// 	if !token.Valid {
// 		return "", "", fmt.Errorf("Token no válido")
// 	}

// 	// Extraer los datos del usuario del token decodificado
// 	userID := claims.UserID
// 	userPhoto := claims.UserPhoto
// 	username := claims.UserName

// 	return userID, userPhoto, username, nil
// }
