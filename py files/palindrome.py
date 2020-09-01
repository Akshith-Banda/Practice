# -*- coding: utf-8 -*-
"""
Created on Tue Jun 16 07:19:39 2020

@author: akshith reddy
"""
'''
def palin(txt):
    if len(txt) % 2 == 0:
        len_txt = len(txt) / 2
    else:
        len_txt = (len(txt) - 1) / 2
    for i in range(int(len_txt)):
        if txt[i] == txt[-(1+i)]:
            continue
        else:
            return 'Not a Palindorme'
    return 'Palindrome'

'''
'''
# recursive function for palindrome
def ispalin(st, s, e):
    if s == e:
        return 'Palindrome'
    if st[s] != st[e]:
        return 'Not a palindrome'
    if s < e:
        return ispalin(st, s+1, e-1)
    return 'Palindrome'

def palin(txt):
    len_txt = len(txt)
    
    if len(txt) == 0:
        return 'Palindrome'
    return ispalin(txt, 0, len_txt - 1)
'''

#print(palin(input('Enter text: ')))

#reverse of a number. (recursive)
def rev(num, temp=0):
    if num != 0:
        temp = int(num % 10) + (int(temp) * 10)
        return rev(int(num / 10), temp)
    return temp

print(rev(int(input('Enter a number: '))))

























