package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"gqltest/item"
	"io"
	"strconv"
)

func MarshalItemType(s item.Type) graphql.Marshaler {
	var mapValueToString = map[item.Type]string{
		item.TypeA: "ItemTypeA",
		item.TypeB: "ItemTypeB",
	}
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(mapValueToString[s]))
	})
}

func UnmarshalItemType(v interface{}) (item.Type, error) {
	if tmp, ok := v.(string); ok {
		var mapStringToValue = map[string]item.Type{
			"ItemTypeA": item.TypeA,
			"ItemTypeB": item.TypeB,
		}

		if tmp2, ok := mapStringToValue[tmp]; ok {
			return tmp2, nil
		}
	}
	return "", fmt.Errorf("invalid SortBy value")
}
