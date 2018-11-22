package client

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/jsonpb"

	log "github.com/sirupsen/logrus"
)

type Metadata struct {
	Name string `json:"name,omitempty"`
}

type LatencyWrapper struct {
	Meta    Metadata `json:"meta"`
	Latency Latency  `json:"latency"`
}

func (lw *LatencyWrapper) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString("{")

	jsonValue, err := json.Marshal(lw.Meta)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"meta\": %s,", jsonValue))

	encoder := jsonpb.Marshaler{OrigName: false}
	jsonBytes, err := encoder.MarshalToString(&lw.Latency)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(fmt.Sprintf("\"%s\":%s", "latency", string(jsonBytes)))

	buffer.WriteString("}")

	return buffer.Bytes(), nil
}

type LatencyWrapperContainer struct {
	Meta    Metadata `json:"meta"`
	Latency string   `json:"latency"`
}

func (lw *LatencyWrapper) UnmarshalJSON(b []byte) error {
	c := LatencyWrapperContainer{}
	log.WithFields(log.Fields{"buffer": string(b), "wrapper": lw}).Info("Unmarshalled")
	err := json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}
	log.WithFields(log.Fields{"wrapper": c}).Info("Unmarshalled")
	return nil
}
