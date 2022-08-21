package main

import (
    "company"
    "person"
)

func main() {
    pers := person.Person{}
    comp := company.Company{}
    
	// мы передаём переменную типа Person в функцию, 
	//аргументом которой является переменная Worker! 
    comp.Hire(pers) 

}
 