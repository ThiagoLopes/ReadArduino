package model

import (
	"log"
	"testing"
)

func TestNewSerialData(t *testing.T) {
	test_data := []byte("33.3,44.33,444.44,33.3,55.55")

	out := NewSerialData(test_data)
	log.Println(out)
}
