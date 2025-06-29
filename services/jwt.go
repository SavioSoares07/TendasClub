package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)
	
	var secretKey []byte

func init() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
    log.Fatal("Erro ao carregar o arquivo .env")
}
	// Carrega a variável TOKEN_SECRET do arquivo .env
	godotenv.Load()
	tokenSecret := os.Getenv("TOKEN_SECRET")


	secretKey = []byte(tokenSecret)
}

// CreateToken cria um token JWT para o usuário com base no email
func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
// VerifyToken verifica se o token JWT é válido
func VerifyToken(tokenString string) error{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil	
	})

	if err != nil {
      return err
   }
  
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}
