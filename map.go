package transform

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// SpecFromMap parses the given map into the spec
func SpecFromMap(data map[string]string, spec interface{}) error {
	if data == nil {
		return fmt.Errorf("data is nil")
	}
	decodeCfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   spec,
		TagName:  "json",
	}
	decoder, err := mapstructure.NewDecoder(decodeCfg)
	if err != nil {
		return err
	}
	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}

// SpecFromMap parses the given map into the spec
func SpecFromMapInterface(data map[string]interface{}, spec interface{}) error {
	if data == nil {
		return fmt.Errorf("data is nil")
	}
	decodeCfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   spec,
		TagName:  "json",
	}
	decoder, err := mapstructure.NewDecoder(decodeCfg)
	if err != nil {
		return err
	}
	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
