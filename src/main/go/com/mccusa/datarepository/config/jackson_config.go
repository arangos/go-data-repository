package config

import (
	"encoding/json"
	"time"
)

// JSONTime is a custom time type for RFC3339 formatting in JSON
type JSONTime time.Time

// MarshalJSON formats the time as an RFC3339 string
func (t JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339))
}

// UnmarshalJSON parses an RFC3339 time string
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	*t = JSONTime(parsed)
	return nil
}

// DefaultObjectMapper is an alias for the standard JSON encoder/decoder
var DefaultObjectMapper = json.Marshal
