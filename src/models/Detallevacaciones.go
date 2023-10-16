package models

type VacacionesDetalle struct {
	Iddetallevacaciones int `json:"iddetallevacaciones"`
	Idvacaciones        int `json:"idvacaciones"`
	Numerodia           int `json:"numerodia"`
	Mes                 int `json:"mes"`
	Anio                int `json:"anio"`
}

type VacacionesDetalleMesAnio struct {
	Idvacaciones   int    `json:"idvacaciones"`
	Identificacion string `json:"identificacion"`
	Cantidaddias   int    `json:"cantidaddias"`
	Estado         string `json:"estado"`
	Numerodia      int    `json:"numerodia"`
	Mes            int    `json:"mes"`
	Anio           int    `json:"anio"`
}
