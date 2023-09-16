package arreglos_slices

import "fmt"

var tabla [10]int = [10]int{10, 1, 59}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54

	tabla2 := [10]string{"pablo", "juan"}

	fmt.Println(tabla2)
	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
