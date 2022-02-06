package main

import (
	"fmt"
	"log"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
	dotenv "github.com/joho/godotenv"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Homepage")
	fmt.Println("User accessed Homepage")
}

func isValid(w http.ResponseWriter, r *http.Request){
	if r.Header["Token"] != nil {
		// Esse interface é um objeto (JSON-like) pode ser usado para melhor compreensão do erro
		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error){
			
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// token = nil, erro = erro
				return nil, fmt.Errorf("Error while trying signing method")
			}

			// token = secret, erro = nil
			return jwtsecret, nil
		})

	} else {
		fmt.Fprintf(w, "Access not authorized")
	}
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	// Retorna uma avaliação
	return http.HandlerFunc(isValid)
}

// Equivalente ao Routes no Node
func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var jwtsecret string
func main(){
	envs, err := dotenv.Read(".env")
	if err != nil {
		log.Fatal("Error while finding .env file")
	}

	jwtsecret := envs["JWTSECRET"]
	handleRequests()
}
