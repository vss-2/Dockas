docker run -p 6379:6379 -v "${PWD}"/data/:/data --name $@ -d redis:7.2-rc2-alpine3.18 redis-server

# Deploy Flask API
python3 main.py

# Subscribe github_user to github_channel using Flask API to save subscriber's list in a REDIS Set
curl --request POST \
  --url http://localhost:8080/subscribe \
  --header 'Content-Type: application/json' \
  --data '{ "channel_id": "github_channel", "user_id": "github_user" }'

# Subscribe github_user starts listening
python3 sub.py github_user

# Publisher sends a "Hello World" to github_channel
python3 pub.py github_channel "Hello\ World"
