package ejercicios

import (
	"strconv"
	"strings"
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
// func NumerosPares(num int) int {
// 	suma := 0
// 	for i := 2; i <= num; i += 2 {
// 		suma += i
// 	}
// 	return suma
// }

// Factorial: Escribe una función que calcule el factorial de un número entero no negativo.
// 3! = 3*2*1
// 4! = 4*3*2*1

func Factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * Factorial(number-1)
}

// Contador de palabras: Escribe una función que tome una cadena de texto como entrada y cuente cuántas palabras contiene. Puedes asumir que las palabras están separadas por espacios en blanco.

func ContadorDePalabras(palabra string) []string {
	palabras := strings.Fields(palabra)

	return palabras

}

// Primos: Escribe una función que determine si un número dado es primo o no.

// number/number => 1
// number/1 => number

func Primo(number int) bool {
	if number <= 1 {
		return false
	}

	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// Inversión de cadena: Escribe una función que invierta una cadena de texto. Por ejemplo, si la entrada es "Hola", la salida debería ser "aloH".
