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
		fileName string
		file     *os.File
		reader   *bufio.Reader
		decoder  *json.Decoder
		jsonType int
		index    uint64
	}
)

func New(fileName string) *Reader {
	r := &Reader{
		fileName: fileName,
	}

	return r
}

func (r *Reader) Start() error {
	r.jsonType = Invalid
	file, err := os.Open(r.fileName)
	if err != nil {
		return fmt.Errorf("file '%s' not found: %w", r.fileName, err)
	}

	r.file = file
	r.reader = bufio.NewReader(file)
	r.decoder = json.NewDecoder(r.reader)

	token, err := r.decoder.Token()
	if err == io.EOF {
		return fmt.Errorf("file '%s' is empty", r.fileName)
	}
	if err != nil {
		return fmt.Errorf("token decoding error: %w", err)
	}

	if token == json.Delim('{') {
		r.jsonType = Object
	} else if token == json.Delim('{') {
		r.jsonType = Array
	} else {
		return fmt.Errorf("file '%s' is not valid json file", r.file.Name())
	}

	return nil
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

func (r *Reader) Close() error {
	if err := r.file.Close(); err != nil {
		return err
	}
	return nil
}
