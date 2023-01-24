package models

type Empleados struct {
	Identificacion string `json:"identificacion"`
	NombreUnido    string `json:"nombreUnido"`
}

type Usuario struct {
	Codigouser        string `json:"codigo"`
	Nombre            string `json:"nombre"`
	Secuencialoficina int    `json:"secuencialOficina"`
	Estadoactivo      bool   `json:"estaActivo"`
	Numeroverificador int    `json:"numeroVerificador"`
	Identificacion    string `json:"identificacion"`
}
