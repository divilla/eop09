package jsonfilereader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/divilla/eop09/internal/domain"
	interfaces2 "github.com/divilla/eop09/internal/interfaces"
	"io"
	"os"
	"sync"
)

type (
	JsonFileReader struct {
		fileName string
		file     *os.File
		decoder  *json.Decoder
		logger   interfaces2.Logger
	}

	Callback func(wg *sync.WaitGroup, parser interface{}, err error)
)

func Init(fileName string, logger interfaces2.Logger) *JsonFileReader {
	file, err := os.Open(fileName)
	if err != nil {
		logger.Fatalf("file '%s' not found: %s", fileName, err.Error())
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	token, err := decoder.Token()
	if err == io.EOF {
		logger.Fatalf("file '%s' is empty", fileName)
	}
	if token != json.Delim('{') {
		logger.Fatalf("file '%s' is not valid json file", fileName)
	}
	if err != nil {
		logger.Fatalf("decoder token error: %w", err)
	}

	return &JsonFileReader{
		fileName: fileName,
		file:     file,
		decoder:  decoder,
	}
}

func (r *JsonFileReader) Parse(parser interface{}, callback Callback) {
	var token json.Token
	var key string
	var err error
	wg := new(sync.WaitGroup)

	for {
		token, err = r.decoder.Token()
		if token == json.Delim('}') {
			break
		}

		switch tt := token.(type) {
		case string:
			key = tt
		default:
			r.logger.Fatalf("invalid token, expected 'string' got %T: %v", token, token)
		}

		wg.Add(1)
		err = r.decoder.Decode(parser)
		if err != nil {
			go callback(wg, nil, fmt.Errorf("unable to parse json: %w", err))
		} else {
			(parser.(domain.Ider)).SetId(key)
			go callback(wg, parser, nil)
		}

		wg.Wait()
	}
}

func (r *JsonFileReader) Close() {
	if err := r.file.Close(); err != nil {
		r.logger.Fatalf("unable to close file: %w", err)
	}
}
