package main

import (
	"log/slog"
	"testing"
)

var data = map[string]interface{}{
	"numDecimal":     42,
	"numOctal":       052,
	"numHexadecimal": 0x2A,
	"pi":             3.14,
	"name":           "Golang",
	"isActive":       true,
	"complexNum":     complex64(1 + 2i),
}

func TestGetType(t *testing.T) {
	log := slog.With(slog.String("op", "test"))
	for _, data := range data {
		getType(data, log)
	}

}

func MyStrokeBuilder(t *testing.T) {
	values := make([]interface{}, 0, len(data))
	for _, v := range data {
		values = append(values, v)
	}
	mystroke := strokeBuilder(values...)
	runeslice := convertToSliceRune(mystroke)
	HashedSliceRune(runeslice, "go-2025")
}
