# -*- coding: utf-8 -*-
"""
Created on Mon Jun  8 09:04:31 2020

@author: akshith reddy
"""

#help('modules')
'''
import pip
installed_pckg = pip._internal.utils.misc.get_installed_distributions()
flat_installed_pckg = [package.project_name 
                       for package in installed_pckg 
                       if package.project_name.startswith('a')]
print(flat_installed_pckg)

import sys
print(repr(sys.byteorder))
'''
import platform
plt = platform.python_implementation()
print(plt)