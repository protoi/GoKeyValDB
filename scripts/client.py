import socket
import time

HOST = 'localhost'
PORT = 4000
messages = [
    "set a 100",
    "set b 200",
    "set b 300",
    "get a",
    "del a",
    "set c 300",
    "set d 400",
    "set greetings hello world",
    "set goodbyes im leaving bye bye",
    "len",
    "del b",
    "get greetings",
    "get goodbyes",
    "get d",
    "get c",
    "get greetings",
    "flush",
    "len",
    "dhfjdfj",
    "djfhfd"
]

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    print("connection successful")
    for raw_msg in messages:
        message = f'{len(raw_msg)}%{raw_msg}'
        print("message being send ---> ", message)
        s.sendall(message.encode())
        data = s.recv(200)
        print('Received:', data.decode())
        time.sleep(1)
