# -*- coding: utf-8 -*-
"""
Created on Mon Jul  6 09:06:20 2020

@author: akshith reddy
"""
'''
import time

def dec(f):
    def wrapper(*args, **kwargs):
        start = time.clock()
        ret = f(*args, **kwargs)
        end = time.clock()
        print("time taken ", end - start)
        return ret
    return wrapper

@dec
def process(n):
    for i in range(n):
        i *= i

print(process(10000))
'''
def dec(f):
    def wrapper(*args, **kwargs):
        print("***********")
        ret = f(*args, **kwargs)
        print("***********")
        return ret
    return wrapper

@dec
def add(x, y):
    print("add value is ", x+y)

@dec
def greet():
    print("Hello! ")

add(3, 6)

greet()