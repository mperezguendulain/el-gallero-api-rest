package jwtserv

import (
	"io/ioutil"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
)

// Credential store the login information
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SecretKey is a secret key for JWT
var SecretKey []byte

// Inicializa el módulo
func init() {

	var err error
	SecretKey, err = ioutil.ReadFile("./keys/public.rsa.pub")
	if err != nil {
		log.Fatal("[RSA] The file with the public key could not be read.")
	}

}

// GetJWT get a JSON Web Token
func GetJWT(username string, role string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["role"] = "admin"

	// Generate encoded token and send it as response.
	stringToken, err := token.SignedString(SecretKey)
	if err != nil {
		log.Println("[JWT] Error when signing token: ", err)
		return "", err
	}

	return stringToken, nil

}
