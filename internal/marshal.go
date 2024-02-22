package internal

func Marshal(
	model interface{},
	method func(model interface{}) ([]byte, error)) ([]byte, error) {

	data, err := method(model)
	if err != nil {
		return data, err
	}

	return data, nil
}
