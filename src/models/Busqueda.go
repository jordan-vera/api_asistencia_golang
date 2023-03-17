package models

type Busqueda struct {
	Idbusqueda      int    `json:"idbusqueda"`
	Identificacion1 string `json:"identificacion1"`
	Identificacion2 string `json:"identificacion2"`
	Fechahora       string `json:"fechahora"`
}
