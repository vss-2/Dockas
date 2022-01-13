docker run -it -p 12345:27017 --network `cat mongonet_id.txt` --name mongoteste mongo
echo Adicione ao .env `docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mongoteste`
