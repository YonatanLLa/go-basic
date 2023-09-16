package arreglos_slices

import (
	"fmt"
)

// var tablaSlice []int = []int{1,2,3}
var arreglo [10]int = [10]int{1,2,3,4,5,6,7}
func MuestroSlice()  {
	// fmt.Println(tablaSlice)

	porsion := arreglo[3:]
	porsion2 := arreglo[:5]
	porsion3 := arreglo[6:7]

	fmt.Println(porsion)
	fmt.Println(porsion2)
	fmt.Println(porsion3)
}

func Capacidad()  {
	// elementos := make([]int, 5, 20)

	// fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0)

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("Largo %d, Capacidad %d", len(nums), cap(nums))
}