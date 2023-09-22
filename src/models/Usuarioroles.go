package models

type UsuarioRoles struct {
	Idusuarioroles int    `json:"idusuarioroles"`
	Identificacion string `json:"identificacion"`
	Nombres        string `json:"nombres"`
	Idtipousuario  int    `json:"idtipousuario"`
}
