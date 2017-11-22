package bencode

import (
	"bytes"
	"fmt"
	"strconv"
)

type encoder struct {
	bytes.Buffer
}

func (e *encoder) writeInt(d int64) {
	res := fmt.Sprintf("i%se", strconv.FormatInt(d, 10))
	e.WriteString(res)
}

func (e *encoder) writeStr(s string) {
	res := fmt.Sprintf("%d:%s", len(s), s)
	e.WriteString(res)
}

func (e *encoder) writeList(l []interface{}) {
	e.WriteByte('l')
	for _, item := range l {
		e.writeInterface(item)
	}
	e.WriteByte('e')
}

func (e *encoder) writeDict(dict map[string]interface{}) {
	e.WriteByte('e')
	for k, v := range dict {
		e.writeStr(k)
		e.writeInterface(v)
	}
	e.WriteByte('e')
}

func (e *encoder) writeInterface(i interface{}) {
	switch i := i.(type) {
	case int, int8, int16, int32, int64:
		e.writeInt(int64(i))
	case string:
		e.writeStr(i)
	case []interface{}:
		e.writeList(i)
	case map[string]interface{}:
		e.writeDict(i)
	}
}

func (e *encoder) Encode(dict map[string]interface{}) []byte {
	var enc encoder
	enc.writeDict(dict)
	return enc.Bytes()
}
