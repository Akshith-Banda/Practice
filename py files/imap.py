# -*- coding: utf-8 -*-
"""
Created on Tue Jun  2 14:51:15 2020

@author: akshith reddy
"""

import  imaplib
import base64
import os
import email

user_name = 'akshith.banda@gmail.com'
password = input('enter password')

mail = imaplib.IMAP4_SSL('imap.gmail.com')
mail.login(user_name, password)
mail.select('INBOX')
i =0 
tpe, data = mail.search(None, 'ALL')
print(tpe)
for num in data[0].split():
    tpe, data = mail.fetch(num, '(RFC822)')
    raw_data = data[0][1]
    raw_data_str = raw_data.decode('utf-8')
    email_msg = email.message_from_string(raw_data_str)
    for part in email_msg.walk():
        if part.get_content_type() == 'text/plain':
            body = part.get_payload(decode=True)
            print(body,'\n\n')
    i += 1
    if i > 2:
        break
'''
    email_subject = email_msg['subject']
    email_body = email_msg.get_payload(decode=True)
    print('subject ',email_subject)
    print('body ',email_body)
    break
'''
