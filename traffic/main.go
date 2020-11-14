package main

import (
	"fmt"
	"time"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"bytes"
)

type Casee struct {
	Name string `json:"name"`;
	Location string `json:"location"`;
	Age int `json:"age"`;
	Infectedtype string `json:"infectedtype"`;
	State string `json:"state"`;
}

type Cases struct {
	Casos [] interface{}
}

var urlLB = "http://34.121.142.125/cases";
var noReq = 1;
var path = "cases.json";

func readFile() (Cases){
	fmt.Println("Leyendo Path actual: ",path)
	b, err := ioutil.ReadFile(path)
    if err != nil { 
     fmt.Print(err) 
	}
	var c Cases
	json.Unmarshal(b, &c);
	//casees = c
	return c;
	/*out, err := json.Marshal(casees)
    if err != nil {
        panic (err)
	}
	fmt.Println(string(out))
	for i:= 0; i<100; i++ {
		v := c.Casos[i].(map[string]interface{})
		fmt.Println(v["name"])
	}
	*/
}

func executeGoR(c Cases){
	fmt.Println("Rutinas a ejecutar: ", noReq);
	for i:= 0; i< noReq; i++ {
		go thread(i, c);
	}
	time.Sleep(10* time.Second);
}

func thread(value int, c Cases) {
	v := c.Casos[value].(map[string]interface{})
	b := &Casee{Name: v["name"].(string),Location: v["location"].(string),Age: int(v["age"].(float64)),Infectedtype: v["infectedtype"].(string),State: v["state"].(string)}
	ff, err := json.Marshal(b);
	if err != nil {
		log.Fatalf("Error en codificacion como JSON: %v", err)
	}
	res, err := http.Post(urlLB,"application/json",bytes.NewBuffer(ff))
	if err != nil {
		log.Fatalf("Ocurrio al enviar la peticion: %v", err)
	}
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Se envio:\t", v["name"]);
	log.Println(string(body))
	time.Sleep(200*time.Millisecond)
}

func main () {
	var tmp1 = ""
	var tmp2 = ""
	var tmp3 = ""
	var tmp4 = ""
	for true {
		fmt.Println("URL del Balanceador de carga.\tActual: ", urlLB)
		fmt.Scanln(&tmp1)
		if(tmp1 != ""){
			urlLB = tmp1;
		}
		fmt.Println("Numero de solicitudes a leer.\tActual: ", noReq)
		fmt.Scanln(&tmp2)
		if(tmp2 != ""){
			noReq, _ = strconv.Atoi(tmp2);
		}
		fmt.Println("Path del archivo con los datos.\tActual: ",path)
		fmt.Scanln(&tmp3)
		if(tmp3 != ""){
			path = tmp3;
		}
		executeGoR(readFile());
		fmt.Println("Desea salir del programa?(S/N)");
		fmt.Scanln(&tmp4)
		if(tmp4 == "S" || tmp4 == "s") {
			break;
		}
		tmp4 = "";
	}
}