package ejercicios

import (
	"strconv"
)

func Ejerccio(dato string) (int, string) {
	// Convertir ese string a entero
	result, err := strconv.Atoi(dato)

	if err != nil {
		return 0, "Hubo un error"
	}

	if result > 100 {
		return result, "Es mayor que 100"
	}

	return result, "No es mayor a 100"
}

/* Suma de números pares: Escribe una función que tome un número entero positivo como entrada y devuelva la suma de todos los números pares desde 1 hasta ese número.*/
func NumerosPares(num int) int {
	suma := 0
	for i := 2; i <= num; i += 2 {
		suma += i
	}
	return suma
}

// Factorial: Escribe una función que calcule el factorial de un número entero no negativo.


