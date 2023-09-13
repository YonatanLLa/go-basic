package juegos

import (
	"fmt"
	"math/rand"
	"time"
)

func Juegos()  {
	rand.Seed(time.Now().UnixNano()) //Inicializamos la semilla para los numeros aleatorios

	numeroSecreto := rand.Intn(100) 
	intentos := 0

	fmt.Println("Bienvenido a Adivina el número!")
	fmt.Println("Estoy pensando en un número entre 0 y 99. !Adivina!")

	for {
		var guess int 
		fmt.Print("Tu suposición: ")	
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Por favor, ingresa un numero valido.")
			continue
		}

		intentos++

		if guess < numeroSecreto {
			fmt.Println("Mas alto...")
		} else if guess > numeroSecreto {
			fmt.Println("Mas bajo...")
		} else {
			fmt.Printf("¡Felicitaciones! Adivinaste el nuemero %d en %d intentos.", numeroSecreto, intentos)
		}

	}
}