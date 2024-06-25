#! /usr/bin/python
# a simple tcp server
import socket, sys

HOST, PORT = sys.argv[1], int(sys.argv[2])
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
    sock.bind((HOST, PORT))  
    sock.listen(5)  
    while True:  
        connection,address = sock.accept()  
        buf = connection.recv(1024)  
        print("received : {}", buf)
        connection.send(buf)    		
        connection.close()