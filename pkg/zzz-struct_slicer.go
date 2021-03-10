// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package pkg

import (
	"reflect"
	"sort"

	"github.com/ikasamt/goslicer/typeslicer"
)

type Channels []Channel

// First is ..
func (anythings Channels) First() (Channel, bool) {
	if len(anythings) == 0 {
		return Channel{}, false
	}
	return anythings[0], true
}

// Where is ..
func (anythings Channels) Where(f func(anything Channel) bool) (result Channels) {
	for _, a := range anythings {
		if f(a) {
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings Channels) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings Channels) CountIf(f func(anything Channel) bool) (counter int) {
	for _, a := range anythings {
		if f(a) {
			counter++
		}
	}
	return
}

// Select is ..
func (anythings Channels) Select(fieldName string) (result typeslicer.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings Channels) SortBy(sortFunc func(Channel, Channel) bool) (result Channels) {
	f := func(i, j int) bool {
		a := anythings[i]
		b := anythings[j]
		return sortFunc(a, b)
	}

	tmp := make(Channels, len(anythings))
	copy(tmp, anythings)
	sort.Slice(tmp, f)
	return tmp
}

// DistinctBy is ..
func (anythings Channels) DistinctBy(f func(anything Channel) interface{}) (result Channels) {
	tmp := map[interface{}]Channel{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}

func (anythings Channels) Mapper(f func(anything Channel) Channel) (result Channels) {
	for _, a := range anythings {
		result = append(result, f(a))
	}
	return
}

type Encodeds []Encoded

// First is ..
func (anythings Encodeds) First() (Encoded, bool) {
	if len(anythings) == 0 {
		return Encoded{}, false
	}
	return anythings[0], true
}

// Where is ..
func (anythings Encodeds) Where(f func(anything Encoded) bool) (result Encodeds) {
	for _, a := range anythings {
		if f(a) {
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings Encodeds) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings Encodeds) CountIf(f func(anything Encoded) bool) (counter int) {
	for _, a := range anythings {
		if f(a) {
			counter++
		}
	}
	return
}

// Select is ..
func (anythings Encodeds) Select(fieldName string) (result typeslicer.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings Encodeds) SortBy(sortFunc func(Encoded, Encoded) bool) (result Encodeds) {
	f := func(i, j int) bool {
		a := anythings[i]
		b := anythings[j]
		return sortFunc(a, b)
	}

	tmp := make(Encodeds, len(anythings))
	copy(tmp, anythings)
	sort.Slice(tmp, f)
	return tmp
}

// DistinctBy is ..
func (anythings Encodeds) DistinctBy(f func(anything Encoded) interface{}) (result Encodeds) {
	tmp := map[interface{}]Encoded{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}

func (anythings Encodeds) Mapper(f func(anything Encoded) Encoded) (result Encodeds) {
	for _, a := range anythings {
		result = append(result, f(a))
	}
	return
}

type EncodedQueues []EncodedQueue

// First is ..
func (anythings EncodedQueues) First() (EncodedQueue, bool) {
	if len(anythings) == 0 {
		return EncodedQueue{}, false
	}
	return anythings[0], true
}

// Where is ..
func (anythings EncodedQueues) Where(f func(anything EncodedQueue) bool) (result EncodedQueues) {
	for _, a := range anythings {
		if f(a) {
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings EncodedQueues) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings EncodedQueues) CountIf(f func(anything EncodedQueue) bool) (counter int) {
	for _, a := range anythings {
		if f(a) {
			counter++
		}
	}
	return
}

// Select is ..
func (anythings EncodedQueues) Select(fieldName string) (result typeslicer.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings EncodedQueues) SortBy(sortFunc func(EncodedQueue, EncodedQueue) bool) (result EncodedQueues) {
	f := func(i, j int) bool {
		a := anythings[i]
		b := anythings[j]
		return sortFunc(a, b)
	}

	tmp := make(EncodedQueues, len(anythings))
	copy(tmp, anythings)
	sort.Slice(tmp, f)
	return tmp
}

// DistinctBy is ..
func (anythings EncodedQueues) DistinctBy(f func(anything EncodedQueue) interface{}) (result EncodedQueues) {
	tmp := map[interface{}]EncodedQueue{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}

func (anythings EncodedQueues) Mapper(f func(anything EncodedQueue) EncodedQueue) (result EncodedQueues) {
	for _, a := range anythings {
		result = append(result, f(a))
	}
	return
}

type Encodings []Encoding

// First is ..
func (anythings Encodings) First() (Encoding, bool) {
	if len(anythings) == 0 {
		return Encoding{}, false
	}
	return anythings[0], true
}

// Where is ..
func (anythings Encodings) Where(f func(anything Encoding) bool) (result Encodings) {
	for _, a := range anythings {
		if f(a) {
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings Encodings) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings Encodings) CountIf(f func(anything Encoding) bool) (counter int) {
	for _, a := range anythings {
		if f(a) {
			counter++
		}
	}
	return
}

// Select is ..
func (anythings Encodings) Select(fieldName string) (result typeslicer.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings Encodings) SortBy(sortFunc func(Encoding, Encoding) bool) (result Encodings) {
	f := func(i, j int) bool {
		a := anythings[i]
		b := anythings[j]
		return sortFunc(a, b)
	}

	tmp := make(Encodings, len(anythings))
	copy(tmp, anythings)
	sort.Slice(tmp, f)
	return tmp
}

// DistinctBy is ..
func (anythings Encodings) DistinctBy(f func(anything Encoding) interface{}) (result Encodings) {
	tmp := map[interface{}]Encoding{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}

func (anythings Encodings) Mapper(f func(anything Encoding) Encoding) (result Encodings) {
	for _, a := range anythings {
		result = append(result, f(a))
	}
	return
}

type Videos []Video

// First is ..
func (anythings Videos) First() (Video, bool) {
	if len(anythings) == 0 {
		return Video{}, false
	}
	return anythings[0], true
}

// Where is ..
func (anythings Videos) Where(f func(anything Video) bool) (result Videos) {
	for _, a := range anythings {
		if f(a) {
			result = append(result, a)
		}
	}
	return
}

// Count is ..
func (anythings Videos) Count() (counter int) {
	return len(anythings)
}

// CountIf is ..
func (anythings Videos) CountIf(f func(anything Video) bool) (counter int) {
	for _, a := range anythings {
		if f(a) {
			counter++
		}
	}
	return
}

// Select is ..
func (anythings Videos) Select(fieldName string) (result typeslicer.InterfaceSlice) {
	for _, a := range anythings {
		i := reflect.ValueOf(a).FieldByName(fieldName).Interface()
		result = append(result, i)
	}
	return
}

// SortBy is ..
func (anythings Videos) SortBy(sortFunc func(Video, Video) bool) (result Videos) {
	f := func(i, j int) bool {
		a := anythings[i]
		b := anythings[j]
		return sortFunc(a, b)
	}

	tmp := make(Videos, len(anythings))
	copy(tmp, anythings)
	sort.Slice(tmp, f)
	return tmp
}

// DistinctBy is ..
func (anythings Videos) DistinctBy(f func(anything Video) interface{}) (result Videos) {
	tmp := map[interface{}]Video{}
	for _, a := range anythings {
		tmp[f(a)] = a
	}
	for _, t := range tmp {
		result = append(result, t)
	}
	return
}

func (anythings Videos) Mapper(f func(anything Video) Video) (result Videos) {
	for _, a := range anythings {
		result = append(result, f(a))
	}
	return
}
