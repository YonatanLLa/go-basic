package main

// admite solo un paque
import (
	"go/src/variables"
	"fmt"
	"reflect"
)

// function principal de go
func main()  {
	estado,	texto := variables.ConviertoText(123)
	tipe := reflect.TypeOf(estado)

	fmt.Println(estado)
	fmt.Println(texto)
	fmt.Println(tipe)
}
