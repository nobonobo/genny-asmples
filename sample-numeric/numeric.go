package numeric

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Type=int,int16,int32,int64"

import (
	"github.com/cheekybits/genny/generic"
)

type Type generic.Number

func SumType(values ...Type) Type {
	res := values[0]
	for _, v := range values[1:] {
		res += v
	}
	return res
}

func MaxType(values ...Type) Type {
	res := values[0]
	for _, v := range values[1:] {
		if res < v {
			res = v
		}
	}
	return res
}

func MinType(values ...Type) Type {
	res := values[0]
	for _, v := range values[1:] {
		if res > v {
			res = v
		}
	}
	return res
}
