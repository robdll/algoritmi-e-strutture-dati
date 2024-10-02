// Data una lunga sequenza di N interi distinti che rappresentano le altezze di una catena montuosa,
// stampare il numero di salite che vanno da sinistra a destra (una salita è una sequenza crescente
// di 2 o più interi, che partono da un punto di minimo e arrivano in un punto di massimo).
// Esempio: nella sequenza 9 1 3 5 2 0 8 6 ci sono due salite: 1 3 5 e 0 8 (1 3 e 3 5 non
// sono salite perché la prima non finisce in un punto di massimo e la seconda non inizia in un
// punto di minimo).
package main

import "fmt"

func main() {
    // define an array with the following values 9, 1, 3, 5, 2, 0, 8, 6, 7
    var sequence = []int{9, 1, 3, 5, 2, 0, 8, 6, 7}
    var upHill = 0
    for i := 1; i < len(sequence); i++ {
        // declare current 
        var current = sequence[i]
        var isLast = i == len(sequence) - 1
        if (current > sequence[i - 1] && (isLast || current > sequence[i + 1])) {
            upHill++
        }
    }
    fmt.Println("uphills:", upHill)
}