package global

import (
	"fmt"
	"math"
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

func CalcularHora(hora string) int {
	concatFecha := fmt.Sprintf("%s %s", FechaActual(), hora)
	currentTime := time.Now()
	loc := currentTime.Location()
	layout := "2006-01-02 15:04"

	pasttime, err := time.ParseInLocation(layout, concatFecha, loc)
	if err != nil {
		fmt.Println(err)
	}

	diff := pasttime.Sub(currentTime)

	return int(Roundf(diff.Minutes()))
}

func Roundf(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

func EsPositivoNeutro(numero int) bool {
	if numero == 0 {
		return true
	} else if numero > 0 {
		return true
	} else {
		return false
	}
}

func NombreDia() string {
	weekday := time.Now().Weekday()
	if int(weekday) == 1 {
		return "lunes"
	} else if int(weekday) == 2 {
		return "martes"
	} else if int(weekday) == 3 {
		return "miércoles"
	} else if int(weekday) == 4 {
		return "jueves"
	} else if int(weekday) == 5 {
		return "viernes"
	} else if int(weekday) == 6 {
		return "sábado"
	} else {
		return "domingo"
	}
}
