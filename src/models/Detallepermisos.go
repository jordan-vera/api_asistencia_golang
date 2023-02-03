package models

type Detallepermisos struct {
	Iddetallepermiso int `json:"iddetallepermiso"`
	Idpermiso        int `json:"idpermiso"`
	Numerodia        int `json:"numerodia"`
	Mes              int `json:"mes"`
	Anio             int `json:"anio"`
}
