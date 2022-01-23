# grpc-calc

gRPC utilizes http2 and protocol buffers to perform bidirectional streaming (multiplexing) and to enable programs from different languages to communicate to each other. gRPC uses protobuf as it's interface definition language and we use proto3 here. The main benefit of using this protobuf ecosystem is that computers can process protobufs faster over parsing a serialized json data and they are significantly of less size than json objects. Faster execution and and lower network bandwidth. 

## Dependencies:
- All the dependencies for golang server and client are managed through go.mod
- Dependencies for python implementation is under ` python-imp/requirements.txt` 

## Generate Native language protobuff implementation
### For golang:
` cd go/ && ./generatepb.sh` 


### For python:
` cd python-imp/ && ./generatepb.sh` 
## Server 
` 
    cd go/server
` 

` go run main.go ` 

## Client
` cd go/client` 

`go run main.go` 

## Client (Python)
` cd python-imp/client/` 

`python client.py` 
