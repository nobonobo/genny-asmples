package fold

import (
	"fmt"
)

func ExampleFoldInt() {
	sum := FoldInt([]int{1, 2, 3, 4}, func(a, b int) int {
		return a + b
	})
	fmt.Println(sum)
	//output:
	//10
}

func ExampleFoldFloat64() {
	sum := FoldFloat64([]float64{1.1, 2.1, 3.1, 4.1}, func(a, b float64) float64 {
		return a + b
	})
	fmt.Println(sum)
	//output:
	//10.4
}

func ExampleFoldString() {
	sum := FoldString([]string{"1", "2", "3", "4"}, func(a, b string) string {
		return a + b
	})
	fmt.Println(sum)
	//output:
	//1234
}

func ExampleFoldBytes() {
	sum := FoldBytes([]Bytes{
		Bytes("1"),
		Bytes("2"),
		Bytes("3"),
		Bytes("4"),
	}, func(a, b Bytes) Bytes {
		return append(a, b...)
	})
	fmt.Println(sum)
	//output:
	//[49 50 51 52]
}

type user string

func (u user) Name() string {
	return string(u)
}

func ExampleFoldUser() {
	sum := FoldUser([]User{
		user("a"),
		user("b"),
		user("c"),
		user("d"),
	}, func(a, b User) User {
		return user(a.Name() + b.Name())
	})
	fmt.Println(sum)
	//output:
	//abcd
}
