package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numero1 int
var err error

func TableNum() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese un Numero: ")

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Ek dato ingresado es incorrecto")
		}
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(numero1, "*", i, "=", (numero1 * i))
	}

}

func TableMult() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Tabla de multiplicar")
	fmt.Print("Ingrese un numero: ")

	var texto string

	for {
		if scanner.Scan() {
			input := scanner.Text()
			numero1, err = strconv.Atoi(input)
		}
		if err != nil {
			fmt.Println("Error:  El dato ingresado no es un nummero valido. Intentalo de nuevo")
			continue
		}

		fmt.Println("tabla de multiplicar de", numero1, ":")
		for i := 0; i < 10; i++ {
			texto += fmt.Sprintf("%d * %d = %d\n", numero1, i, numero1*i)

		}

		return texto
	}

}

func TableFinal() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Tabla de multiplicar")
	fmt.Print("Ingrese un numero: ")

	for {
		if scanner.Scan() {
			input := scanner.Text()
			if input == "exit" {
				fmt.Println("Saliendo del programa")
				return
			}

			numero1, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Error: El dato ingresado no es un número valido, Inténtelo de nuevo")
				continue
			}
			fmt.Println("Especifique el rango de la tabla (Ejemplo: 5 - 10): ")
			if scanner.Scan() {
				rangeInput := scanner.Text()
				start, end, err := parseRangeInput(rangeInput)

				if err != nil {
					fmt.Println("Error: Rango no válido. Inténtelo de nuevo.")
				}

				fmt.Printf("Tabla de multiplicar de %d en el ranfo %d-%d:\n", numero1, start, end)
				for i := start; i < end; i++ {
					fmt.Printf("%d * %d = %d\n", numero1, i, numero1*i)
				}
				break
			}
		}
	}
}

func parseRangeInput(input string) (int, int, error) {
	parts := strings.Split(input, "-")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("Entrada no vlaida")
	}

	// convertir las partes a números enteros.
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Inicio del ranod no válido")
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Fin del ranog no valido")
	}

	// validar que el inicio sea menor o igual al final.
	if start > end {
		return 0, 0, fmt.Errorf("El inicio del ranfo debe ser menor o iual al final")
	}
	return start, end, nil
}
