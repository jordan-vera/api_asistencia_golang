package models

type Vacaciones struct {
	Idvacaciones   int    `json:"idvacaciones"`
	Identificacion string `json:"identificacion"`
	Cantidaddias   int    `json:"cantidaddias"`
	Fechainicio    string `json:"fechainicio"`
	Fechafin       string `json:"fechafin"`
	Estado         string `json:"estado"`
	Anio           int    `json:"anio"`
	Nombre         string `json:"nombre"`
}

type VacacionesDetalleFilter struct {
	Idvacaciones   int    `json:"idvacaciones"`
	Identificacion string `json:"identificacion"`
	Cantidaddias   int    `json:"cantidaddias"`
	Fechainicio    string `json:"fechainicio"`
	Fechafin       string `json:"fechafin"`
	Estado         string `json:"estado"`
	Anio           int    `json:"anio"`
	Detalle        []VacacionesDetalle
}
