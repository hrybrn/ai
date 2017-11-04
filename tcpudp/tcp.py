import socket

address = 'wsn.ecs.soton.ac.uk'
port = 5002
bufferSize = 1024
message = 'hb15g16'

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((address, port))
s.send(message)
data = s.recv(bufferSize)
s.close()
print "received data:", data