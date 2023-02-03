package models

type Sucursales struct {
	IDSUCURSAL int    `json:"IDSUCURSAL"`
	SUCURSAL   string `json:"SUCURSAL"`
	LONGITUD   string `json:"LONGITUD"`
	LATITUD    string `json:"LATITUD"`
}
