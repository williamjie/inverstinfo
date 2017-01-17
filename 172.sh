#! /usr/bin/expect -f 
#scp -r * root@10.2.113.172:/root/vueAdmin 
scp -r ./server/* root@10.2.113.172:/root/vueAdmin/server

#expect "*password:" 
#send "your password\r"
#         zaq1@WSX
#expect eof 



#!/usr/bin/expect