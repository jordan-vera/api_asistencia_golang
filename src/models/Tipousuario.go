package models

type Tipousuario struct {
	Idtipousuario int    `json:"idtipousuario"`
	Tipouser      string `json:"tipouser"`
	Estado        int    `json:"estado"`
}
