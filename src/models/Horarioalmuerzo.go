package models

type Horarioalmuerzo struct {
	Idhorarioalmuerzo int    `json:"idhorarioalmuerzo"`
	Entrada           string `json:"entrada"`
	Salida            string `json:"salida"`
	Salidaacasa       string `json:"salida_a_casa"`
	Identificacion    string `json:"identificacion"`
}

type HorarioalmuerzoColaborador struct {
	Idhorarioalmuerzo int    `json:"idhorarioalmuerzo"`
	Entrada           string `json:"entrada"`
	Salida            string `json:"salida"`
	Salidaacasa       string `json:"salida_a_casa"`
	Identificacion    string `json:"identificacion"`
	Nombres           string `json:"nombres"`
}
