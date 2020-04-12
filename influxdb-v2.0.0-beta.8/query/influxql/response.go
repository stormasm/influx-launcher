package influxql

// all of this code is copied more or less verbatim from the influxdb repo.
// we copy instead of sharing because we want to prevent inadvertent breaking
// changes introduced by the transpiler vs the actual InfluxQL engine.
// By copying the code, we'll be able to detect more explicitly that the
// results generated by the transpiler diverge from InfluxQL.

type Response struct {
	Results []Result `json:"results,omitempty"`
	Err     string   `json:"error,omitempty"`
}

func (r *Response) error(err error) {
	r.Results = nil
	r.Err = err.Error()
}

// Message represents a user-facing message to be included with the result.
type Message struct {
	Level string `json:"level"`
	Text  string `json:"text"`
}

// Result represents a resultset returned from a single statement.
// Rows represents a list of rows that can be sorted consistently by name/tag.
type Result struct {
	// StatementID is just the statement's position in the query. It's used
	// to combine statement results if they're being buffered in memory.
	StatementID int        `json:"statement_id"`
	Series      []*Row     `json:"series,omitempty"`
	Messages    []*Message `json:"messages,omitempty"`
	Partial     bool       `json:"partial,omitempty"`
	Err         string     `json:"error,omitempty"`
}

// Row represents a single row returned from the execution of a statement.
type Row struct {
	Name    string            `json:"name,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
	Columns []string          `json:"columns,omitempty"`
	Values  [][]interface{}   `json:"values,omitempty"`
	Partial bool              `json:"partial,omitempty"`
}
