# -*- coding: utf-8 -*-
"""
Created on Tue Jun 16 09:36:34 2020

@author: akshith reddy
"""
'''
#recursive function
def fact(n, res):
    if n > 1:
        res *= n
        return fact(n-1, res)
    return res
    
def factorial(num):
    if num == 0|1:
        return 1
    r = 1
    return fact(num, r)
'''

'''
#recursive
def factorial(num):
    if num == 0|1:
        return 1
    return num * factorial(num - 1)
'''

'''
def factorial(num):
    if num == 0|1:
        return 1
    res = 1
    for i in range(2, num + 1):
        res *= i
    return res
'''

#print(factorial(int(input('Enter a number: '))))