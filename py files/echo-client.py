# -*- coding: utf-8 -*-
"""
Created on Thu Jun 11 09:28:14 2020

@author: akshith reddy
"""

import socket

HOST = '127.0.0.1'
PORT = 65431

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))     
    s.sendall(input('enter numbers in tuples: ').encode('utf-8'))
    data = s.recv(1024).decode('utf-8')

res = data.split(',')
for msg in res:
    print('rec ',msg)