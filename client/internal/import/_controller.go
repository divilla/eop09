package importer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/divilla/eop09/internal/domain"
	interfaces2 "github.com/divilla/eop09/internal/interfaces"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

type controller struct {
	logger interfaces2.Logger
}

func Controller(e *echo.Echo) {
	ctrl := &controller{
		logger: e.Logger,
		//ser: &service{},
	}

	e.GET("/import", ctrl.importer)
}

func (c *controller) importer(ctx echo.Context) error {
	var err error

	fileName := "data/ports.json"
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("file '%s' not found: %s", fileName, err.Error())
	}

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	token, err := decoder.Token()
	fmt.Println(token)
	if err == io.EOF {
		return fmt.Errorf("file '%s' is empty", fileName)
	}
	if token != json.Delim('{') {
		return fmt.Errorf("file '%s' is not valid json file", fileName)
	}
	if err != nil {
		c.logger.Fatal(err)
	}

	//client := gclient.Persistence()
	//stream, err := client.Imp(r.Context())
	//if err != nil {
	//	logger.Fatalf("client.Imp failed to register stream: %v", err)
	//}

	//comp, err := zlib.NewWriterLevel(&cbuf, zlib.BestSpeed )
	//if err != nil {
	//	logger.Fatal("Failed to init zlib writer")
	//}

	var port domain.Port
	for {
		//var key string
		token, err = decoder.Token()
		if token == json.Delim('}') {
			break
		}

		//switch tt := token.(type) {
		//case string:
		//	key = tt
		//default:
		//	c.logger.Fatalf("Invalid token. Expected 'string' got %T: %v", token, token)
		//}

		//decoder.More()
		err = decoder.Decode(&port)
		if err != nil {
			c.logger.Errorf("Unable to parse json: ", err)
		}
		//port.Id = key

		//c.logger.Info(port)

		//var buf bytes.Buffer
		//encoder := gob.NewEncoder(&buf)
		//
		//if err = encoder.Encode(port); err != nil {
		//	panic(err)
		//}

		//entity := &pb.Entity{
		//	Payload: buf.Bytes(),
		//}
		//if err = stream.Send(entity); err != nil {
		//	logger.Errorf("%v.Send(%v) = %v", stream, entity, err)
		//}

		//inp := buf.Bytes()
		//ibuf := bytes.NewBuffer(inp)
		//dec := gob.NewDecoder(ibuf)
		//if err = dec.Decode(&port); err != nil {
		//	c.logger.Fatal(err)
		//}
		//
		//fmt.Println(&port)
		//enc := json.NewEncoder(w)
		//enc.SetIndent("", "  ")
		//err = enc.Encode(port)
		//if err != nil {
		//	logger.LogError(err)
		//	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		//}
	}
	if err = file.Close(); err != nil {
		return fmt.Errorf("unable to close file: %w", err)
	}
	//reply, err := stream.CloseAndRecv()
	//if err != nil {
	//	c.logger.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	//}
	//c.logger.Printf("Route summary: %v", reply)

	return ctx.NoContent(http.StatusOK)
}
