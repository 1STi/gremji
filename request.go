package gremji

import (
    "github.com/satori/go.uuid"
)

type Request struct {
    RequstId  interface{}  `json:"requestId"`
    Op        string       `json:"op"`
    Processor string       `json:"processor"`
    Args      *RequestArgs `json:"args"`
}

type RequestArgs struct {
    Gremlin           string            `json:"gremlin,omitempty"`
    Session           string            `json:"session,omitempty"`
    Bindings          Bind              `json:"bindings,omitempty"`
    Language          string            `json:"language,omitempty"`
    Rebindings        Bind              `json:"rebindings,omitempty"`
    Sasl              string            `json:"sasl,omitempty"`
    BatchSize         int               `json:"batchSize,omitempty"`
    ManageTransaction bool              `json:"manageTransaction,omitempty"`
    Aliases           map[string]string `json:"aliases,omitempty"`
}

type Bind map[string]interface{}

type QueryArgs struct {
    query string
    bindings Bind
    rebindings Bind
}

func Query(query QueryArgs) *Request {
    args := &RequestArgs{
        Gremlin: query.query,
        Language: "gremlin-groovy",
        Bindings: query.bindings,
        Rebindings: query.rebindings,
    }

    id := uuid.NewV4().String()

    rId := map[string]string{"@type": "g:UUID", "@value": id}

    req := &Request{
        RequstId: rId,
        Op: "eval",
        Processor: "",
        Args: args,
    }

    return req
}