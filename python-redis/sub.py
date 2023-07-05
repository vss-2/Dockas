from time import sleep
import requests
import redis
from json import loads
from sys import argv
from json import dumps

# Default localhost and 8080
FLASK_URL='http://127.0.0.1'
FLASK_PORT=8080

client = redis.Redis(host='localhost', port=6379)
p = client.pubsub()

req = {'status_code': -1}
res = {'channels': []}

if len(argv) < 1:
    print('Missing user_id')
    exit(1)

def channel_finder() -> bool:
    while(True):
        req = \
        requests.get(f'{FLASK_URL}:{FLASK_PORT}/list/subscribed',
            headers = {
                'Content-Type': 'application/json',
            },
            data = dumps({'user_id': argv[1]})
        )

        if req.ok:
            global res
            res = req.json()
            req.close()

            return True

        sleep(10)


while len(res['channels']) == 0:
    channel_finder()
    if len(res['channels']) == 0:
        print('No channels subscribed. Retrying connection!')

p.subscribe(res['channels'])

while True:
    # Waits before checking for new update
    msg = p.get_message()

    if msg != None:
        
        if msg['type'] == 'subscribe':
            print('Subscribing to channel', msg['channel'].decode('utf-8'))
        elif msg['type'] == 'message':
            print('<',msg['channel'].decode('utf-8'),'>: ',msg['data'].decode('utf-8'), sep='')
        else:
            print('Debug:', msg)

    sleep(5)
