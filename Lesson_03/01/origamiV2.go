package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

// superflui
type Piega struct {
	ascissa string
	valore  int
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	//creo le slices
	var coordinates []Coordinate
	var folds []Piega
	maxX, maxY := 0, 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text()) //rimozione spazi inutili

		if strings.Contains(line, "fold") {
			parts := strings.Split(line, "=")
			ascissa := strings.TrimSpace(parts[0][len(parts[0])-1:])
			valore, _ := strconv.Atoi(parts[1])
			folds = append(folds, Piega{ascissa, valore})

		} else {
			parts := strings.Split(line, ",")
			x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			coordinates = append(coordinates, Coordinate{x, y})

		}
	}

	for _, coord := range coordinates {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}
	//controllo le coordinate
	// for _, k := range coordinates {
	// 	fmt.Println(k.x, k.y)
	// }

	grid := make([][]rune, maxY+1)
	for i := range grid {
		grid[i] = make([]rune, maxX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for _, coord := range coordinates {
		grid[coord.y][coord.x] = '#'
	}
	//prima griglia
	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }

	//l := 1
	//println("io maxX e: ", maxX, "e il maxY:", maxY)
	for _, fold := range folds {
		if fold.ascissa == "y" {
			// Piega lungo y
			for y := fold.valore + 1; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {

					//println("entrato: ", l)
					// l += 1
					if grid[y][x] == '#' {
						//println("qua")

						newY := fold.valore - (y - fold.valore)
						//println("questo e il newY: ", newY)
						grid[newY][x] = '#'
					}
				}
			}

			//println("dfuiogfdfgiuyfdyudokiuytrewssdrftyuiopoiuytrewasdfghjklòlkjhgfdsazxcvbnjklòloiuytfrd")

			grid = grid[:fold.valore] // Tagliare la griglia alla linea di piegatura
			// fmt.Println("griglia dopo il taglio di y=", fold.valore)
			// for _, row := range grid {
			// 	fmt.Println(string(row))
			// }
			// println("arrivi quaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			// lunghezza := len(grid)
			// println("questa e la lunghezza della griglia dopo", lunghezza)
		} else if fold.ascissa == "x" {
			// Piega lungo x
			for y := 0; y < len(grid); y++ {
				for x := fold.valore + 1; x <= maxX; x++ {
					if grid[y][x] == '#' {
						newX := fold.valore - (x - fold.valore)
						grid[y][newX] = '#'
						// println("stampa il newX", newX)
					}
				}
			}
			for y := range grid {
				if fold.valore < len(grid[y]) {
					grid[y] = grid[y][:fold.valore] // Tagliare le colonne alla linea di piegatura
				}
			}
			// fmt.Println("griglia dopo il taglio di x=", fold.valore)
			// for _, row := range grid {
			// 	fmt.Println(string(row))
			// }
		}
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	// fmt.Println("Folds:")
	// for _, fold := range folds {
	// 	fmt.Println(fold)
	// }

}
