package funcs

import (
	"fmt"
)

func tabla(valor int) func() int  {
	numero := valor
	secuencia := 0
	
	return func() int {
		secuencia++
		return numero*secuencia
	}
}

func LlamarClousure()  {
	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 0; i < 10; i++ {
		fmt.Println(MiTabla())
	}
}