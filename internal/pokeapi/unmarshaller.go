package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func unmarshalJsonBodyIntoGivenStruct[T any](res *http.Response, target *T) ([]byte, error) {
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return data, err
	}

	return data, nil
}
