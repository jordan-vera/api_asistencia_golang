package models

type Trabajocampo struct {
	Idcampo        int    `json:"idcampo"`
	Identificacion string `json:"identificacion"`
	Fecha          string `json:"fecha"`
	Comentario     string `json:"comentario"`
	Dia            int    `json:"dia"`
	Mes            int    `json:"mes"`
	Anio           int    `json:"anio"`
}
