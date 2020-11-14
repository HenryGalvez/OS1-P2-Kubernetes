# we take the image of Golang
FROM golang

# we set the directory to use inside container
WORKDIR /go

# we copy the files of our api inside actual folder and send inside container folder
COPY . /go

# we install all necessary dependencies
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/proto
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u github.com/gorilla/mux

# we open the port 3000
EXPOSE 5000

# we create a folder to mount and use to read the kernel modules
RUN mkdir -p /go

# we run the command inside the container to run the api
CMD [ "go", "run", "main.go" ]
