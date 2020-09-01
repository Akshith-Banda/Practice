# -*- coding: utf-8 -*-
"""
Created on Tue Jun 16 13:19:59 2020

@author: akshith reddy
"""

def TowerOfHanoi(n , from_rod, to_rod, aux_rod): 
    if n == 1: 
        print('first func')
        print("Move disk '1' from rod",from_rod,"to rod",to_rod)
        return
    print('second func')
    TowerOfHanoi(n-1, from_rod, aux_rod, to_rod) 
    print("Move disk",n,"from rod",from_rod,"to rod",to_rod)
    print('third func')
    TowerOfHanoi(n-1, aux_rod, to_rod, from_rod) 
    print('after third')
n = 3
TowerOfHanoi(n, 'A', 'C', 'B')