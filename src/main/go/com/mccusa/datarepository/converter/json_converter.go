package converter

import (
	"encoding/json"
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
)

// JsonConverter handles marshalling/unmarshalling of JsonMetadata slices to/from JSON strings.
type JsonConverter struct{}

// NewJsonConverter constructs a new JsonConverter.
func NewJsonConverter() JsonConverter {
	return JsonConverter{}
}

// ToDatabaseColumn converts a slice of JsonMetadata to a JSON string.
func (jc JsonConverter) ToDatabaseColumn(metadata []model.JsonMetadata) (string, error) {
	data, err := json.Marshal(metadata)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}
	return string(data), nil
}

// FromDatabaseColumn converts a JSON string to a slice of JsonMetadata.
func (jc JsonConverter) FromDatabaseColumn(dbData string) ([]model.JsonMetadata, error) {
	if dbData == "" {
		return nil, nil
	}
	var metadata []model.JsonMetadata
	if err := json.Unmarshal([]byte(dbData), &metadata); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}
	return metadata, nil
}
