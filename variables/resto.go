package variables

import (
	"fmt"
	"time"
)
// todo los archivos pesados lo importamos de manera implicita
var Nombe string 
var Estado bool
var Sueldo1 float32
var Sueldo2 float64
var Fecha time.Time

func RestoVariables()  {
	Nombe = "Pedro"
	Estado = true
	Sueldo1 = 155.77
	Fecha = time.Now()

	fmt.Println(Nombe)
	fmt.Println(Estado)
	fmt.Println(Sueldo1)
	fmt.Println(Fecha)
}

func ConviertoText(numero int)(bool, string)  {
	// vat texto string
	texto = number

}