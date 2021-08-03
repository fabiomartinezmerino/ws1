package main

import (
	"fmt"
)

func main() {

	/*nil map*/

	var nilMap map[string][]string
	fmt.Printf("Map - nil: %v\n", nilMap)

	/*map examples*/

	stringEmptyMap := map[string][]string{"Real Madrid": []string{"Curtoi", "Sergio Ramos", "Hazard"}}
	fmt.Printf("Map - empty:%v\n", stringEmptyMap)

	stringMakemap := make(map[string][]string, 20)
	fmt.Printf("Map created with make(): %v\n Length of this map %d\n", stringMakemap, len(stringMakemap))
	stringMakemap["Barsa"] = []string{"Portero", "Defensa", "Medio", "Delantero"}
	fmt.Printf("Map created with make() and now initialized qith a literal: %v\n Length of this map %d\n", stringMakemap, len(stringMakemap))

	/*map reading and insert*/

	totalWins := map[string]int{} /*map: key --> string value --> int. Initialized using an empty literal*/

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 3
	totalWins["Kittens"]++

	fmt.Printf("Wins of Orcas = %d\n", totalWins["Orcas"])
	fmt.Printf("Wins of Lions = %d\n", totalWins["Lions"])
	fmt.Printf("Wins of Kittens= %d\n", totalWins["Kittens"])

}
