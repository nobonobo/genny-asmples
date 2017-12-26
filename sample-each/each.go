package each

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string ValueType=string,int,bool"

import "github.com/cheekybits/genny/generic"

type KeyType generic.Type
type ValueType generic.Type
type KeyTypeValueTypeMap map[KeyType]ValueType

func (m KeyTypeValueTypeMap) Each(aply func(key KeyType, value ValueType) bool) {
	for k, v := range m {
		if !aply(k, v) {
			break
		}
	}
}
