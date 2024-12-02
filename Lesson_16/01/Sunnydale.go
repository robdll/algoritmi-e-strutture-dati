package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	destinazione int
	peso int
	next *listNode
}

type linkedList struct {
	iniziolista *listNode
}



type informazioni struct{
	svincoli, gallerie, posHarm, posSarah int
}


func main(){


	var info informazioni

	file, err := os.Open("VampireDates.txt")
	if err != nil {
		Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()

	// Scanner per leggere il file riga per riga
	scanner := bufio.NewScanner(file)

	if scanner.Scan(){
		Input := strings.Fields(scanner.Text())

		info.svincoli,_ = strconv.Atoi(Input[0])
		info.gallerie,_ = strconv.Atoi(Input[1])
		info.posHarm,_ = strconv.Atoi(Input[2])
		info.posSarah,_ = strconv.Atoi(Input[3])

		// Inizializza il grafo
		graph := make([]*linkedList, info.svincoli+1) // Usa 1-based indexing
		for i := range graph {
			graph[i] = &linkedList{}
		}
		Print()
	}
	


	

	 
	

}