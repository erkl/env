package env

import (
	"os"
	"reflect"
	"strconv"
)

// Load accepts the pointer to a struct, iterates over its fields and
// for each one attempts to load and parse the environment variable with
// a corresponding name.
func Load(v interface{}) {
	rv := reflect.ValueOf(v)

	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		if rf := rt.Field(i); rf.PkgPath == "" {
			load(rv.Field(i), rf.Name)
		}
	}
}

func load(rv reflect.Value, name string) error {
	raw, ok := os.LookupEnv(name)
	if !ok {
		return nil
	}

	switch rv.Kind() {
	case reflect.Bool:
		x, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return err
		}
		rv.SetBool(x != 0)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		x, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return err
		}
		rv.SetInt(x)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		x, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return err
		}
		rv.SetUint(x)

	case reflect.Float32, reflect.Float64:
		x, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return err
		}
		rv.SetFloat(x)

	case reflect.String:
		rv.SetString(raw)
	}

	return nil
}
