package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"context"
	//"errors"
	//"os"
    //"path/filepath"
    //"strconv"
    "time"

	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"

	pb "./cases"
)

type casee struct {
	Name			string	`json:"name"`
	Location		string	`json:"location"`
	Age				int32	`json:"age"`
	Infectedtype 	string	`json:"infectedtype"`
	State 			string	`json:"state"`
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GO API!")
}

func getConnection(host string) (*grpc.ClientConn, error) {
	return grpc.Dial(host, grpc.WithInsecure())
	//return grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock());
}

func newCase(casees pb.Case) string {
    host := "localhost:4000"//os.Getenv("DISCOUNT_SERVICE_HOST")
    /*if len(host) == 0 {
        host = "localhost:4000"
	}*/
	conn, err := getConnection(host)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInsertClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var req = pb.CaseReq{
		Case : &pb.Case{Id: "sad", Name: "asd", Location: "sad", Age: 2, Infectedtype: "fdf", State: "gg"},
	}
	//var tmp casee;
	//json.Unmarshal(req.Case, &tmp)
	log.Println("F3.", req.Case)
	r, err := c.InsertCase(ctx, &req)
	if err == nil {
		log.Println("RES.", r);
		return "Case inserted.";
	} else {
		log.Println("Failed to insert case.", err)
		return "Failed to insert case.";
	}
}

func createCase(w http.ResponseWriter, r *http.Request) {
	var newcase casee
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Request Invalido")
	}
	
	json.Unmarshal(reqBody, &newcase)
	//var res = newCase(newcase)
	log.Println("Name.", newcase.Name)
	sendCase := pb.Case{Id: "0XsdJOk", Name: "hola", Location: "mundo", Age: 12, Infectedtype: "com", State: "sd"};
	//sendCase.Name = newcase.Name;
	//sendCase.Location = newcase.Location;
	//sendCase.Age = newcase.Age;
	//sendCase.Infectedtype = newcase.Infectedtype;
	//sendCase.State = newcase.State;

	host := "localhost:4000"
	conn, err := getConnection(host)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewInsertClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := pb.CaseReq{ Case: &sendCase}
	
	var res = "";
	resp, err := c.InsertCase(ctx, &req);
	if err == nil {
		log.Println("RES.", resp);
		res = "Case inserted.";
	} else {
		log.Println("Failed to insert case.", err)
		res = "Failed to insert case.";
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cases", createCase).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}

