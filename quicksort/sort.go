package quicksort

import (
	"errors"
	"fmt"
	"reflect"
)

var handle = make(map[string]sort)

type sort interface {
	Sort(val interface{})
}

func register(typ string, h sort) {
	if _, dup := handle[typ]; dup {
		fmt.Println(typ, " have already registered.")
	} else {
		handle[typ] = h
	}
}

func QuickSort(val interface{}) error {

	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Slice {
		if v.Len() <= 1 {
			return nil
		}
		switch v.Index(0).Kind() {
		case reflect.Int:
			handle["int"].Sort(val)
			return nil
		case reflect.Int8:
			handle["int8"].Sort(val)
			return nil
		case reflect.Int16:
			handle["int16"].Sort(val)
			return nil
		case reflect.Int32:
			handle["int32"].Sort(val)
			return nil
		case reflect.Int64:
			handle["int64"].Sort(val)
			return nil
		case reflect.Float32:
			handle["float32"].Sort(val)
			return nil
		case reflect.Float64:
			handle["float64"].Sort(val)
			return nil
		case reflect.String:
			handle["string"].Sort(val)
			return nil
		case reflect.Uint:
			handle["uint"].Sort(val)
			return nil
		case reflect.Uint8:
			handle["uint8"].Sort(val)
			return nil
		case reflect.Uint16:
			handle["uint16"].Sort(val)
			return nil
		case reflect.Uint32:
			handle["uint32"].Sort(val)
			return nil
		case reflect.Uint64:
			handle["uint64"].Sort(val)
			return nil
		default:
			fmt.Println("unsupported type.")
		}
	}

	return errors.New("unsupported type.")

}
