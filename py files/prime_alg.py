# -*- coding: utf-8 -*-
"""
Created on Wed Jun 10 08:58:11 2020

@author: akshith reddy
"""

def prime(num):
    dummy_int = 0
    if num == 2:
        return 'Prime number '
    elif num % 2 != 0:
        for i in range(2, num):
            if num % i != 0:
                dummy_int += 1
                continue
            else:
                dummy_int = 0
                return 'Not a prime '
    else:
        return 'Not a prime '
    if dummy_int >= 1:
        return 'Prime number '

'''
num = input('Enter a number: ')
print(prime(int(num)))

numbers = list(map(int, input("Enter multiple numbers with ',': ").split(',')))
for n in numbers:
    print(n, 'is ',prime(n))
'''

def co_prime(num1, num2):
    dummy_int = 0
    num = num1 if num1 < num2 else num2
    for i in range(2, num+1):
        if (num1 % i == 0) & (num2 % i == 0):
            dummy_int += 1
            break      
    if dummy_int >= 1:
        return 'Not co-prime '
    else:
        return 'co-prime '
'''  
num1, num2 = map(int, input("Enter two numbers with ',': ").split(','))     
print(co_prime(num1, num2))    

numbers = list(input('enter numbers in tuples: ').split(','))
lis = list()
for ele in numbers:
    if numbers.index(ele) % 2 == 0:
        l_ele = ele.split('(')
        lis.append(int(l_ele[1]))
    elif numbers.index(ele) % 2 == 1:
        r_ele = ele.split(')')
        lis.append(int(r_ele[0]))
lis_tup = list()
for j in range(0,len(lis),2):
    a = lis[j]
    b= lis[j+1]
    lis_tup.append(tuple((a, b)))
for num in lis_tup:
    print(co_prime(num[0], num[1]))
'''
