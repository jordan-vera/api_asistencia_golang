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
	Horafinpermiso        string `json:"horafinpermiso"`
	Anio                  int    `json:"anio"`
	Iddetallepermiso      int    `json:"iddetallepermiso"`
}

type Tipopermisos struct {
	Idtipopermiso int    `json:"idtipopermiso"`
	Tipo          string `json:"tipo"`
}

type PermisosAnioMesDia struct {
	Numerodia      int    `json:"numerodia"`
	Mes            int    `json:"mes"`
	Anio           int    `json:"anio"`
	Identificacion string `json:"identificacion"`
}

type DetallepermisosVerificar struct {
	Iddetallepermiso int `json:"iddetallepermiso"`
}
