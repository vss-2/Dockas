docker run -p 6379:6379 -v "${PWD}"/data/:/data --name $@ -d redis redis-server
