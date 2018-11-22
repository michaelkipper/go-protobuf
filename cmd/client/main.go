package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"github.com/michaelkipper/go-protobuf/pkg/client"

	log "github.com/sirupsen/logrus"
)

func testLatencyMarshal(w *client.LatencyWrapper) {
	m := jsonpb.Marshaler{}
	s, err := m.MarshalToString(&w.Latency)
	if err != nil {
		panic(err)
	}
	log.WithField("val", s).Info("With jsonpb.Marshal")

	b, err := json.Marshal(w.Latency)
	if err != nil {
		panic(err)
	}
	log.WithField("val", string(b)).Info("With json.Marshal")
}

func testLatencyWrapperMarshal(w *client.LatencyWrapper) []byte {
	b, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}
	log.WithField("val", string(b)).Info("With json.Marshal")
	return b
}

func testLatencyWrapperUnmarshal(w *client.LatencyWrapper) {
	b := testLatencyWrapperMarshal(w)
	ret := client.LatencyWrapper{}
	json.Unmarshal(b, &ret)
	log.WithField("ret", fmt.Sprintf("%+v", ret)).Info("With json.Unmarshal")
}

func main() {
	w := client.LatencyWrapper{
		Meta: client.Metadata{
			Name: "Michael",
		},
		Latency: client.Latency{
			LatencyType: &client.Latency_Normal{
				Normal: &client.NormalLatency{
					Mean:   123,
					Stddev: 456,
				},
			},
		},
	}

	testLatencyMarshal(&w)
	testLatencyWrapperUnmarshal(&w)
}
