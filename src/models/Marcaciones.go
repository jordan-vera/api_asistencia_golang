package models

type Marcaciones struct {
	IDMARCACION  int    `json:"IDMARCACION"`
	IDASISTENCIA int    `json:"IDASISTENCIA"`
	HORA         string `json:"HORA"`
	TIPO         string `json:"TIPO"`
	IDSUCURSAL   int    `json:"IDSUCURSAL"`
	IMAGEN       string `json:"IMAGEN"`
	FILE         string `json:"FILE"`
}
