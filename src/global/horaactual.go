package global

import (
	"fmt"
	"time"
)

func FechaActual() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d",
		t.Year(), t.Month(), t.Day())

	return fecha
}

func HoraActual() string {
	t := time.Now()
	fecha := fmt.Sprintf("%02d:%02d:%02d",
		t.Hour(), t.Minute(), t.Second())

	return fecha
}

func NumAnioActual() int {
	t := time.Now()
	return t.Year()
}

func NumMesActual() int {
	t := time.Now()
	return int(t.Month())
}

func NumDiaActual() int {
	t := time.Now()
	return int(t.Day())
}

func NumHoraActual() int {
	t := time.Now()
	return t.Hour()
}

func NumMinutoActual() int {
	t := time.Now()
	return t.Minute()
}
