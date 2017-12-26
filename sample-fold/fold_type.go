package fold

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Type=string,Bytes,User"

import (
	"github.com/cheekybits/genny/generic"
)

type Type generic.Type

func FoldType(l []Type, apply func(a Type, b Type) Type) (res Type) {
	res = l[0]
	for _, v := range l[1:] {
		res = apply(res, v)
	}
	return res
}
