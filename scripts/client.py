import socket
import time

HOST = 'localhost'
PORT = 4000
messages = [
    # "KV INIT mymap",
    # "KV SET mymap aaa 111",
    # "KV GET mymap aaa",
    # "KV GET mymapbbb",
    # "KV SET mymap bbb 222",
    # "KV GET mymap bbb",
    # "KV SET mymap ccc 333",
    # "KV SET mymap ddd 444",
    # "KV SET mymap eee 555",
    # "KV SET mymap fff HELLOWORLD",
    # "KV GET mymap fff",
    # "KV DEL mymap fff",
    # "KV GET mymap fff",
    #
    # "KV FLUSH mymap",
    # "KV GET mymap bbb",
    # "KV INIT xxx",
    # "KV SET xxx aaa HELLOWORLD",
    # "KV SET xxx bbb ITISME",
    # "KV FLUSH xxx",
    # "KV FLUSH xxx",
    "PROTOI $AAAABBBBCCCCDDDD%%",
    "ZSET INIT firstset ",
    "ZSET INSERT firstset hello 1 ",
    "ZSET INSERT firstset world 3 ",
    "ZSET INSERT firstset whats 12 ",
    "ZSET INSERT firstset up 4 ",
    "ZSET INSERT firstset everyone 5 ",
    "ZSET INSERT firstset hii 10 ",
    "ZSET POPMIN firstset ",
    "ZSET POPMIN firstset ",
    "ZSET POPMIN firstset ",
    "ZSET POPMIN firstset ",
    "ZSET POPMIN firstset ",
    "ZSET POPMIN firstset ",
    "ZSET POPMIN firstset"

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
        # time.sleep(1)
