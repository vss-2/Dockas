docker build --name nodeteste .
docker run -p 0.0.0.0:4001:4001 --network mongonet nodeteste
