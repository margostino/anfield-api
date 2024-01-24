package db

import (
	"github.com/margostino/anfield-api/fpl"
)

func loadTypes(fplElementTypes []*fpl.ElementType) map[int]*ElementType {
	elementTypes := make(map[int]*ElementType)
	for _, elementType := range fplElementTypes {
		elementTypes[elementType.ID] = &ElementType{
			ID:                elementType.ID,
			PluralName:        elementType.PluralName,
			PluralNameShort:   elementType.PluralNameShort,
			SingularName:      elementType.SingularName,
			SingularNameShort: elementType.SingularNameShort,
		}
	}
	return elementTypes
}
