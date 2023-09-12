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