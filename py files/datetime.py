# -*- coding: utf-8 -*-
"""
Created on Wed Jul 22 09:59:26 2020

@author: akshith reddy
"""
from datetime import timezone

import datetime

todays_date = datetime.date.today()
print(todays_date)

past_date =datetime.date.today()
print(past_date)

diff = todays_date - past_date
print('difference in dates ', diff.days)
