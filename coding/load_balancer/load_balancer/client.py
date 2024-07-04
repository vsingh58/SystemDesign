#! /usr/bin/python
# a simple tcp server
import socket, sys

HOST, PORT = sys.argv[1], int(sys.argv[2])
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
    sock.connect((HOST, PORT))
    
    while True:
        user_input = input("Type a mesage:")
        if user_input == "exit":
            break 
        if user_input:            
            sock.sendall(bytes(user_input, "utf-8"))
            print("sent:", user_input)

        data = sock.recv(1024)
        if data:
            print("received:", str(data.decode('utf-8')))
        


    