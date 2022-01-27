# Ordem de execuçao

- Executar em um terminal &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; ```bash ./create_mongonet.sh```
- Executar em outro terminal &nbsp;&nbsp; ```bash ./mongo-teste.sh```
- Executar em outro terminal &nbsp;&nbsp;&nbsp;```bash ./mongo-expressteste.sh```
- Criar ```.env``` e adicionar MONGO_URI vinda do ```docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mongoteste```
- Executar ```docker build```
- Executar ```npm start 4001```
- Opcional ```docker exec -it {{id nodeteste}} bash```
- Verificar se está tudo funcionando. Enviar requests para 0.0.0.0:4001.
