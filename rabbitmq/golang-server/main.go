package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"math/rand"
    "github.com/streadway/amqp"
)

type casee struct {
	Name			string	`json:"name"`
	Location		string	`json:"location"`
	Age				int		`json:"age"`
	Infectedtype 	string	`json:"infectedtype"`
	State 			string	`json:"state"`
}

type responsee struct {
	Status		bool `json:"status"`
	Message		string `json:"message"`
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API!")
}

func createCase(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Request Invalido")
	}
    res, err := llamada(string(reqBody))
	failOnError(err, "Failed to handle RPC request")
	
	var ress responsee
	json.Unmarshal([]byte(res), &ress)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ress)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cases", createCase).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}




//Funciones RabbitMQ
func failOnError(err error, msg string) {
	if err != nil {
			log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
			bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func llamada(n string) (res string, err error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // noWait
			nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrId := randomString(32)

	err = ch.Publish(
			"",          // exchange
			"rpc_queue", // routing key
			false,       // mandatory
			false,       // immediate
			amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: corrId,
					ReplyTo:       q.Name,
					Body:          []byte(n),
			})
	failOnError(err, "Failed to publish a message")

	for d := range msgs {
			if corrId == d.CorrelationId {
					res = string(d.Body)
					break
			}
	}

	return
}