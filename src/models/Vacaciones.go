package models

type Vacaciones struct {
	Idvacaciones   int    `json:"idvacaciones"`
	Identificacion string `json:"identificacion"`
	Cantidaddias   int    `json:"cantidaddias"`
	Estado         string `json:"estado"`
	Anio           int    `json:"anio"`
	Nombre         string `json:"nombre"`
}

type VacacionesDetalleFilter struct {
	Idvacaciones   int    `json:"idvacaciones"`
	Identificacion string `json:"identificacion"`
	Cantidaddias   int    `json:"cantidaddias"`
	Estado         string `json:"estado"`
	Anio           int    `json:"anio"`
	Detalle        []VacacionesDetalle
}
