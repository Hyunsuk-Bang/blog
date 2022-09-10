# blog
Simple blog API with gRpc and MongoDB
A simple Blog gRPC capable of CRUD with MongoDB. 

## Featue:
All gRPC requests and responses are encrypted with self-signed CA. 

## 0. Generate private keys and public keys for self-signed CA
$ cd ssl <br>
$ sh ./ssl.sh

## 1. Generate binary file 
$ make build <br>
This command will generate binary file for server and client. <br>
For Server: /bin/server <br>
For Client: /bin/client

