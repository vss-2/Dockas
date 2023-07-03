# jwt-node

## Description
### Implements JWT token authentication and authorization standard using NodeJS container. A MongoDB container is also included and it is required to be running to store session keys and users data. Containers are connected by creating a docker network. API port, Mongo_URI (internal network variable ID) and JWT token-key can be changed in .env file (check .env.example for expected system environment variables). 

## Requirements
* Mongo 5.0.5
* Mongo-express 1.0.0
* NodeJS (check package.json for specific packages reqs.)

## Usage example
Check the following tutorial section about how executing the containers. Tests are also implemented, check app.test.js .

## Setting up tutorial
- Executar em um terminal &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ```bash ./create_mongonet.sh```
- Executar em outro terminal &nbsp;&nbsp; ```bash ./mongo-teste.sh```
- Executar em outro terminal &nbsp;&nbsp;&nbsp;```bash ./mongo-expressteste.sh```
- Criar ```.env``` e adicionar MONGO_URI vinda do ```docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mongoteste```
- Executar ```docker build```
- Executar ```npm start 4001```
- Opcional ```docker exec -it {{id nodeteste}} bash```
- Verificar se está tudo funcionando. Enviar requests para 0.0.0.0:4001.
