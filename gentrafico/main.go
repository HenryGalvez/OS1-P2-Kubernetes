package main

import "fmt"
import "io/ioutil"
import "encoding/json"

type infocase struct {
    Name   string
	Location string
	Age int
	Infectedtype string
	State string
}

type cases struct {
	Casos []infocase
}

func main() {
	var op int
	var gorutinas int
	var nsolicitudes int
	var urlbalancer string
	var cases cases
	var path string
	op = 0;
	gorutinas = 0;
	nsolicitudes = 0;
	
	for op!=6 {
		fmt.Println("Menu")
		fmt.Println("1. URL del balanceador de carga a donde se desea enviar")
		fmt.Println("2. Cantidad de gorutinas que se ejecutarán para enviar los datos")
		fmt.Println("3. Cantidad de solicitudes que tiene el archivos")
		fmt.Println("4. Ruta del archivo que se desea cargar")
		fmt.Println("5. Enviar datos")
		fmt.Println("6. Salir")
		fmt.Print("Opcion a escoger: ")
		fmt.Scanf("%d", &op)
		switch op {
		case 1:
			fmt.Println("\033[2J")
			fmt.Print("Ingrese la ruta: ")
			fmt.Scanf("%s", &urlbalancer)
		case 2:
			fmt.Println("\033[2J")
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scanf("%d", &gorutinas)
		case 3:
			fmt.Println("\033[2J")
			fmt.Print("Ingrese la cantidad: ")
			fmt.Scanf("%d", &nsolicitudes)
		case 4:
			fmt.Println("\033[2J")
			fmt.Print("Ingrese la ruta: ")
			fmt.Scanf("%s", &path)
			cases = readfile(path)
		}
		fmt.Println("\033[2J")
	}
}

func readfile(path string) cases{
	var cases cases
	data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("File reading error", err)
        return cases
	}
	json.Unmarshal([]byte(string(data)), &cases)
	return cases
}