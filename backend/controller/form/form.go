package form

import "Dotato-di-una-libreria/backend/model"

// Form ...
type Form interface {
	ParseToDto() model.Dto
}
