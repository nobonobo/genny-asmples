/*
orderedmap package
Copyright (c) 2017 Noboru Irieda
This software is released under the MIT License.
http://opensource.org/licenses/mit-license.php
*/

package orderedmap

// github.com/cheekybits/genny 参照
//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "Type=Sample"

import (
	"github.com/cheekybits/genny/generic"
)

// Type ...
type Type generic.Type

// TypeItem ...
type TypeItem struct {
	index int
	value *Type
}

// TypeMap ...
type TypeMap struct {
	keys  []string
	items map[string]*TypeItem
}

// NewTypeMap ...
func NewTypeMap() *TypeMap {
	return &TypeMap{
		keys:  []string{},
		items: map[string]*TypeItem{},
	}
}

func (m *TypeMap) Len() int { return len(m.keys) }
func (m *TypeMap) Swap(i, j int) {
	iv := m.items[m.keys[i]]
	jv := m.items[m.keys[j]]
	jv.index, iv.index = iv.index, jv.index
	m.keys[i], m.keys[j] = m.keys[j], m.keys[i]
}
func (m *TypeMap) Less(i, j int) bool { return m.keys[i] < m.keys[j] }

// Get ...
func (m *TypeMap) Get(key string) *Type {
	item, ok := m.items[key]
	if ok {
		return item.value
	}
	return nil
}

// Set ...
func (m *TypeMap) Set(key string, value *Type) {
	m.items[key] = &TypeItem{len(m.keys), value}
	m.keys = append(m.keys, key)
}

// Del ...
func (m *TypeMap) Del(key string) {
	item, ok := m.items[key]
	if ok {
		m.keys = append(m.keys[:item.index], m.keys[item.index+1:]...)
		delete(m.items, key)
	}
}

// Iter ...
func (m *TypeMap) Iter(f func(key string, value *Type) bool) {
	for _, key := range m.keys {
		TypeItem, ok := m.items[key]
		if ok {
			if !f(key, TypeItem.value) {
				break
			}
		}
	}
}
