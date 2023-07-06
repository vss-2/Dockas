from time import sleep
import requests
import redis
import asyncio
from sys import argv
from json import dumps

# Default localhost and 8080
FLASK_URL='http://127.0.0.1'
FLASK_PORT=8080

client = redis.Redis(host='localhost', port=6379)
p = client.pubsub()

req = {'status_code': -1}
res = {'channels': []}
current_subscribed = set()

if len(argv) < 1:
    print('Missing user_id')
    exit(1)

def channel_finder() -> bool:
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
    return False

def subscribe() -> None:
    # Fetchs user subscribed channels
    while len(res['channels']) == 0:
        ok = channel_finder()
        if not ok:
            print('Some error trying to find channel."\
                  "Check address, port, firewall and if REDIS is up!')
            sleep(10)
            continue
        if len(res['channels']) == 0:
            print('No channels subscribed. Retrying connection!')
            sleep(5)

async def check_new_subs() -> list:    
    # This method avoids subscribing over and over with no changes
    
    while True:
        global current_subscribed
        if len(current_subscribed) == 0:
            # First time running this method
            current_subscribed = set(res['channels'])
        else:
            # Subscribe to new channel (that is not in current subscribed ones)
            channel_finder()
            channel_diff = set(res['channels']).difference(current_subscribed)
            if len(channel_diff) > 0:
                global p
                p.subscribe(list(channel_diff))
            
                # Update current channels
                current_subscribed = res['channels']
                return current_subscribed
            else:
                return []
        asyncio.sleep(5)

subscribe()

p.subscribe(res['channels'])

asyncio.run(check_new_subs())

def event_receiver() -> None:
    while True:
        # Waits before checking for new update
        msg = p.get_message()

        if msg != None:
            
            if msg['type'] == 'subscribe':
                print('Subscribing to channel', msg['channel'].decode('utf-8'))
            elif msg['type'] == 'message':
                print('<{channel}>: {data}'
                    .format(channel = msg['channel'].decode('utf-8'), data = msg['data'].decode('utf-8')))
            else:
                print('Debug:', msg)

        sleep(5)

event_receiver()
