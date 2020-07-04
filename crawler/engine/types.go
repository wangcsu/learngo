package engine

// Request stores a url and its corresponding handler func
type Request struct {
	URL        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult stores a list of parsed results
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// NilParser handles undefined parsers
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
