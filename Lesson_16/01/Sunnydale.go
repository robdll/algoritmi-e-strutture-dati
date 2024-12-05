package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

type Junction struct {
	id int
	galleries []*Gallery
	visited bool
}

type Gallery struct {
	from int
	to int
	brightness int
}

// func AddEdge(graph []*linkedList, from int, to int , peso int){
// 	newNode := &listNode{destinazione: to, peso: peso, next: graph[from].iniziolista}
// 	graph[from].iniziolista = newNode
// }


// func PrintGraph( graph []*linkedList){
// 	for i,lista:=range graph{
// 		if lista.iniziolista != nil {
// 			Print(" nodo", i , ":")
// 			current := lista.iniziolista
// 			for current != nil{
// 				Print(" -> ",current.destinazione, "( peso: ", current.peso, ")" )
// 				current = current.next
// 			}
// 			Println()
// 		}
// 	}
// }

func main(){

	file, err := os.Open("VampireDates.txt")
	if err != nil {
		Println("Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var startPath, endPath int

	if scanner.Scan(){
		input := strings.Fields(scanner.Text())
		start, _ := strconv.Atoi(input[2])
		end, _ := strconv.Atoi(input[3])
		startPath = start
		endPath = end
	}


	junctionMap := make(map[int]Junction)
		
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		from, _ := strconv.Atoi(input[0])
		to, _ := strconv.Atoi(input[1])
		brightness, _ := strconv.Atoi(input[2])

		if junction, exists := junctionMap[from]; exists{
			var newGallery = &Gallery{from, to, brightness}
			junction.galleries = append(junctionMap[from].galleries, newGallery)
			junctionMap[from] = junction
		}	else {
			var gallery = []*Gallery{{from, to, brightness}}
			var junction = Junction{id: from, galleries: gallery, visited: false}
			junctionMap[from] = junction
		}
	}

	currentJunction, _ := junctionMap[startPath]
	endJunction, _ := junctionMap[endPath]
	
	var steps int
	for currentJunction.id != endJunction.id && currentJunction.visited == false {
		steps++
		currentJunction.visited = true
		var darkestGallery *Gallery
		for _, gallery := range currentJunction.galleries {
			if darkestGallery == nil || gallery.brightness < darkestGallery.brightness {
				darkestGallery = gallery
			}
		}
		Println("going to: ", darkestGallery.to)
		currentJunction = junctionMap[darkestGallery.to]
	}


	if currentJunction.id == endJunction.id {
		Println("arrived in ", steps, "steps")
	} else {
		Println("Loop after #", steps, "steps")
	}
}