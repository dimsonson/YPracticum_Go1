package main

import (
	"fmt"

	"github.com/dimsonson/YPracticum_Go1/S3T2L1_interface_concept/company"
	"github.com/dimsonson/YPracticum_Go1/S3T2L1_interface_concept/person"
	//"company"
	//	"person"
)

func main() {
	pers := person.Person{}
	comp := company.Company{}

	// мы передаём переменную типа Person в функцию,
	//аргументом которой является переменная Worker!
	comp.Hire(pers)
	fmt.Println(pers)
	fmt.Println(comp)
	
}
