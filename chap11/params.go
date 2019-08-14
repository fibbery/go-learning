package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/search",search)
	log.Fatal(http.ListenAndServe("127.0.0.1:9090", nil))
}

func search(response http.ResponseWriter, request *http.Request) {
	var data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}
	data.MaxResults = 10
	if err := Unpack(request, &data); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(response, "Search : %+v\n", data)
}

func Unpack(request *http.Request, ptr interface{}) error {
	if err := request.ParseForm(); err != nil {
		return err
	}

	fields := make(map[string]reflect.Value)
	value := reflect.ValueOf(ptr).Elem()
	for i := 0; i < value.NumField(); i++ {
		fieldInfo := value.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = value.Field(i)
	}

	for name, values := range request.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s : %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s : %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		i, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(i)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
