# Flora-spring

## Description
### Contains a container back-end implementation of both SQL and NoSQL conections (using MySQL and MongoDB) to save Flora's user and plant data.

## Requirements
* Java 11 with Spring Boot 2.7.1. Flora is build using Maven.
* MongoDB is expected to be running at port 27017 (and also a created database "florafiles")
* MySQL is expected tu be running at port 3306 (and also two created tables: "plants" and "users")
* Other Spring boot (login, password, database name...) can be edited at (flora-spring/flora/src/main/resources/application.properties) file.

<br>

## Usage example

* Save plant sensors' info
```
curl -X POST http://localhost:8080/vase/save \
     -H "Content-Type: application/json" \
     -d '{"temperature": 17,"luminosity": 6, "umidity": 86}' 
```

* Save new user
```
curl -X POST http://localhost:8080/user/register \
     -H "Content-Type: application/json" \
     -d '{ \
            "name": "User", \
            "surname": "Flora", \
            "email": "user@flora", \
            "password": "eqwh87yef79tsd8fa", \
            "loginMethod": "email", \
            "location": "Brazil", \
            "language": "Portuguese-BR" \
         }'
```
* Get list with all users
```
curl -X GET http://localhost:8080/user/list/all
```

## Install
```sh
git clone --depth 1 https://github.com/vss-2/Dockas
mv ./flora-spring/* ../
cd ../
rm -rf ./Dockas/
cd ./flora-spring/flora
maven install
```

## Initialize
```sh
maven spring-boot:run
```