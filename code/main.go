package main

import (
    "fmt"
)

type Analyst struct{
    nombre string
    work_station int
}

func main(){

    h := Analyst{"Humberto", 62}

    fmt.Println(h)
}