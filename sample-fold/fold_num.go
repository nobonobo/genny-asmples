package fold

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Num=float64,int"

import (
	"github.com/cheekybits/genny/generic"
)

type Num generic.Number

func FoldNum(l []Num, apply func(a Num, b Num) Num) (res Num) {
	res = l[0]
	for _, v := range l[1:] {
		res = apply(res, v)
	}
	return res
}
