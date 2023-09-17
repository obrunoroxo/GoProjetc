package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a chave secreta para assinar o token. Mantenha isso em um local seguro.
var secretKey = []byte("sua_chave_secreta_aqui")

func main() {
	// Crie um dicionário (mapa) com informações personalizadas
	dict := map[string]interface{}{
		"user_id":  123,
		"username": "exemplo_usuario",
	}

	// Crie um token com as informações do dicionário
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Preencha as reivindicações (claims) do token com os dados do dicionário
	for key, value := range dict {
		claims[key] = value
	}

	// Defina o tempo de expiração do token (por exemplo, 24 horas a partir de agora)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Assine o token com a chave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Erro ao assinar o token:", err)
		return
	}

	fmt.Println("Token JWT gerado:")
	fmt.Println(tokenString)
}
