# gRPC in Python

## Dependencies
We need to have installed to Python:
```
$ pip install grpcio grpcio-tools
```

Commands to execute:
```
$ python -m grpc_tools.protoc -I.. --python_out=. --grpc_python_out=. ../cases.proto
```
or
```
$ python -m grpc_tools.protoc -I=.. --python_out=. --grpc_python_out=. ../cases.proto
```

*this command is used to compile the **.proto** file and obtain the **_pb2.py** files to create the gRPC server*

To run the projecto will be execute
```
$ python app.py
```