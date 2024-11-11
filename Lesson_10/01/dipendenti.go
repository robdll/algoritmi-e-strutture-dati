// Anna è supervisore di Bruno, Carlo e Daniela. Bruno è supervisore di Enrico e
// Francesco. Gianni è supervisore di Harry. Francesco è supervisore di Irene. Il numero di
// dipendenti senza subordinati è 5 (Carlo, Daniela, Enrico, Harry, Irene). La lista di dipendenti
// che si trovano sopra Irene è: Francesco, Bruno, Anna. Questo è l’elenco dei dipendenti in ordine
// di livello: A, G (massimo livello), B, D, H, C (subordinati di dipendenti di massimo livello, non
// importa il loro ordine), F, E, I.

package main

import (
	"fmt"
)

type Employee struct {
	name string
	subordinates []*Employee
	manager *Employee
}

type Company struct {
	employees []*Employee
}

func main () {
	var anna = Employee{name: "Anna"}
	var gianni = Employee{name: "Gianni"}
	var bruno = Employee{name: "Bruno", manager: &anna}
	var carlo = Employee{name: "Carlo", manager: &anna}
	var daniela = Employee{name: "Daniela", manager: &anna}
	var enrico = Employee{name: "Enrico", manager: &bruno}
	var francesco = Employee{name: "Francesco", manager: &bruno}
	var harry = Employee{name: "Harry", manager: &gianni}
	var irene = Employee{name: "Irene", manager: &francesco}

	anna.subordinates = []*Employee{&bruno, &carlo, &daniela}
	bruno.subordinates = []*Employee{&enrico, &francesco}
	gianni.subordinates = []*Employee{&harry}
	francesco.subordinates = []*Employee{&irene}

	var company = Company{
		employees: []*Employee{
			&anna,
			&gianni,
			&bruno,
			&carlo,
			&daniela,
			&enrico,
			&francesco,
			&harry,
			&irene,
		},
	}
	fmt.Println("##################")
	stampaSubordinati(anna)
	quantiSenzaSubordinati(anna)
	fmt.Println("##################")
	stampaSubordinati(gianni)
	quantiSenzaSubordinati(gianni)
	fmt.Println("##################")
	stampaImpiegatiSopra(harry)
	fmt.Println("##################")
	stampaImpiegatiConSupervisore(company)
}

func stampaSubordinati (d Employee) {
	for _, s := range d.subordinates {
		fmt.Println("-", s.name)
	}
}

func quantiSenzaSubordinati (d Employee) {
	fmt.Print("Nel Team di ", d.name, " ci sono ")
	count := 0
	for _, s := range d.subordinates {
		if len(s.subordinates) == 0 {
			count++
		}
	}
	fmt.Println(count, "dipendenti senza subordinati")
}

func supervisore (d Employee) {
	fmt.Println("Supervisore di ", d.name, " è ", d.manager.name)
}

func stampaImpiegatiSopra (d Employee) {
	fmt.Println("I dipendenti sopra ", d.name, " sono: ")
	for d.manager != nil {
		fmt.Print(d.manager.name, " ")
		d = *d.manager
	}
	fmt.Println()
}

func stampaImpiegatiConSupervisore (c Company) {
	for _, s := range c.employees {
		fmt.Print(s.name, " ")
		if(s.manager != nil) {
			fmt.Print("- ",s.manager.name)
		}
		fmt.Println()
	}
}