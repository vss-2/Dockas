import redis
from flask import Flask, request, jsonify

redis = redis.Redis(host='localhost', port=6379, decode_responses=True)
FLASK_PORT = 8080
app = Flask('python-redis')


def simple_start() -> bool:
	if redis.set('test', 'ok'):
		if redis.get('test') == 'ok':
			print('System is running and looks fine!')
			redis.delete('test')
			return True
	return False

@app.route('/subscribe', methods=['POST'])
def sub_subs_new_channel():
	# Adds a new channel to pair (user: channels) [str: set[str]] in REDIS

	req = request.json
	if req['channel_id'] and req['user_id']:
		# TODO: Check user auth

		redis.sadd(req['user_id'], req['channel_id'])

		response = jsonify({'msg': 'success'})
		response.headers.add('Access-Control-Allow-Origin', '*')
	else:
		if not req['channel_id']:
			response = jsonify({'msg': 'invalid channel_id'})
		else:
			response = jsonify({'msg': 'invalid user_id'})

	return response

@app.route('/list/subscribed', methods=['GET'])
def sub_requests_all_channels():
	# Returns all subscribed channels in a list format for a given user

	req = request.json

	# TODO: Check user auth
	if req['user_id']:
		channels = redis.smembers(req['user_id'])
		channels = list(channels)
		return jsonify({'channels': channels})

@app.route('/')
def initializer():
	return 'Hello World, python-redis!'

if __name__ == "__main__":
	if not simple_start():
		print('Something is wrong, check if server is up and running at port 6379.')
		exit(1)
	initializer()

app.run(port=FLASK_PORT)
