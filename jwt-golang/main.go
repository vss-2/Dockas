package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	dotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global variables
var jwtsecret string
var apiport string
var mongodbURI string
var workdir string
var collection *mongo.Collection

// User datatype
// in mongo, create database and collection: 
//   use User;
//   db.User.user.insert({...})
type User struct {
	Name_first	string `bson:"text"`
	Name_last	string `bson:"text"`
	Email		string `bson:"text"`
	Password	string `bson:"text"`
	Token		string `bson:"text"`
}

func doNothing(w http.ResponseWriter, r *http.Request){}

func favicon(w http.ResponseWriter, r *http.Request){
	envs, err := dotenv.Read(".env")
	
	if err != nil {
		log.Fatal("Error while finding .env file")
	}

	// Sample source: https://favicon.cc/?action=icon&file_id=107462
	http.ServeFile(w, r, string(envs["WORKDIR"])+string("favicon.ico"))
}

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
				return nil, fmt.Errorf("Error while trying signing method", ok)
			}

			// token = secret, erro = nil
			return jwtsecret, nil
		})
		_, _ = token, err
	} else {
		fmt.Fprintf(w, "Access not authorized")
	}
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	// Retorna uma avaliação
	return http.HandlerFunc(isValid)
}

// Equivalente ao Routes no Node
func handleRequests(envs map[string]string){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/favicon.ico", favicon)
	log.Fatal(http.ListenAndServe(string(":")+string(envs["APIPORT"]), nil))
}

func variables() map[string]string {
	envs, err := dotenv.Read(".env")
	
	if err != nil {
		log.Fatal("Error while finding .env file")
	}

	return envs
}

// https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver-pt
func init(){
	envs := variables()
	ctx := context.TODO()
	mongodbURI := string(envs["MONGODBURI"])
	fmt.Print(mongodbURI+"\n")
	clientOptions := options.Client().ApplyURI(envs["MONGODBURI"])
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("User").Collection("user")
	fmt.Print("Server conectado!\n")
}

func main(){
	envs := variables()
	jwtsecret := envs["JWTSECRET"]
	apiport := envs["APIPORT"]
	// mongodbURI := envs["MONGODBURI"]
	workdir := envs["WORKDIR"]
	fmt.Println("JWTSECRET: "+jwtsecret)
	fmt.Println("APIPORT: "+apiport)
	// fmt.Println("MONGODBURI: "+mongodbURI)
	fmt.Println("WORKDIR: "+workdir)
	handleRequests(envs)
}
