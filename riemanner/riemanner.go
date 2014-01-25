package riemanner

import (
	"bufio"
	"encoding/json"
	"github.com/amir/raidman"
	"io"
	"log"
)

// RaidmanClient is the interface exposed by raidman.
type RaidmanClient interface {
	Send(event *raidman.Event) error
}

// Riemanner is an io.Writer that you can send events to as lines of json objects.
type Riemanner struct {
	rc    RaidmanClient
	input io.Reader
}

// NewRiemanner initializes a Riemanner.
func NewRiemanner(rc RaidmanClient, input io.Reader) Riemanner {
	return Riemanner{rc: rc, input: input}
}

// Run Takes json input and sends it to riemann
func (r *Riemanner) Run() error {

	scanner := bufio.NewScanner(r.input)

	for scanner.Scan() {

		event := &raidman.Event{}
		jsonErr := json.Unmarshal(scanner.Bytes(), &event)
		if jsonErr == nil {
			must(r.rc.Send(event))
		} else {
			log.Println("Skipping invalid JSON line:", jsonErr)
		}

	}
	return scanner.Err()
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
