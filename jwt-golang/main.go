package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	dotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

// https://www.honeybadger.io/blog/go-web-services/
// https://golangbyexample.com/200-http-status-response-golang/
func userRegister(w http.ResponseWriter, r *http.Request){
	if (r.Header["Name_first"] != nil && 
		r.Header["Name_last"] != nil &&
		r.Header["Email"] != nil &&
		r.Header["Password"] != nil){
			http.Error(w, "Missing one or more parameters: first name, last name, email or password", http.StatusBadRequest)
			return
	} else {
		cursor, err := collection.Find(ctx, bson.M{})
		if(err != nil){
			log.Fatal(err)
		}

		// Query if there's already a user with the same email registered
		var user []bson.M
		if err = cursor.All(ctx, bson.M{"Email": r.Header["Email"]}, &user); err != nil {
			log.Fatal(err)
		}
		// If there isn't insert one in
		if(len(user) == 0){
			token, err := jwt.createToken(envs["JWTSECRET"]); err != nil {
				log.Fatal(err)
			}
			if res, err := cursor.InsertOne(ctx.Background(), bson.M{
				"Name_first": r.Header["Name_firl"], 
				"Name_last": r.Header["Name_last"], 
				"Email": r.Header["Email"],
				"Password": r.Header["Password"],
				"Token": token,
				"expiresAfterSeconds": 60*60*12
			});
			err != nil {
				log.Fatal(err)
			} else {
				// Return request correctly
				resp := make(map[string]string)
				resp["message"] = "Registered with success"
				resp["Token"] = token
				if jsonResp, err := json.Marshal(resp); err != nil {
					log.Fatal(err)
				}
				w.WriterHeader(http.StatusOK)
				w.Write(jsonResp)
				return
			}
		} else {
			http.Error(w, "User already registered in database, did you forget you password?", http.Conclict)
		}
	}
}

func userLogin(w http.ResponseWriter, r *http.Request){
	if (r.Header["Name_first"] != nil && 
		r.Header["Name_last"] != nil &&
		r.Header["Password"] != nil){
			r.Header["Email"] != nil &&
			// fmt.Fprintf("Missing email or password")
			http.Error(w, "Missing email or password", http.StatusBadRequest)
			return
	} else {
		var user []bson.M
		if err = cursor.All(ctx, bson.M{"Email": r.Header["Email"]}, &user); err != nil {
			log.Fatal(err)
		}
		if(len(user) == 1){
			token, err := jwt.createToken(envs["JWTSECRET"]); err != nil {
				log.Fatal(err)
			}
			resp := make(map[string]string)
			resp["message"] = "Login with success"
			resp["Token"] = token
			if jsonResp, err := json.Marshal(resp); err != nil {
				log.Fatal(err)
			}
			w.WriterHeader(http.StatusOK)
			w.Write(jsonResp)
			return
			}
		} else {
			http.Error(w, "Incorrect email or password", http.Forbidden)
		}
	}
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
		// fmt.Fprintf(w, "Access unauthorized")
		http.Error(w, "Access unauthorized, try to login first", http.Unauthorized)
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
	http.HandleFunc("/register", userRegister)
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
