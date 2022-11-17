# Author : Nemuel Wainaina

from socket import *

HOST = "127.0.0.1"
PORT = 12345

s = socket(AF_INET, SOCK_STREAM)
s.bind((HOST, PORT))
s.listen()
print(f"Listening for incoming connections")

while True:
    conn, addr = s.accept()
    print(f"Accepted connection from {addr[0]}:{addr[1]}")
    cmd = input("(Master)# ").encode('utf-8')
    print(cmd, file=s)
    result = s.recv(4096).decode('utf-8')
    print(result)
