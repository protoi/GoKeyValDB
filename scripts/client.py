import socket
import time

HOST = 'localhost'
PORT = 4000
messages = [
    "get mykey ",
    "set something sdjkkdjf ",
    "del mykey ",
    "del ddfks fskdfsdkf dksjk fdjskjkf ",
    "set fkdj ",
    "set ",
    "get dh djf dh hfhdfh ",
    "del ",
]

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    print("connection successful")
    for message in messages:
        print("message being send ---> ", message)
        s.sendall(message.encode())
        data = s.recv(1024)
        print('Received:', data.decode())
        time.sleep(1)
