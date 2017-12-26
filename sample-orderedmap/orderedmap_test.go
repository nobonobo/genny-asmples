package orderedmap

import (
	"fmt"
	"sort"
)

func ExampleSampleMap() {
	m := NewSampleMap()
	m.Set("hoge1", &Sample{})
	m.Set("moge1", &Sample{})
	m.Set("hoge2", &Sample{})
	m.Set("moge2", &Sample{})
	sort.Sort(m)
	m.Del("moge1")
	m.Iter(func(k string, v *Sample) bool {
		fmt.Println(k, v)
		return true
	})
	// output:
	// hoge1 &{}
	// hoge2 &{}
	// moge2 &{}
}
