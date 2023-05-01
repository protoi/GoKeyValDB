import socket
import time

HOST = 'localhost'
PORT = 4000
messages = [
        "LIST INIT mylist",
        "LIST INIT mynewlist",
        
        "LIST PUSHFRONT mylist hello",
        "LIST PUSHFRONT mylist world",
        "LIST PUSHFRONT mylist aaaa",
        "LIST PUSHFRONT mylist 111",
        "LIST PUSHFRONT mylist 222",
        "LIST PUSHFRONT mylist 333",
       
        "LIST PUSHBACK mylist xxx",
        "LIST PUSHBACK mylist yyy",
        "LIST PUSHBACK mylist zzz",
        "LIST PUSHBACK mylist cccc",
        "LIST PUSHBACK mylist bbbbb",
        

        "LIST POPFRONT mylist",
        "LIST POPFRONT mylist",
        "LIST POPFRONT mylist",

        "LIST POPBACK mylist",
        "LIST POPBACK mylist",
        "LIST POPBACK mylist",

        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",
        "LIST PEEKFRONT mylist",

        "LIST INIT mylist",

        "LIST PUSHFRONT mynewlist hello",
        "LIST PUSHFRONT mynewlist world",
        "LIST PUSHFRONT mynewlist aaaa",
        "LIST PUSHFRONT mynewlist 111",
        "LIST PUSHFRONT mynewlist 222",
        "LIST PUSHFRONT mynewlist  333",
       
        "LIST PUSHBACK mynewlist xxx",
        "LIST PUSHBACK mynewlist yyy",
        "LIST PUSHBACK mynewlist zzz",
        "LIST PUSHBACK mynewlist cccc",
        "LIST PUSHBACK mynewlist bbbbb",
        

        "LIST POPFRONT mynewlist",
        "LIST POPFRONT mynewlist",
        "LIST POPFRONT mynewlist",

        "LIST POPBACK mynewlist",
        "LIST POPBACK mynewlist",
        "LIST POPBACK mynewlist",

        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",
        "LIST PEEKFRONT mynewlist",


        "LIST POPFRONT mynewlist",
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        "LIST POPFRONT mynewlist", 
        






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
