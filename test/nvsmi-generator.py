import random

data = f'''#Time        gpu   pwr gtemp mtemp    sm   mem   enc   dec  mclk  pclk
#HH:MM:SS    Idx     W     C     C     %     %     %     %   MHz   MHz
 17:36:40      0    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}     {random.randint(0,100)}   {random.randint(0,1500)}   {random.randint(0,1500)}
 17:36:40      1    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}     {random.randint(0,100)}   {random.randint(0,1500)}   {random.randint(0,1500)}
 17:36:40      2    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}     {random.randint(0,100)}   {random.randint(0,1500)}   {random.randint(0,1500)}
 17:36:40      3    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}    {random.randint(0,100)}    {random.randint(0,100)}     {random.randint(0,100)}     {random.randint(0,100)}   {random.randint(0,1500)}   {random.randint(0,1500)}
'''

print(data)
