package json5

import "encoding/json"

func Unmarshal(data []byte, v interface{}) error {
	p := Parser{}
	parsedValue, err := p.Parse(data)
	if err != nil {
		return err
	}

	jsonValue, err := json.Marshal(parsedValue)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonValue, v)
}
