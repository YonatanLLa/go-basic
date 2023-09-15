package funcs

import (
	"fmt"
)

func Calculos()  {

	suma := func (num1 int, num2 int) int  {
		return num1 + num2
	}
	fmt.Println(suma(12,13))
}