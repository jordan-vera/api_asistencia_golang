package models

type Bloqueofuerarango struct {
	Idbloqueofuerarando int    `json:"idbloqueofuerarando"`
	Identificacion      string `json:"identificación"`
	Tipobloqueo         string `json:"tipobloqueo"`
	Dia                 int    `json:"dia"`
	Mes                 int    `json:"mes"`
	Anio                int    `json:"anio"`
	Horab               string `json:"horab"`
	Estado              int    `json:"estado"`
	Justificacion       string `json:"justificacion"`
	Autorizadorb        string `json:"autorizadorb"`
}
