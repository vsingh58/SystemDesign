#! /usr/bin/python
# a simple tcp server
import socket, sys

HOST, PORT = sys.argv[1], int(sys.argv[2])
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
    sock.bind((HOST, PORT))  
    sock.listen()      
    print("API Server listening on port", PORT)
    connection,address = sock.accept()  
    print("API Server accepted connection")
    with connection:
        while True:              
            data = connection.recv(1024) 
            if data: 
                print("received :", str(data.decode('utf-8')))
                connection.sendall(data)    		
    
    