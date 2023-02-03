package models

type Empleados struct {
	Identificacion    string `json:"identificacion"`
	NombreUnido       string `json:"nombreUnido"`
	SecuencialOficina string `json:"secuencialOficina"`
}

type EmpleadosAsistencias struct {
	Identificacion string `json:"identificacion"`
	NombreUnido    string `json:"nombreUnido"`
	Asistencias    []AsistenciasMarcaciones
}

type Usuario struct {
	Codigouser        string `json:"codigo"`
	Nombre            string `json:"nombre"`
	Secuencialoficina int    `json:"secuencialOficina"`
	Estadoactivo      bool   `json:"estaActivo"`
	Numeroverificador int    `json:"numeroVerificador"`
	Identificacion    string `json:"identificacion"`
}
