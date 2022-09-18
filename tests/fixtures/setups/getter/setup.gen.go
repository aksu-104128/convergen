// Code generated by github.com/reedom/convergen
// DO NOT EDIT.

package getter

import (
	"github.com/reedom/convergen/tests/fixtures/data/ddd/domain"
	"github.com/reedom/convergen/tests/fixtures/data/ddd/model"
)

// DomainToModel copies domain.Pet to model.Pet.
func DomainToModel(pet *domain.Pet) (dst *model.Pet) {
	dst = &model.Pet{}
	dst.ID = pet.ID()
	// no match: dst.Category
	dst.Name = pet.Name()
	// skip: dst.PhotoUrls
	// no match: dst.Status

	return
}
