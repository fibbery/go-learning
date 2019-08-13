package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func Any(v interface{}) string {
	return formatAtom(reflect.ValueOf(v))
}

func formatAtom(value reflect.Value) string {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.String:
		return strconv.Quote(value.String())
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Func, reflect.Ptr:
		return value.Type().String() + " 0x" + strconv.FormatUint(uint64(value.Pointer()), 16)
	default:
		return value.Type().String() + " value"
	}
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T)\n ", name, x)
	display(name, reflect.ValueOf(x))
}

func display(name string, x reflect.Value) {
	switch x.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", name)
	case reflect.Slice, reflect.Array:
		for i := 0; i < x.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", name, i), x.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			display(fmt.Sprintf("%s.%s", name, x.Type().Field(i).Name), x.Field(i))
		}
	case reflect.Map:
		for _, key := range x.MapKeys() {
			display(fmt.Sprintf("%s[%s]", name, formatAtom(key)), x.MapIndex(key))
		}
	case reflect.Ptr:
		if x.IsNil() {
			fmt.Printf("%s=nil\n", name)
		} else {
			display(fmt.Sprintf("(*%s)", name), x.Elem())
		}
	case reflect.Interface:
		if x.IsNil() {
			fmt.Printf("%s=nil\n", name)
		} else {
			fmt.Printf("%s.Type = %s\n", name, x.Elem().Type())
			display(name+".value", x.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", name, formatAtom(x))
	}
}

func main() {
	var a interface{} = 3
	//Display("a", a)
	Display("&a", &a)
}
