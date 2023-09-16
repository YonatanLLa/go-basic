package ejer_interfaces

import (
	"fmt"
	"go/src/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("soy un/ %s, estoy respirando \n", hu.Sexo())
}
