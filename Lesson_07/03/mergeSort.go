// Scrivete una funzione che implementa l’algoritmo ricorsivo mergeSort. La funzione:
// 1. divide il vettore in due sotto-vettori di dimensione circa uguale;
// 2. ordina il sotto-vettore di sinistra richiamando se stessa;
// 3. ordina il sotto-vettore di destra richiamando se stessa;
// 4. integra (merge) i due vettori in un vettore ordinato.
// La base della ricorsione è, anche qui, data dai vattori di lunghezza 0 o 1, che sono sempre
// ordinati.
// La parte di integrazione (merge) di due vettore ordinati a1 e a2 funziona con un vettore di
// supporto: si scorrono entrambi i vettori da sinistra a destra usando due indicatori i1 e i2 rispetti-
// vamente; ad ogni passo si confronta a1[i1] con a2[i2] e si sceglie l’elemento più piccolo, lo si
// copia nel vettore di supporto (nella prima posizione libera) e si incrementa l’indicatore relativo
// ad esso. Quando i1 esce da a1 oppure i2 esce da a2, la parte rimanente dell’altro vettore viene
// copiata nel vettore di supporto. Alla fine si copia il contenuto del vettore di supporto nel vettore
// originale.
// 2

package main

import "fmt"

func main() {
	s := []int{15, 96, 44, 22, 54, 28, 83}
	fmt.Println("Before: ", s)
	fmt.Println("After: ", mergeSort(s))
}

func mergeSort(slice []int) []int {
	fmt.Println("mergeSort: ", slice)
	if len(slice) == 1 {
		return slice
	} else if len(slice) == 2 {
		if slice[0] > slice[1] {
			slice[0], slice[1] = slice[1], slice[0]
		}
		return slice
	} else {
		mid := len(slice) / 2
		left := mergeSort(slice[:mid])
		right := mergeSort(slice[mid:])
		return merge(left, right...)
	}
}

func merge(left []int, right ...int) []int {
	result := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	return result
}