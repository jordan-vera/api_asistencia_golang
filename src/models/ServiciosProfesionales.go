package models

type ServiciosProfesionales struct {
	Idservicio     int    `json:"idservicio"`
	Nombres        string `json:"nombres"`
	Usuario        string `json:"usuario"`
	Clave          string `json:"clave"`
	Identificacion string `json:"identificacion"`
	Idsucursal     int    `json:"idsucursal"`
	Estado         int    `json:"estado"`
}

type ServiciosProfesionalesAsistencia struct {
	Idservicio     int    `json:"idservicio"`
	Nombres        string `json:"nombres"`
	Usuario        string `json:"usuario"`
	Clave          string `json:"clave"`
	Identificacion string `json:"identificacion"`
	Idsucursal     int    `json:"idsucursal"`
	Estado         int    `json:"estado"`
	Asistencias    []AsistenciasMarcaciones
}
