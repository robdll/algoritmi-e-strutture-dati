// 2 Conto corrente
// Scrivere un programma che legge da riga di comando un intero che rappresenta il saldo di un
// conto corrente. Il programma legge poi da standard input una serie di numeri interi che rappre-
// sentano spese da addebitare sul conto e stampa il saldo finale. La lettura si interrompe quando il
// saldo è <=0.

package main

import "fmt"

func main() {
	var balance int
	fmt.Print("Provide the initial balance: ")
	fmt.Scan(&balance)
	var expense int
	for balance > 0 {
		fmt.Print("Provide an expense: ")
		fmt.Scan(&expense)
		balance -= expense
	}
	fmt.Println("Final balance: ", balance)
	
}