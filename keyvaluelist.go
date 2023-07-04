package prompushgw

import (
	"fmt"
	"strings"
)

type KeyValueList struct {
	keyPrefix string
	Items     []KeyValue
}

func NewKeyValueList(keyPrefix string) *KeyValueList {
	return &KeyValueList{
		Items:     make([]KeyValue, 0),
		keyPrefix: keyPrefix,
	}
}

func (l *KeyValueList) Add(keyVal KeyValue) {
	keyVal.Key = l.keyPrefix + keyVal.Key
	l.Items = append(l.Items, keyVal)
}

func (l *KeyValueList) AddInt(key string, val int) {
	l.Add(KeyValue{
		Key: key,
		Val: fmt.Sprintf("%d", val),
	})
}

func (l *KeyValueList) AddStr(key, val string) {
	l.Add(KeyValue{
		Key: key,
		Val: val,
	})
}

func (l *KeyValueList) AsText() []byte {
	var builder strings.Builder
	for _, item := range l.Items {
		builder.WriteString(fmt.Sprintf("%s %s\n", item.Key, item.Val))
	}
	return []byte(builder.String())
}
