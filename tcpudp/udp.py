import socket

address = 'wsn.ecs.soton.ac.uk'
port = 5005
message = 'hb15g16'
bufferSize = 4096

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.sendto(message, (address, port))

data, server = sock.recvfrom(bufferSize)

print data