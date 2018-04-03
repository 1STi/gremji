package gremji

import (
	"github.com/satori/go.uuid"
	"encoding/json"
)

type Request struct {
	RequestId string       `json:"requestId"`
	Op        string       `json:"op"`
	Processor string       `json:"processor"`
	Args      *RequestArgs `json:"args"`
}

type RequestArgs struct {
	Gremlin           string            `json:"gremlin,omitempty"`
	Session           string            `json:"session,omitempty"`
	Bindings          Bind              `json:"bindings,omitempty"`
	Language          string            `json:"language,omitempty"`
	Accept            string            `json:"accept,omitempty"`
	Rebindings        Bind              `json:"rebindings,omitempty"`
	Sasl              string            `json:"sasl,omitempty"`
	BatchSize         int               `json:"batchSize,omitempty"`
	ManageTransaction bool              `json:"manageTransaction,omitempty"`
	Aliases           map[string]string `json:"aliases,omitempty"`
}

type Bind map[string]interface{}

type QueryArgs struct {
	Query      string
	Bindings   Bind
	Rebindings Bind
}

func Query(query QueryArgs) *Request {
	args := &RequestArgs{
		Gremlin:    query.Query,
		Language:   "gremlin-groovy",
		Accept:     "application/vnd.gremlin-v2.0+json",
		Bindings:   query.Bindings,
		Rebindings: query.Rebindings,
	}

    uid, err := uuid.NewV4()

	if err != nil {
		return nil
	}

	id := uid.String()

	req := &Request{
		RequestId:  id,
		Op:        "eval",
		Processor: "",
		Args:      args,
	}

	return req
}

// Formats the requests in the appropriate way
type FormattedReq struct {
	Op        string       `json:"op"`
	RequestId interface{}  `json:"requestId"`
	Args      *RequestArgs `json:"args"`
	Processor string       `json:"processor"`
}

func GraphSONSerializer(req *Request) ([]byte, error) {
	form := NewFormattedReq(req)
	msg, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	mimeType := []byte("application/vnd.gremlin-v2.0+json")
	var mimeLen= []byte{0x21}
	res := append(mimeLen, mimeType...)
	res = append(res, msg...)
	return res, nil
}

func NewFormattedReq(req *Request) FormattedReq {
	rId := map[string]string{"@type": "g:UUID", "@value": req.RequestId}
	sr := FormattedReq{RequestId: rId, Processor: req.Processor, Op: req.Op, Args: req.Args}

	return sr
}