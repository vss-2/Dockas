import redis
from sys import argv

client = redis.Redis(host='localhost', port=6379)

def send_message() -> int:
    # TODO: Check user auth
    if len(argv) == 3:
        # publish "channel" in "message"
        num_delivered = client.publish(argv[1], argv[2])
        return num_delivered
    else:
        print("Invalid number of arguments: must be exactly 2. "
            "If your message contains whitespaces, " 
            "make sure format by adding \\ before every whitespace. "
            "Other special characters must be treated!")
        return -1

send_message()
