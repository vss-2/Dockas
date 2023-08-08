# python-redis

## Description
### A simple Python API to save temporary keys in REDIS (Remote Dictionary Server) database. A PubSub example is included.

## Requirements
* Redis 7.0.2
* Python 3.7 or latest (check `requirements.txt` to see pip dependencies)
* cURL

## Usage Example
Check `startup.sh` to run without using dockerfile and run passing a `name`.
Ex: `bash startup.sh my-redis`

Using the dockerfile the volume won't be automatically mapped to `/data` folder and the name is hardcoded to `python-redis`.

Run `pip3 install -r requirements.txt && python3 main.py` to check if REDIS is running.

## Specs
Check `redis.conf` to review the customized variables used.
Flask default URL and PORT can (and if so, MUST BE) eddited in all .py files (main, sub, pub).

## PubSub step-by-step tutorial
* Keep running a single `python3 main.py` instance
* Send a request to `localhost:8080/add/subscribe` with `$user_id` and `$channel_id` defined in JSON body
* Have fun creating `python3 sub.py $user_id` (for receiving messages) or `python3 pub.py $channel_id $msg` (for publishing messages). 

## Usage example
* Terminal 1
```sh
python3 main.py &
curl -X POST localhost:8080/add/subscribe \
     -H "Content-Type: application/json" \
     -d '{ \
            "user_id": "github_tester", \
            "channel_id": "python_redis"
         }'
python3 sub.py github_tester python_redis
```

* Terminal 2
```sh
python3 pub.py python_redis "Hello python_redis subscribers World!"
```