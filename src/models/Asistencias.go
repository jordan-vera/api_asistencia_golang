package models

type Asistencias struct {
	IDASISTENCIA      int    `json:"IDASISTENCIA"`
	IDENTIFICACION    string `json:"IDENTIFICACION"`
	FECHA             string `json:"FECHA"`
	MES               int    `json:"MES"`
	ANIO              int    `json:"ANIO"`
	DIA               int    `json:"DIA"`
	NOMBREDIA         string `json:"NOMBREDIA"`
	JUSTIFICACION     string `json:"JUSTIFICACION"`
	HORASJUSTIFICADAS int    `json:"HORASJUSTIFICADAS"`
}

type AsistenciasMarcaciones struct {
	IDASISTENCIA      int    `json:"IDASISTENCIA"`
	IDENTIFICACION    string `json:"IDENTIFICACION"`
	FECHA             string `json:"FECHA"`
	MES               int    `json:"MES"`
	ANIO              int    `json:"ANIO"`
	DIA               int    `json:"DIA"`
	NOMBREDIA         string `json:"NOMBREDIA"`
	JUSTIFICACION     string `json:"JUSTIFICACION"`
	HORASJUSTIFICADAS int    `json:"HORASJUSTIFICADAS"`
	MARCACIONES       []Marcaciones
}
