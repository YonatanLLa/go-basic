package funcs

import (
	"fmt"
)
func Exponencia(valor int)  {
	if valor > 100000 {
		return
	}
	fmt.Println(valor)
	Exponencia(valor*2)
}