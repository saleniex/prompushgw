package prompushgw

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValueList_Add(t *testing.T) {
	kv := NewKeyValueList("prefix_")
	kv.Add(KeyValue{"_K01_", "_V01_"})
	kv.Add(KeyValue{"_K02_", "_V02_"})
	assert.Len(t, kv.Items, 2)
	assert.Equal(t, kv.Items[0].Key, "prefix__K01_")
	assert.Equal(t, kv.Items[1].Key, "prefix__K02_")
}

func TestKeyValueList_AddInt(t *testing.T) {
	kv := NewKeyValueList("prefix_")
	kv.AddInt("_K01_", 123)
	kv.AddInt("_K02_", 234)
	assert.Len(t, kv.Items, 2)
	assert.Equal(t, kv.Items[0].Val, "123")
	assert.Equal(t, kv.Items[1].Val, "234")
}

func TestKeyValueList_AddStr(t *testing.T) {
	kv := NewKeyValueList("prefix_")
	kv.AddStr("_K01_", "_V01_")
	kv.AddStr("_K02_", "_V02_")
	assert.Len(t, kv.Items, 2)
	assert.Equal(t, kv.Items[0].Key, "prefix__K01_")
	assert.Equal(t, kv.Items[1].Key, "prefix__K02_")
}

func TestKeyValueList_AsText(t *testing.T) {
	kv := NewKeyValueList("prefix_")
	kv.Add(KeyValue{"_K01_", "_V01_"})
	kv.AddInt("_K02_", 234)
	kv.AddStr("_K03_", "_V03_")
	assert.Equal(t,
		"prefix__K01_ _V01_\nprefix__K02_ 234\nprefix__K03_ _V03_\n",
		string(kv.AsText()))
}
