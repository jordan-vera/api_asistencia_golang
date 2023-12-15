package models

type Permisos struct {
	Idpermiso             int    `json:"idpermiso"`
	Idtipopermiso         int    `json:"idtipopermiso"`
	Identificacion        string `json:"identificacion"`
	Desde                 string `json:"desde"`
	Hasta                 string `json:"hasta"`
	Motivo                string `json:"motivo"`
	Estadojefe            string `json:"estadojefe"`
	Fechasolicitud        string `json:"fechasolicitud"`
	Tiempoestimado        string `json:"tiempoestimado"`
	Tipo                  string `json:"tipo"`
	Numerodia             int    `json:"numerodia"`
	Autorizador           string `json:"autorizador"`
	Calculadoenvacaciones int    `json:"calculadoenvacaciones"`
	Escargovacaciones     int    `json:"escargovacaciones"`
	Horainiciopermiso     string `json:"horainiciopermiso"`
}

type Tipopermisos struct {
	Idtipopermiso int    `json:"idtipopermiso"`
	Tipo          string `json:"tipo"`
}
