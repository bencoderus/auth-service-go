package http

import (
	"encoding/json"
)

func ParseBody[K any](body []byte) K {
	var payload K

	json.Unmarshal(body, &payload)

	return payload
}
