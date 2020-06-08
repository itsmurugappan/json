package pkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"

	validator "github.com/go-playground/validator/v10"
)

// ErrInvalidInput indicates invalid input
// to parse the json string into
var ErrInvalidInput = errors.New("input must be a struct pointer")
var validate = validator.New()

// JsonToSpec transforms the byte array to the given struct
func JsonToSpec(data []byte, spec interface{}) error {
	s := reflect.ValueOf(spec)
	if s.Kind() != reflect.Ptr {
		return ErrInvalidInput
	}

	t := reflect.TypeOf(spec)
	new := reflect.New(t.Elem())
	i := new.Interface()
	json.Unmarshal(data, i)

	s.Elem().Set((reflect.ValueOf(i)).Elem())

	return nil
}

// ParseReqBodyToSpec parses HTTP request body into the given struct
func ParseReqBodyToSpec(r *http.Request, spec interface{}) error {
	body, _ := ioutil.ReadAll(r.Body)
	return JsonToSpec(body, spec)
}

// ParseAndValidate parses the HTTP request body into struct
// and Validate based on struct tags
func ParseAndValidate(r *http.Request, spec interface{}) error {
	if err := ParseReqBodyToSpec(r, spec); err != nil {
		return err
	}
	switch (reflect.ValueOf(spec)).Elem().Kind() {
	case reflect.Struct:
		return validate.Struct(spec)
	case reflect.Slice:
		return validate.Var(spec, "dive")
	default:
		return validate.Struct(spec)
	}
}
