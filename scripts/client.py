import socket
import time

HOST = 'localhost'
PORT = 4000
messages = ["hello\n", "hey\n", "how\n", "whats up\n", "hi\n"]

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    print("connection successful")
    for message in messages:
        print("message being send ---> ", message)
        s.sendall(message.encode())
        data = s.recv(1024)
        print('Received:', data.decode())
        time.sleep(1)
