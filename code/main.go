package main

import (
    "fmt"
)

type Analyst struct{
    nombre string
    work_station int
    edad int
}

func main(){

    // Entrada de datos simulada
    h := Analyst{"Humberto", 62, 24}

    // Métodos aplicado a la instancia
    h.saludar()
    h.birthDay()
}

func(a *Analyst) saludar(){
    fmt.Printf("Hola, soy %s y mi lugar es el #%d\n", a.nombre, a.work_station)
}

func (a *Analyst) birthDay(){
    a.edad += 1
    fmt.Printf("Felicidades %s, cumpliste %d años\n", a.nombre, a.edad)
}