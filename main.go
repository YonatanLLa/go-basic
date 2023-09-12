package main

// admite solo un paque
import (
	"fmt"
	// "go/src/variables"
	"go/src/ejercicios"
	// "reflect"
	// "runtime"
)

// function principal de go
func main() {
	// estado, texto := variables.ConviertoText(123)
	// tipe := reflect.TypeOf(estado)

	// fmt.Println(estado)
	// fmt.Println(texto)
	// fmt.Println(tipe)

	

	// if os := runtime.GOOS; os == "Linux." || os == "OS X." {
	// 	fmt.Println("Esto no es windows")
	// } else {
	// 	fmt.Println("Esto es windows")
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es un Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n ", os)
	// }

	// result, dato := ejercicios.Ejerccio("fff")
	// fmt.Println(result, dato)

	// fmt.Println(dato)

	// ejerccio2:
	num := ejercicios.NumerosPares(12)
	fmt.Println(num)

	
}
