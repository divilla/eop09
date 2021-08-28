package largejsonreader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	Invalid = iota
	Object
	Array
)

type (
	Reader struct {
		file     *os.File
		reader   *bufio.Reader
		decoder  *json.Decoder
		jsonType int
		index    uint64
	}
)

func New(fileName string) (*Reader, error) {
	jsonType := Invalid
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("file '%s' not found: %w", fileName, err)
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	token, err := decoder.Token()
	if err == io.EOF {
		return nil, fmt.Errorf("file '%s' is empty", fileName)
	}
	if err != nil {
		return nil, fmt.Errorf("token decoding error: %w", err)
	}

	if token == json.Delim('{') {
		jsonType = Object
	} else if token == json.Delim('{') {
		jsonType = Array
	} else {
		return nil, fmt.Errorf("file '%s' is not valid json file", fileName)
	}

	return &Reader{
		file:     file,
		reader:   reader,
		decoder:  decoder,
		jsonType: jsonType,
	}, nil
}

func (r *Reader) Read(index *uint64, key *string, value *json.RawMessage) error {
	var token json.Token
	var err error

	if r.jsonType == Object {
		token, err = r.decoder.Token()
		if err != nil {
			return fmt.Errorf("token decoding error: %w", err)
		}

		if token == json.Delim('}') {
			return io.EOF
		}

		switch tt := token.(type) {
		case string:
			*key = tt
		default:
			return fmt.Errorf("string token expected: %w", err)
		}
	}

	err = r.decoder.Decode(value)
	if err == io.EOF {
		return err
	} else if err != nil {
		return fmt.Errorf("invalid or malformed json object: %w", err)
	}

	*index = r.index
	r.index++

	return nil
}

func (r *Reader) Reset() {
	r.reader.Reset(r.file)
}

func (r *Reader) Close() error {
	if err := r.file.Close(); err != nil {
		return err
	}
	return nil
}
