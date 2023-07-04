# python-redis

## Description
### A simple Python API to save temporary keys in a simple REDIS (Remote Dictionary Server) database.

## Requirements
* Redis 7.0.2
* Python 3.7 or latest (check `requirements.txt` to see pip dependencies)


## Usage Example
Check `startup.sh` to run without using dockerfile and run passing a `name`.
Ex: `bash startup.sh my-redis`

Using the dockerfile the volume won't be automatically mapped to `/data` folder and the name is hardcoded to `python-redis`.

Run `pip3 install -r requirements.txt && python3 main.py` to check if REDIS is running.

## Specs
Check `redis.conf` to review customized variables used.
