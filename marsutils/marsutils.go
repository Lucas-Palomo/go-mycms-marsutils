package marsutils

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Lucas-Palomo/go-mycms-marsutils/internal"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

type Type int

const (
	XML  Type = iota
	JSON Type = iota
	TOML Type = iota
	YAML Type = iota
)

func (t Type) Marshal() func(model interface{}) ([]byte, error) {
	marshal := []func(model interface{}) ([]byte, error){
		xml.Marshal,
		json.Marshal,
		toml.Marshal,
		yaml.Marshal,
	}

	return marshal[t]
}

func (t Type) Unmarshal() func(data []byte, model interface{}) error {
	unmarshal := []func(data []byte, model interface{}) error{
		xml.Unmarshal,
		json.Unmarshal,
		toml.Unmarshal,
		yaml.Unmarshal,
	}

	return unmarshal[t]
}

func ReadFile(filepath string, model interface{}, filetype Type) error {
	data, err := internal.ReadFile(filepath)

	if err != nil {
		return err
	}

	return internal.Unmarshal(data, model, filetype.Unmarshal())
}

func WriteFile(filepath string, model interface{}, filetype Type) error {
	content, err := ToString(model, filetype)

	if err != nil {
		return err
	}

	return internal.WriteFile(filepath, []byte(content))
}

func FromString(content string, model interface{}, filetype Type) error {
	return internal.Unmarshal([]byte(content), model, filetype.Unmarshal())
}

func ToString(model interface{}, filetype Type) (string, error) {
	content, err := internal.Marshal(model, filetype.Marshal())
	if err != nil {
		return "", err
	}

	return string(content), nil
}
