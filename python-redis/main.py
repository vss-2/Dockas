import redis

redis = redis.Redis(host='localhost', port=6379, decode_responses=True)

if redis.set('test', 'ok'):
	if redis.get('test') == 'ok':
		print('System is running and looks fine!')
