package common

import (
	"fmt"
	"math/rand"
	"regexp"
	"runtime"
	"time"
)

const (
	UniqueCode = "unique_code"
)

// MakeTimestampMilli returns current ms timestamp
func MakeTimestampMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func RemoveNonAlphanumeric(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func Str2ptr(i string) *string {
	return &i
}

func Int2prt(i int32) *int32 {
	return &i
}

func Bool2prt(i bool) *bool {
	return &i
}

func RandomCharacters(size int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, size)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func MemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	result := fmt.Sprintf("memoryusage::Alloc = %v MB::TotalAlloc = %v MB::Sys = %v MB::tNumGC = %v", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
	return result
}

// useful for getting how much ram we are currently using in case we ever want that
func AllocatedMemUsageMb() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return bToMb(m.Alloc)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
