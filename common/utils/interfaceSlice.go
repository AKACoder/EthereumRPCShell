package utils

type toInterfaceSlice struct{}

func (t toInterfaceSlice) FromUint64Slice(i []uint64) []interface{} {
	var r []interface{}
	for _, v := range i {
		r = append(r, v)
	}

	return r
}

var ToInterfaceSlice = toInterfaceSlice{}
