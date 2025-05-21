package util

import (
	"encoding/xml"
	"fmt"
	"os"
	"sync"
)

// QueryLoader loads SQL queries from an XML file.
type QueryLoader struct {
	queries map[string]string
	once    sync.Once
}

var loader = &QueryLoader{}

// load reads and parses queries.xml from the resources directory.
func (q *QueryLoader) load() {
	data, err := os.ReadFile("src/main/resources/queries.xml")
	if err != nil {
		panic(fmt.Errorf("failed to read queries.xml: %w", err))
	}
	var root struct {
		XMLName xml.Name `xml:"queries"`
		Queries []struct {
			ID  string `xml:"id,attr"`
			SQL string `xml:",chardata"`
		} `xml:"query"`
	}
	if err := xml.Unmarshal(data, &root); err != nil {
		panic(fmt.Errorf("failed to parse queries.xml: %w", err))
	}
	q.queries = make(map[string]string)
	for _, el := range root.Queries {
		q.queries[el.ID] = el.SQL
	}
}

// GetQuery returns the SQL for the given query ID.
func GetQuery(id string) string {
	loader.once.Do(loader.load)
	return loader.queries[id]
}
