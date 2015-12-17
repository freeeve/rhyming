package main

import "testing"

func TestAppendBytes(t *testing.T) {
	bytes := []byte{}
	AppendBytes(bytes, 32)
}

func TestDecodeBytes(t *testing.T) {
	bytes := []byte{}
	expected := []int32{32, 33, 1337, 133337, 133339}
	for _, x := range expected {
		bytes = AppendBytes(bytes, x)
	}
	ids := DecodeBytes(bytes)
	if len(ids) != len(expected) {
		t.Errorf("length of ids array incorrect: %d", len(ids))
		return
	}
	for i, _ := range expected {
		if ids[i] != expected[i] {
			t.Errorf("ids[%d] was %d not %d", i, ids[i], expected[i])
		}
	}
}
