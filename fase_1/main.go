package main

import (
	"fmt"
	"os"
	"strconv"
	//"errors"
)

func main(){
	
	// Ingreso de datos
	fmt.Println("Caluladora")
	var firstNum string
	fmt.Scan(&firstNum)

	// Conversión de datos
	num1, err := strconv.ParseFloat(firstNum, 32)
	if err != nil{
		fmt.Println("Deberás ingresar un número")
		os.Exit(1)
	}
	
	// Resultado
	fmt.Printf("%T\n",num1)
}