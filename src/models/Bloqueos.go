package models

type Bloqueos struct {
	Idbloqueo      int    `json:"idbloqueo"`
	Identificacion string `json:"identificacion"`
	Anio           int    `json:"anio"`
	Mes            int    `json:"mes"`
	Dia            int    `json:"dia"`
	Estado         int    `json:"estado"`
	Hora           string `json:"hora"`
	Autorizador    string `json:"autorizador"`
}
