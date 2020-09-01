# -*- coding: utf-8 -*-
"""
Created on Thu Jun 11 09:18:47 2020

@author: akshith reddy
"""

import socket
import prime_alg

HOST = '127.0.0.1'
PORT = 65431

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.bind((HOST, PORT))
    s.listen()
    conn, addr = s.accept()
    with conn:
        print('connected by ',addr)
        while True:
            data = conn.recv(1024)
            if not data:
                break
            numbers = list(str(data).split(','))
            lis = list()
            for ele in numbers:
                if numbers.index(ele) % 2 == 0:
                    l_ele = ele.split('(')
                    lis.append(int(l_ele[1]))
                elif numbers.index(ele) % 2 == 1:
                    r_ele = ele.split(')')
                    lis.append(int(r_ele[0]))
            lis_tup = list()
            f_res = list()
            for j in range(0,len(lis),2):
                a = lis[j]
                b= lis[j+1]
                lis_tup.append(tuple((a, b)))
            for num in lis_tup:
                f_res.append(prime_alg.co_prime(num[0], num[1]))
                res = ','.join(f_res)
            conn.sendall(res.encode('utf-8'))
        