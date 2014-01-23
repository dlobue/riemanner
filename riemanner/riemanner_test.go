package riemanner

import (
	"bytes"
	"fmt"
	"github.com/amir/raidman"
	"testing"
)

type MockRaidmanClient struct {
	Events []raidman.Event
}

func (m *MockRaidmanClient) Send(event *raidman.Event) error {
	m.Events = append(m.Events, *event)
	return nil
}

func TestIt(t *testing.T) {
	var input bytes.Buffer
	input.Write([]byte(`{"Ttl":10,"Time":0,"Tags":[],"Host":"host","State":"success","Service":"foo","Metric":100,"Description":""}
{"Ttl":100,"Time":0,"Tags":["asdf"],"Host":"raidman","State":"success","Service":"bar","Metric":100,"Description":""}


`))
	mock := MockRaidmanClient{}

	riemanner := NewRiemanner(&mock, &input)
	if err := riemanner.Run(); err != nil {
		t.Fatal(err)
	}

	expected := []raidman.Event{
		raidman.Event{Ttl: 10, Time: 0, Tags: []string{}, Host: "host", State: "success", Service: "foo", Metric: 100, Description: ""},
		raidman.Event{Ttl: 100, Time: 0, Tags: []string{"asdf"}, Host: "raidman", State: "success", Service: "bar", Metric: 100, Description: ""},
	}

	expectedstr := fmt.Sprintf("%#v", expected)
	actualstr := fmt.Sprintf("%#v", mock.Events)

	if expectedstr != actualstr {
		t.Fatalf("Expected:\n%s\ngot:\n%s", expectedstr, actualstr)
	}
}
