package models

type Seguidor struct {
	UsuarioID  uint64 `json:"usuarioId,omitempty"`
	SeguidorID uint64 `json:"seguidorID,omitempty"`
}
