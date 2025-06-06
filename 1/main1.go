package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log/slog"
	"reflect"
	"strings"
)

func main() {
	log := slog.With(slog.String("op", "main"))
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i
	getType(numDecimal, log)
	getType(numOctal, log)
	getType(numHexadecimal, log)
	getType(pi, log)
	getType(name, log)
	getType(isActive, log)
	getType(complexNum, log)
	strokearg := strokeBuilder(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	log.Info("strokeBuilder", "stroke", strokearg)
	runeArray := convertToSliceRune(strokearg)
	log.Info("convertToSliceRune", "slice", runeArray)
	log.Info("sha256", "sha256", HashedSliceRune(runeArray, "go-2025"))
}

func getType(v interface{}, log *slog.Logger) string {
	mytpye := reflect.TypeOf(v).String()
	log.Info("getType", "type", mytpye)
	return mytpye
}

func strokeBuilder(values ...interface{}) string {
	builder := strings.Builder{}
	for _, value := range values {
		builder.WriteString(fmt.Sprintf("%v", value))
	}
	return builder.String()
}

func convertToSliceRune(str string) []rune {
	return []rune(str)
}

func HashedSliceRune(runes []rune, salt string) string {
	mid := len(runes) / 2
	withSalt := append(runes[:mid], append([]rune(salt), runes[mid:]...)...)
	hash := sha256.Sum256([]byte(string(withSalt)))
	return hex.EncodeToString(hash[:])
}
