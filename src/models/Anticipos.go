package models

type Anticipos struct {
	Idanticipo           int    `json:"idanticipo"`
	Fecha                string `json:"fecha"`
	Identificacion       string `json:"identificacion"`
	Cantidadanticipo     string `json:"cantidadanticipo"`
	Motivo_si_es_segundo string `json:"motivo_si_es_segundo"`
	Meses_a_deducir      int    `json:"meses_a_deducir"`
	Anio                 int    `json:"anio"`
	Mes                  int    `json:"mes"`
	Dia                  int    `json:"dia"`
	Estodogerente        string `json:"estodogerente"`
}
