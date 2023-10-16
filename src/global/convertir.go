package global

import (
	"log"
	"strconv"
)

func ConvertirAhInt(numero string) int {
	entero, err := strconv.Atoi(numero)
	if err != nil {
		log.Fatal(err)
	}
	return entero
}
