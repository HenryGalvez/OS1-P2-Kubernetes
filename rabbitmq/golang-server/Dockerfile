FROM golang
COPY main.go .

RUN go get -u github.com/gorilla/mux
RUN go get github.com/streadway/amqp

EXPOSE 3000

CMD ["go", "run", "main.go"]