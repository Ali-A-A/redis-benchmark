package parser

import "encoding/json"

// JsonParser represents the main interface of the parser
type JsonParser interface {
	encode(in interface{}) error
	decode() (interface{}, error)
}

// JsonParserImpl implements JsonParser interface
type JsonParserImpl struct {
	Data []byte
}

// NewJsonParser returns new JsonParserImpl
func NewJsonParser() *JsonParserImpl {
	return &JsonParserImpl{}
}

// Encode marshals the input
func (jp *JsonParserImpl) Encode(in interface{}) error {
	d, err := json.Marshal(in)
	if err != nil {
		return err
	}
	jp.Data = d
	return nil
}

// Decode decodes stored data
func (jp *JsonParserImpl) Decode() (interface{}, error) {
	var v interface{}
	err := json.Unmarshal(jp.Data, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// DataLen returns data length
func (jp *JsonParserImpl) DataLen() int {
	if jp == nil {
		panic("jp could not nil")
	} else {
		return len(jp.Data)
	}
}

