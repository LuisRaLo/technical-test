package models

type Response struct {
	Folio   string `json:"folio,omitempty" example:"" format:"string"`
	Message string `json:"mensaje" example:"Operación realizada con éxito" format:"string"`
}

type ResponseWithResult struct {
	*Response
	Result interface{} `json:"resultado"`
}

type Response400WithResult struct {
	Folio   string   `json:"folio,omitempty" example:"" format:"string"`
	Message string   `json:"mensaje" example:"Operación realizada sin éxito" format:"string"`
	Details []string `json:"detalles" example:"Error en la petición" format:"string"`
}
type Response401WithResult struct {
	Folio   string   `json:"folio,omitempty" example:"" format:"string"`
	Message string   `json:"mensaje" example:"Operación realizada sin éxito" format:"string"`
	Details []string `json:"detalles" example:"No se encontró el recurso solicitado" format:"string"`
}
type Response403WithResult struct {
	Folio   string `json:"folio,omitempty" example:"" format:"string"`
	Message string `json:"mensaje" example:"Operación realizada sin éxito" format:"string"`
}
type Response404WithResult struct {
	Folio   string   `json:"folio,omitempty" example:"" format:"string"`
	Message string   `json:"mensaje" example:"Operación realizada con éxito" format:"string"`
	Details []string `json:"detalles" example:"No se encontró el recurso solicitado" format:"string"`
}
type Response500WithResult struct {
	Folio   string   `json:"folio,omitempty" example:"" format:"string"`
	Message string   `json:"mensaje" example:"Operación fallida" format:"string"`
	Details []string `json:"detalles" example:"Error interno del servidor" format:"string"`
}

type Response409WithResult struct {
	Folio   string `json:"folio,omitempty" example:"" format:"string"`
	Message string `json:"mensaje" example:"Operación realizada sin éxito" format:"string"`
}
