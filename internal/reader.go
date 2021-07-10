package internal

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
)

type Reader struct {
	Filename string
	Channel  chan []string
}

func NewReader(filename string, channel chan []string) *Reader {
	return &Reader{filename, channel}
}

func (r *Reader) ReadCsvToChannel() {

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(path.Join(workdir, r.Filename))

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	defer close(r.Channel)

	headerParsed := false
	parser := csv.NewReader(f)

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if !headerParsed {
			headerParsed = true
			continue
		}
		r.Channel <- record
	}
}
