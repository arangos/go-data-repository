package model

// JsonMetadata represents a key-value metadata item stored in JSON columns
type JsonMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
