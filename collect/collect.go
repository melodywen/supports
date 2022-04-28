package collect

import (
	"encoding/json"
	"fmt"
	"github.com/melodywen/supports/constracts"
	"github.com/melodywen/supports/exceptions"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

// TypeTransformOrFail [V any]
//  @Description: interface data type transform
//  @param param
//  @return response
func TypeTransformOrFail[I, V any](param I) (response V) {
	var data interface{}
	data = param
	if response, ok := data.(V); ok {
		return response
	}
	panic(exceptions.NewInvalidParamError(
		fmt.Sprintf(
			"type transform error, please check code,current type:%s, want type:%s",
			reflect.TypeOf(param).Kind().String(),
			reflect.TypeOf(response).Kind().String(),
		),
	))
}

// MapSlice [V, S any]
//  @Description:Run a map over each of the items.
//  @param subject
//  @param callback
//  @return response
func MapSlice[V, S any](subject []V, callback func(int, V) S) (response []S) {
	if subject == nil {
		return response
	}
	response = make([]S, len(subject))
	for key, item := range subject {
		response[key] = callback(key, item)
	}
	return response
}

// MapMap [K comparable, V, S any]
//  @Description: Run a map over each of the items.
//  @param subject
//  @param callback
//  @return response
func MapMap[K comparable, V, S any](subject map[K]V, callback func(K, V) S) (response map[K]S) {
	if subject == nil {
		return response
	}
	response = make(map[K]S, len(subject))
	for key, item := range subject {
		response[key] = callback(key, item)
	}
	return response
}

// MapSliceWithKeys [K comparable, V, S any]
//  @Description:Run an associative map over each of the items
//				The callback should return an associative array with a single key/value pair.
//  @param subject
//  @param callback
//  @return response
func MapSliceWithKeys[K comparable, V, S any](subject []V, callback func(int, V) (K, S)) (response map[K]S) {
	if subject == nil {
		return response
	}
	response = make(map[K]S, len(subject))
	for key, item := range subject {
		mKey, mValue := callback(key, item)
		response[mKey] = mValue
	}
	return response
}

// MapMapWithKeys [K, SK comparable, V, S any]
//  @Description:Run an associative map over each of the items
//				The callback should return an associative array with a single key/value pair.
//  @param subject
//  @param callback
//  @return response
func MapMapWithKeys[K, SK comparable, V, S any](subject map[K]V, callback func(K, V) (SK, S)) (response map[SK]S) {
	if subject == nil {
		return response
	}
	response = make(map[SK]S, len(subject))
	for key, item := range subject {
		mKey, mValue := callback(key, item)
		response[mKey] = mValue
	}
	return response
}

// FilterSlice [V, S any]
//  @Description: Run a filter over each of the items.
//  @param subject
//  @param callback
//  @return response
func FilterSlice[V any](subject []V, callback func(int, V) bool) (response []V) {
	if subject == nil {
		return response
	}
	response = []V{}
	for key, item := range subject {
		if callback(key, item) {
			response = append(response, item)
		}
	}
	return response
}

// FilterMap [K comparable, V any]
//  @Description: Run a filter over each of the items.
//  @param subject
//  @param callback
//  @return response
func FilterMap[K comparable, V any](subject map[K]V, callback func(K, V) bool) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for key, item := range subject {
		if callback(key, item) {
			response[key] = item
		}
	}
	return response
}

// EachSlice [V any]
//  @Description:Execute a callback over each item.
//  @param subject
//  @param callback
func EachSlice[V any](subject []V, callback func(int, V) bool) {
	for key, item := range subject {
		if !callback(key, item) {
			break
		}
	}
}

// EachMap [K comparable, V any]
//  @Description: Execute a callback over each item.
//  @param subject
//  @param callback
func EachMap[K comparable, V any](subject map[K]V, callback func(K, V) bool) {
	for key, item := range subject {
		if !callback(key, item) {
			break
		}
	}
}

// EverySlice [V any]
//  @Description: Determine if all items pass the given truth test.
//  @param subject
//  @param callback
//  @return bool
func EverySlice[V any](subject []V, callback func(int, V) bool) bool {
	for key, item := range subject {
		if !callback(key, item) {
			return false
		}
	}
	return true
}

// EveryMap [K comparable, V any]
//  @Description: Determine if all items pass the given truth test.
//  @param subject
//  @param callback
//  @return bool
func EveryMap[K comparable, V any](subject map[K]V, callback func(K, V) bool) bool {
	for key, item := range subject {
		if !callback(key, item) {
			return false
		}
	}
	return true
}

// ContainsSlice [V comparable]
//  @Description: Determine if an item exists in the collection.
//  @param subject
//  @param item
//  @return response
func ContainsSlice[V comparable](subject []V, item V) (response bool) {
	if subject == nil {
		return response
	}
	for _, v := range subject {
		if v == item {
			return true
		}
	}
	return response
}

// ContainsMap [K, V comparable]
//  @Description: Determine if an item exists in the collection.
//  @param subject
//  @param key
//  @param value
//  @return response
func ContainsMap[K, V comparable](subject []map[K]V, key K, value V) (response bool) {
	if subject == nil {
		return response
	}
	for _, item := range subject {
		if current, ok := item[key]; ok && current == value {
			return true
		}
	}
	return response
}

// DoesNotContainsSlice [V comparable]
//  @Description:
//  @param subject
//  @param item
//  @return response
func DoesNotContainsSlice[V comparable](subject []V, item V) (response bool) {
	if subject == nil {
		return true
	}
	for _, v := range subject {
		if v == item {
			return false
		}
	}
	return true
}

// DoesNotContainsMap [K, V comparable]
//  @Description: Determine if an item is not contained in the collection.
//  @param subject
//  @param key
//  @param value
//  @return response
func DoesNotContainsMap[K, V comparable](subject []map[K]V, key K, value V) (response bool) {
	if subject == nil {
		return true
	}
	for _, item := range subject {
		if current, ok := item[key]; ok && current == value {
			return false
		}
	}
	return true
}

// Sort [V constracts.SortInterFaceGenerics]
//  @Description:
//  @param subject
//  @return response
func Sort[V constracts.SortInterFaceGenerics](subject []V) (response []V) {
	if subject == nil || len(subject) == 0 {
		return response
	}
	item := subject[0]
	switch TypeTransformOrFail[V, interface{}](item).(type) {
	case int:
		tmp := TypeTransformOrFail[[]V, []int](subject)
		sort.Ints(tmp)
		response = TypeTransformOrFail[[]int, []V](tmp)
	case string:
		tmp := TypeTransformOrFail[[]V, []string](subject)
		sort.Strings(tmp)
		response = TypeTransformOrFail[[]string, []V](tmp)
	case float64:
		tmp := TypeTransformOrFail[[]V, []float64](subject)
		sort.Float64s(tmp)
		response = TypeTransformOrFail[[]float64, []V](tmp)
	}
	return response
}

// SortDesc [V constracts.SortInterFaceGenerics]
//  @Description: Sort items in descending order.
//  @param subject
//  @return response
func SortDesc[V constracts.SortInterFaceGenerics](subject []V) (response []V) {
	if subject == nil || len(subject) == 0 {
		return response
	}
	item := subject[0]
	switch TypeTransformOrFail[V, interface{}](item).(type) {
	case int:
		tmp := TypeTransformOrFail[[]V, []int](subject)
		sort.Sort(sort.Reverse(sort.IntSlice(tmp)))
		response = TypeTransformOrFail[[]int, []V](tmp)
	case string:
		tmp := TypeTransformOrFail[[]V, []string](subject)
		sort.Sort(sort.Reverse(sort.StringSlice(tmp)))
		response = TypeTransformOrFail[[]string, []V](tmp)
	case float64:
		tmp := TypeTransformOrFail[[]V, []float64](subject)
		sort.Sort(sort.Reverse(sort.Float64Slice(tmp)))
		response = TypeTransformOrFail[[]float64, []V](tmp)
	}
	return response
}

// sortSlice[V any, ST constracts.SortInterFaceGenerics]
//  @Description: Sort items.
//  @param subject
//  @param callback
//  @param isAsc
//  @return response
func sortSlice[V any, ST constracts.SortInterFaceGenerics](subject []V, callback func(int, V) ST, isAsc bool) (response []V) {
	if subject == nil {
		return response
	}
	response = make([]V, len(subject))
	sortMap := make(map[ST]int, len(subject))
	sortContain := make([]ST, len(subject))
	for index, item := range subject {
		sortData := callback(index, item)
		sortMap[sortData] = index
		sortContain[index] = sortData
	}
	if isAsc {
		Sort(sortContain)
	} else {
		SortDesc(sortContain)
	}
	for index, item := range sortContain {
		response[index] = subject[sortMap[item]]
	}
	return response
}

// SortBy [V any, ST constracts.SortInterFaceGenerics]
//  @Description:Sort the collection using the given callback.
//  @param subject
//  @param callback
//  @return response
func SortBy[V any, ST constracts.SortInterFaceGenerics](subject []V, callback func(int, V) ST) (response []V) {
	return sortSlice(subject, callback, true)
}

// SortByDesc [V any, ST constracts.SortInterFaceGenerics]
//  @Description:  Sort the collection in descending order using the given callback.
//  @param subject
//  @param callback
//  @return response
func SortByDesc[V any, ST constracts.SortInterFaceGenerics](subject []V, callback func(int, V) ST) (response []V) {
	return sortSlice(subject, callback, false)
}

// Keys [K comparable, V any]
//  @Description: Get the keys of the collection items.
//  @param subject
//  @return response
func Keys[K comparable, V any](subject map[K]V) (response []K) {
	if subject == nil {
		return response
	}
	response = make([]K, len(subject))
	i := 0
	for key, _ := range subject {
		response[i] = key
		i++
	}
	return response
}

// Values [K comparable, V any]
//  @Description:Reset the keys on the underlying array.
//  @param subject
//  @return response
func Values[K comparable, V any](subject map[K]V) (response []V) {
	if subject == nil {
		return response
	}
	response = make([]V, len(subject))
	i := 0
	for _, item := range subject {
		response[i] = item
		i++
	}
	return response
}

// Except [K comparable, V any]
//  @Description:get all items except for those with the specified keys.
//  @param m
//  @param excepts
func Except[K comparable, V any](subject map[K]V, excepts []K) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for key, item := range subject {
		if !ContainsSlice(excepts, key) {
			response[key] = item
		}
	}
	return response
}

// Only [K comparable, V any]
//  @Description:Get the items with the specified keys.
//  @param maps
//  @param keys
//  @return response
func Only[K comparable, V any](subject map[K]V, keys []K) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for key, item := range subject {
		if ContainsSlice(keys, key) {
			response[key] = item
		}
	}
	return response
}

// Sum [V constracts.NumberInterFaceGenerics]
//  @Description: Get the sum of the given values.
//  @param subject
//  @return response
func Sum[V constracts.NumberInterFaceGenerics](subject []V) (response V) {
	for _, item := range subject {
		response += item
	}
	return response
}

// SumSlice [V any, CS constracts.NumberInterFaceGenerics]
//  @Description:Get the sum of the given values.
//  @param subject
//  @param callback
//  @return response
func SumSlice[V any, CS constracts.NumberInterFaceGenerics](subject []V, callback func(int, V) CS) (response CS) {
	for index, item := range subject {
		response += callback(index, item)
	}
	return response
}

// Average [V constracts.NumberInterFaceGenerics]
//  @Description:Get the average value of a given key.
//  @param subject
//  @param a
//  @param b
//  @return response
func Average[V constracts.NumberInterFaceGenerics](subject []V) (response V) {
	if len(subject) == 0 {
		return response
	}
	response = Sum(subject)
	length := TypeTransformOrFail[int, V](len(subject))
	return response / length
}

// AverageSlice [V any, CS constracts.NumberInterFaceGenerics]
//  @Description:Get the average value of a given key.
//  @param subject
//  @param callback
//  @return response
func AverageSlice[V any, CS constracts.NumberInterFaceGenerics](subject []V, callback func(int, V) CS) (response CS) {
	if len(subject) == 0 {
		return response
	}
	for index, item := range subject {
		response += callback(index, item)
	}
	length := TypeTransformOrFail[int, CS](len(subject))
	return response / length
}

// Chunk [V any]
//  @Description:Chunk the collection into chunks of the given size.
//  @param subject
//  @param size
//  @return response
func Chunk[V any](subject []V, size int) (response [][]V) {
	if subject == nil {
		return response
	}
	response = make([][]V, int(math.Ceil(float64(len(subject))/float64(size))))
	for index, item := range subject {
		indexOut := index / size
		indexIn := index % size
		if response[indexOut] == nil {
			response[indexOut] = make([]V, size)
		}
		response[indexOut][indexIn] = item
	}
	if lastLen := len(subject) % size; lastLen != 0 {
		response[len(response)-1] = response[len(response)-1][:len(subject)%size]
	}
	return response
}

// Collapse [T any]
//  @Description:Collapse an array of arrays into a single array.
//  @param value
//  @return response
func Collapse[T any](subject [][]T) (response []T) {
	if subject == nil {
		return response
	}
	response = []T{}
	for _, items := range subject {
		response = append(response, items...)
	}
	return response
}

// Flatten [K comparable, V any]
//  @Description:Get a flattened array of the items in the collection.
//  @param subject
//  @return response
func Flatten[K comparable, V any](subject map[K][]V) (response []V) {
	if subject == nil {
		return response
	}
	response = []V{}
	for _, items := range subject {
		response = append(response, items...)
	}
	return response
}

// Combine [K comparable, V any]
//  @Description:Chunk the collection into chunks of the given size.
//  @param keys
//  @param values
//  @return response
func Combine[K comparable, V any](keys []K, values []V) (response map[K]V) {
	if keys == nil || values == nil {
		return response
	}
	response = make(map[K]V, len(keys))
	valueLen := len(values)
	for indexKey, key := range keys {
		if indexKey < valueLen {
			response[key] = values[indexKey]
		}
	}
	return response
}

// Flip [K, V comparable]
//  @Description:Flip the items in the collection.
//  @param subject
//  @return response
func Flip[K, V comparable](subject map[K]V) (response map[V]K) {
	if subject == nil {
		return response
	}
	response = make(map[V]K, len(subject))
	for key, item := range subject {
		response[item] = key
	}
	return response
}

// FlatMap [K, NK comparable, V any]
//  @Description:
//  @param subject
//  @param callback
//  @return response
func FlatMap[K, NK comparable, V any](subject map[K]V, callback func(K, V) NK) (response map[NK]V) {
	if subject == nil {
		return response
	}
	response = make(map[NK]V, len(subject))
	for key, item := range subject {
		response[callback(key, item)] = item
	}
	return response
}

// Forget [K, V comparable]
//  @Description:Remove an item from the collection by key.
//  @param subject
//  @return response
func Forget[K comparable, V any](subject map[K]V, keys []K) (response map[K]V) {
	if subject == nil {
		return response
	}
	for _, key := range keys {
		delete(subject, key)
	}
	return subject
}

// Get [K comparable, V any]
//  @Description:Get an item from the collection by key.
//  @param subject
//  @param key
//  @return response
func Get[K comparable, V any](subject map[K]V, key K) (response V) {
	if subject == nil {
		return response
	}
	return subject[key]
}

// GetOrDefault [K comparable, V any]
//  @Description:Get an item from the collection by key.
//  @param subject
//  @param key
//  @param def
//  @return response
func GetOrDefault[K comparable, V any](subject map[K]V, key K, def V) (response V) {
	if _, ok := subject[key]; !ok {
		return def
	}
	return subject[key]
}

// Has [K comparable, V any]
//  @Description:Determine if an item exists in the collection by key.
//  @param subject
//  @param keys
//  @return response
func Has[K comparable, V any](subject map[K]V, keys []K) (response bool) {
	if subject == nil {
		return response
	}
	for _, key := range keys {
		if _, ok := subject[key]; !ok {
			return false
		}
	}
	return true
}

// First [V any]
//  @Description: Get the first item from the collection passing the given truth test.
//  @param subject
//  @param callback
//  @return response
func First[V any](subject []V, callback func(key int, value V) bool) (response V) {
	if subject == nil {
		return response
	}
	for i, item := range subject {
		if callback(i, item) {
			return item
		}
	}
	return response
}

// Last [V any]
//  @Description:
//  @param subject
//  @param callback
//  @return response
func Last[V any](subject []V, callback func(key int, value V) bool) (response V) {
	if subject == nil {
		return response
	}
	for i := len(subject) - 1; i >= 0; i-- {
		if callback(i, subject[i]) {
			return subject[i]
		}
	}
	return response
}

// CountBy [K comparable, V any]
//  @Description:Count the number of items in the collection by a field or using a callback.
//  @param subject
//  @param callback
//  @return response
func CountBy[K comparable, V any](subject []V, callback func(int, V) K) (response map[K]int) {
	if subject == nil {
		return response
	}
	response = map[K]int{}
	for index, item := range subject {
		response[callback(index, item)] += 1
	}
	return response
}

// CrossJoin [V any]
//  @Description:Cross join with the given lists, returning all possible permutations.
//  @param subject
//  @param object
//  @return response
func CrossJoin[V any](subject [][]V, object [][]V) (response [][]V) {
	if subject == nil || object == nil {
		return response
	}
	objLen := len(object)
	response = make([][]V, len(subject)*objLen)
	for key1, item1 := range subject {
		for key2, item2 := range object {
			response[key1*objLen+key2] = append(response[key1*objLen+key2], item1...)
			response[key1*objLen+key2] = append(response[key1*objLen+key2], item2...)
		}
	}
	return response
}

// Diff [V comparable]
//  @Description:Get the items in the collection that are not present in the given items.
//  @param subject
//  @param refs
//  @return response
func Diff[V comparable](subject []V, refs []V) (response []V) {
	if subject == nil {
		return response
	}
	response = []V{}
	for _, item := range subject {
		if !ContainsSlice(refs, item) {
			response = append(response, item)
		}
	}
	return response
}

// DiffAssoc [K, V comparable]
//  @Description:Get the items in the collection whose keys and values are not present in the given items.
//  @param subject
//  @param refs
//  @return response
func DiffAssoc[K, V comparable](subject map[K]V, refs map[K]V) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for key, item := range subject {
		if tmp, ok := refs[key]; !ok || (tmp != item) {
			response[key] = item
		}
	}
	return response
}

// DiffKeys [K comparable, V any]
//  @Description:Get the items in the collection whose keys are not present in the given items.
//  @param subject
//  @param refs
//  @return response
func DiffKeys[K comparable, V any](subject map[K]V, refs map[K]V) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for key, item := range subject {
		if _, ok := refs[key]; !ok {
			response[key] = item
		}
	}
	return response
}

// Intersect [V comparable]
//  @Description:
//  @param subject
//  @param refs
//  @return response
func Intersect[V comparable](subject []V, refs []V) (response []V) {
	if subject == nil {
		return response
	}
	response = []V{}
	for _, item := range subject {
		if ContainsSlice(refs, item) {
			response = append(response, item)
		}
	}
	return response
}

// IntersectByKeys [K comparable, V any]
//  @Description:
//  @param subject
//  @param refs
//  @return response
func IntersectByKeys[K comparable, V any](subject map[K]V, refs map[K]V) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for key, item := range subject {
		if _, ok := refs[key]; ok {
			response[key] = item
		}
	}
	return response
}

// Min [V constracts.NumberInterFaceGenerics]
//  @Description:
//  @param subject
//  @return response
func Min[V constracts.NumberInterFaceGenerics](subject []V) (response V) {
	for index, item := range subject {
		if index == 0 || item < response {
			response = item
		}
	}
	return response
}

// MinSlice [V any, CS constracts.NumberInterFaceGenerics]
//  @Description:
//  @param subject
//  @param callback
//  @return response
func MinSlice[V any, CS constracts.NumberInterFaceGenerics](subject []V, callback func(int, V) CS) (response CS) {
	for index, item := range subject {
		if tmp := callback(index, item); index == 0 || tmp < response {
			response = tmp
		}
	}
	return response
}

// Max [V constracts.NumberInterFaceGenerics]
//  @Description:
//  @param subject
//  @return response
func Max[V constracts.NumberInterFaceGenerics](subject []V) (response V) {
	for index, item := range subject {
		if index == 0 || item > response {
			response = item
		}
	}
	return response
}

// MaxSlice [V any, CS constracts.NumberInterFaceGenerics]
//  @Description:
//  @param subject
//  @param callback
//  @return response
func MaxSlice[V any, CS constracts.NumberInterFaceGenerics](subject []V, callback func(int, V) CS) (response CS) {
	for index, item := range subject {
		if tmp := callback(index, item); index == 0 || tmp > response {
			response = tmp
		}
	}
	return response
}

// Pop [V any]
//  @Description:get and remove the last N items from the collection.
//  @param subject
//  @return response
func Pop[V any](subject []V) (response []V, item V) {
	if subject == nil {
		return response, item
	}
	subLen := len(subject)
	if subLen == 0 {
		return []V{}, item
	} else if subLen == 1 {
		return []V{}, subject[0]
	}
	return subject[:subLen-1], subject[subLen-1]
}

// Prepend [V any]
//  @Description:Push an item onto the beginning of the collection.
//  @param subject
//  @param item
//  @return response
func Prepend[V any](subject []V, item ...V) (response []V) {
	return append(item, subject...)
}

// Pull [K comparable, V any]
//  @Description:Get and remove an item from the collection.
//  @param subject
//  @param key
//  @return response
//  @return item
func Pull[K comparable, V any](subject map[K]V, key K) (response map[K]V, item V) {
	if subject == nil {
		return response, item
	}
	item = subject[key]
	delete(subject, key)
	return subject, item
}

// Push [V any]
//  @Description:Push one or more items onto the end of the collection.
//  @param subject
//  @param item
//  @return response
func Push[V any](subject []V, item ...V) (response []V) {
	return append(subject, item...)
}

// Put [K comparable, V any]
//  @Description:Put an item in the collection by key.
//  @param subject
//  @param key
//  @param item
//  @return response
func Put[K comparable, V any](subject map[K]V, key K, item V) (response map[K]V) {
	if subject == nil {
		response = map[K]V{}
	} else {
		response = subject
	}
	response[key] = item
	return response
}

// IsEmpty [V any]
//  @Description:Determine if the collection is empty or not.
//  @param subject
//  @return bool
func IsEmpty[V any](subject V) bool {
	switch reflect.TypeOf(subject).Kind() {
	case reflect.String:
		if TypeTransformOrFail[V, string](subject) == "" {
			return true
		}
	case reflect.Int:
		if TypeTransformOrFail[V, int](subject) == 0 {
			return true
		}
	case reflect.Float64:
		if TypeTransformOrFail[V, float64](subject) == 0 {
			return true
		}
	case reflect.Bool:
		if TypeTransformOrFail[V, bool](subject) == false {
			return true
		}
	default:
		panic(exceptions.NewInvalidParamError(
			fmt.Sprintf(
				"type transform error, please check code(is empty method)",
			),
		))
	}
	return false
}

// IsEmptySlice [V any]
//  @Description:Determine if the collection is empty or not.
//  @param subject
//  @return bool
func IsEmptySlice[V any](subject []V) bool {
	if subject == nil || len(subject) == 0 {
		return true
	}
	return false
}

// IsEmptyMap [K comparable, V any]
//  @Description: Determine if the collection is empty or not.
//  @param subject
//  @return bool
func IsEmptyMap[K comparable, V any](subject map[K]V) bool {
	if subject == nil || len(subject) == 0 {
		return true
	}
	return false
}

// IsNotEmpty [V any]
//  @Description: Determine if the collection is not empty.
//  @param subject
//  @return bool
func IsNotEmpty[V any](subject V) bool {
	return !IsEmpty(subject)
}

// IsNotEmptySlice [V any]
//  @Description:
//  @param subject
//  @return bool
func IsNotEmptySlice[V any](subject []V) bool {
	return !IsEmptySlice(subject)
}

// IsNotEmptyMap [K comparable, V any]
//  @Description:
//  @param subject
//  @return bool
func IsNotEmptyMap[K comparable, V any](subject map[K]V) bool {
	return !IsEmptyMap(subject)
}

// Implode [V any]
//  @Description:
//  @param subject
//  @param callback
//  @param glue
//  @return response
func Implode[V any](subject []V, callback func(int, V) string, glue string) (response string) {
	if subject == nil {
		return response
	}
	stringSlice := MapSlice(subject, func(index int, item V) string {
		return callback(index, item)
	})
	return strings.Join(stringSlice, glue)
}

// KeyBy [K comparable, V any]
//  @Description:Key an associative  using a callback.
//  @param subject
//  @param callback
//  @return response
func KeyBy[K comparable, V any](subject []V, callback func(int, V) K) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = make(map[K]V, len(subject))
	for index, item := range subject {
		response[callback(index, item)] = item
	}
	return response
}

// Merge [K comparable, V any]
// @Description: map merge
// @param arg
func Merge[K comparable, V any](subject ...map[K]V) (response map[K]V) {
	if subject == nil {
		return response
	}
	response = map[K]V{}
	for _, item := range subject {
		for key, value := range item {
			response[key] = value
		}
	}
	return response
}

// MergeRecursive [K comparable, V any]
//  @Description:Recursively merge the collection with the given items.
//  @param subject
//  @return response
func MergeRecursive[K comparable, V any](subject ...map[K][]V) (response map[K][]V) {
	if subject == nil {
		return response
	}
	response = map[K][]V{}
	for _, item := range subject {
		for key, value := range item {
			response[key] = append(response[key], value...)
		}
	}
	return response
}

// Skip [V any]
//  @Description: Skip the first {$count} items.
//  @param subject
//  @param skip
//  @return response
func Skip[V any](subject []V, offset int) (response []V) {
	if subject == nil {
		return response
	}
	if offset > len(subject) {
		return []V{}
	}
	return subject[offset:]
}

// Slice [V any]
//  @Description:Slice the underlying collection array.
//  @param subject
//  @param offset
//  @param length
//  @return response
func Slice[V any](subject []V, offset int, length int) (response []V) {
	if subject == nil {
		return response
	}
	subLen := len(subject)
	if offset > subLen || offset+subLen < 0 {
		return []V{}
	}
	if offset < 0 {
		offset += subLen
	}
	last := offset + length
	if last > subLen {
		last = subLen
	}
	return subject[offset:last]
}

// Nth [V any]
//  @Description:Create a new collection consisting of every n-th element.
//  @param subject
//  @param offset
//  @return response
func Nth[V any](subject []V, offset int) (response []V) {
	if subject == nil {
		return response
	}
	response = make([]V, len(subject)/offset)
	resLen := len(response)
	for i := 0; i < resLen; i++ {
		response[i] = subject[i*offset]
	}
	return response
}

// Pad [V any]
//  @Description: Pad collection to the specified length with a value.
//  @param subject
//  @param size
//  @param item
//  @return response
func Pad[V any](subject []V, size int, item V) (response []V) {
	isAfter := true
	if size < 0 {
		isAfter = false
		size = -size
	}
	if len(subject) >= size {
		return subject
	}
	result := make([]V, size-len(subject))
	for i := size - len(subject) - 1; i >= 0; i-- {
		result[i] = item
	}
	if isAfter {
		return append(subject, result...)
	}
	return append(result, subject...)
}

// Partition [V any]
//  @Description:Partition the collection into two arrays using the given callback .
//  @param subject
//  @param callback
//  @return pass
//  @return fail
func Partition[V any](subject []V, callback func(int, V) bool) (pass []V, fail []V) {
	if subject == nil {
		return pass, fail
	}
	pass = []V{}
	fail = []V{}
	for index, item := range subject {
		if callback(index, item) {
			pass = append(pass, item)
		} else {
			fail = append(fail, item)
		}
	}
	return pass, fail
}

// Pluck [K comparable, V any]
//  @Description:Get the values of a given key.
//  @param subject
//  @param key
//  @return response
func Pluck[K comparable, V any](subject []map[K]V, key K) (response []V) {
	if subject == nil {
		return response
	}
	for _, items := range subject {
		if item, ok := items[key]; ok {
			response = append(response, item)
		}
	}
	return response
}

// Range
//  @Description:Create a collection with the given range.
//  @param from
//  @param to
//  @return response
func Range(from, to int) (response []int) {
	response = make([]int, to-from+1)
	for i := from; i <= to; i++ {
		response[i-from] = i
	}
	return response
}

// ForPage [V any]
//  @Description:"Paginate" the collection by slicing it into a smaller collection.
//  @param subject
//  @param page
//  @param perPage
//  @return response
func ForPage[V any](subject []V, page, perPage int) (response []V) {
	if subject == nil {
		return response
	}
	offset := Max([]int{0, (page - 1) * perPage})

	return Slice(subject, offset, perPage)
}

// GroupBy [K comparable, V any]
//  @Description:Group an associative array by a field or using a callback.
//  @param subject
//  @param callback
//  @return response
func GroupBy[K comparable, V any](subject []V, callback func(int, V) K) (response map[K][]V) {
	if subject == nil {
		return response
	}
	response = map[K][]V{}
	for index, item := range subject {
		tmpKey := callback(index, item)
		response[tmpKey] = append(response[tmpKey], item)
	}
	return response
}

// Shuffle [V any]
// @Description:Shuffle the given array and return the result.
// @param slices
// @return response
func Shuffle[V any](subject []V) (response []V) {
	if subject == nil {
		return response
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(subject), func(i, j int) {
		subject[i], subject[j] = subject[j], subject[i]
	})
	return subject
}

// Random [V any]
//  @Description:Get a specified number of items randomly from the collection.
//  @param subject
//  @param number
//  @return response
func Random[V any](subject []V, number int) (response []V) {
	if number > len(subject) {
		number = len(subject)
	}
	response = Shuffle(subject)[:number]
	return response
}

// Reduce [V, S any]
//  @Description: Reduce the collection to a single value.
//  @param subject
//  @param callback
//  @return response
func Reduce[V, S any](subject []V, callback func(S, int, V) S) (response S) {
	if subject == nil {
		return response
	}
	for index, item := range subject {
		response = callback(response, index, item)
	}
	return response
}

// Reverse [K, V comparable]
//  @Description:Reverse items order.
//  @param subject
//  @return response
func Reverse[K, V comparable](subject map[K]V) (response map[V]K) {
	if subject == nil {
		return response
	}
	response = make(map[V]K, len(subject))
	for key, item := range subject {
		response[item] = key
	}
	return response
}

// SearchSlice [V comparable]
//  @Description: Search the collection for a given value and return the corresponding key if successful.
//  @param subject
//  @param search
//  @return response
//  @return ok
func SearchSlice[V comparable](subject []V, search V) (response int, ok bool) {
	if subject == nil {
		return response, ok
	}
	for index, item := range subject {
		if item == search {
			return index, true
		}
	}
	return response, false
}

// SearchMap [K, V comparable]
//  @Description:Search the collection for a given value and return the corresponding key if successful.
//  @param subject
//  @param search
//  @return response
//  @return ok
func SearchMap[K, V comparable](subject map[K]V, search V) (response K, ok bool) {
	if subject == nil {
		return response, ok
	}
	for key, item := range subject {
		if item == search {
			return key, true
		}
	}
	return response, false

}

// Shift [V any]
//  @Description:Search the collection for a given value and return the corresponding key if successful.
//  @param subject
//  @return response
//  @return item
func Shift[V any](subject []V) (response []V, item V) {
	if subject == nil {
		return response, item
	}
	subLen := len(subject)
	if subLen == 0 {
		return []V{}, item
	} else if subLen == 1 {
		return []V{}, subject[0]
	}
	return subject[1:subLen], subject[0]
}

// Times [V any]
//  @Description:Create a new collection by invoking the callback a given amount of times.
//  @param times
//  @param callback
//  @return response
func Times[V any](times int, callback func(index int) V) (response []V) {
	if times <= 0 {
		return response
	}
	response = make([]V, times)
	for i := 1; i <= times; i++ {
		response[i-1] = callback(i)
	}
	return response
}

// Sliding [V any]
//  @Description: Create chunks representing a "sliding window" view of the items in the collection.
//  @param subject
//  @param size
//  @param step
//  @return response
func Sliding[V any](subject []V, size, step int) (response [][]V) {
	if subject == nil || len(subject)-size < 0 {
		return response
	}
	chunks := ((len(subject) - size) / step) + 1
	return Times(chunks, func(index int) []V {
		return Slice(subject, (index-1)*step, size)
	})
}

// Splice [V any]
//  @Description:Splice a portion of the underlying collection array.
//  @param subject
//  @param offset
//  @param length
//  @return response
//  @return other
func Splice[V any](subject []V, offset, length int) (response []V, other []V) {
	if subject == nil {
		return response, other
	}
	subLen := len(subject)
	if offset > subLen || offset+subLen < 0 {
		return []V{}, subject
	}
	if offset < 0 {
		offset += subLen
	}
	last := offset + length
	if last > subLen {
		last = subLen
	}
	response = []V{}
	other = []V{}
	for i := 0; i < subLen; i++ {
		if i >= offset && i < last {
			response = append(response, subject[i])
		} else {
			other = append(other, subject[i])
		}
	}
	return response, other
}

// Split [V any]
//  @Description:Split a collection into a certain number of groups, and fill the first groups completely.
//  @param subject
//  @param numberOfGroups
//  @return response
func Split[V any](subject []V, numberOfGroups int) (response [][]V) {
	if subject == nil || numberOfGroups < 1 {
		return response
	}
	groupSize := int(math.Floor(float64(len(subject)) / float64(numberOfGroups)))

	remain := len(subject) % numberOfGroups
	start := 0
	for i := 0; i < numberOfGroups; i++ {
		size := groupSize
		if i < remain {
			size++
		}
		if size != 0 {
			response = Push(response, Slice(subject, start, size))
			start += size
		}
	}
	return response
}

// ToJson [V any]
//  @Description:Get the collection of items as JSON.
//  @param subject
//  @return string
func ToJson[V any](subject V) string {
	response, e := json.Marshal(subject)
	if e != nil {
		panic(exceptions.NewInvalidParamErrorWithData(e.Error(), string(response)))
	}
	return string(response)
}

// Union [V any]
//  @Description:Union the collection with the given items.
//  @param subject
//  @return response
func Union[V any](subject ...[]V) (response []V) {
	for _, items := range subject {
		response = append(response, items...)
	}
	return response
}

// Unique [K comparable, V any]
//  @Description:Return only unique items from the collection array.
//  @param subject
//  @param callback
//  @return response
func Unique[K comparable, V any](subject []V, callback func(int, V) K) (response []V) {
	if subject == nil {
		return response
	}
	uniMap := map[K]int{}
	for index, item := range subject {
		uniMap[callback(index, item)] = index
	}
	revMap := Reverse(uniMap)
	response = FilterSlice(subject, func(key int, _ V) bool {
		_, ok := revMap[key]
		return ok
	})
	return response
}

// Add [K comparable, V any]
//  @Description:Add an element to an array using "dot" notation if it doesn't exist.
//  @param m
//  @param key
//  @param value
func Add[K comparable, V any](subject map[K]V, key K, value V) (response map[K]V) {
	if subject == nil {
		subject = map[K]V{}
	}
	if _, ok := subject[key]; !ok {
		subject[key] = value
	}
	return subject
}

// Wrap [V, S any]
//  @Description:Wrap the given value in a collection if applicable.
//  @param subject
//  @return response
func Wrap[V, S any](subject V) (response []S) {
	if reflect.TypeOf(subject).Kind() == reflect.Slice {
		return TypeTransformOrFail[V, []S](subject)
	}
	return []S{TypeTransformOrFail[V, S](subject)}
}

// Zip [V any]
//  @Description:Zip the collection together with one or more arrays.
//  @param keys
//  @param values
//  @return response
func Zip[V any](keys []V, values []V) (response [][]V) {
	if keys == nil || values == nil {
		return response
	}
	valueLen := len(values)
	response = make([][]V, Min([]int{len(keys), valueLen}))
	for indexKey, key := range keys {
		if indexKey >= valueLen {
			break
		}
		response[indexKey] = []V{key, values[indexKey]}
	}
	return response
}
