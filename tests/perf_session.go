package tests

import (
	"github.com/go-openapi/strfmt"
	"time"
)

type PerfSession struct {
}

func (self *PerfSession) CheckBool(expected bool, actual bool, message string) {
	if expected != actual {
		panic(message)
	}
}
func (self *PerfSession) CheckInt(expected int, actual int, message string) {
	if expected != actual {
		log.Panicf("Unexpected data (%d != %d): %s", expected, actual, message)
	}
}
func (self *PerfSession) CheckInt32(expected int32, actual int32, message string) {
	if expected != actual {
		log.Panicf("Unexpected data (%d != %d): %s", expected, actual, message)
	}
}
func (self *PerfSession) CheckInt64(expected int64, actual int64, message string) {
	if expected != actual {
		log.Panicf("Unexpected data (%d != %d): %s", expected, actual, message)
	}
}
func (self *PerfSession) CheckStr(expected string, actual string, message string) {
	if expected != actual {
		log.Panicf("Unexpected string value (%s != %s: %s", expected, actual, message)
	}
}
func (self *PerfSession) CheckHash(expected PHash, actual string, message string) {
	if expected != Hash(actual) {
		log.Panicf("Unexpected string (%s): %s", actual, message)
	}
}
func (self *PerfSession) CheckDate(expected *strfmt.DateTime, actual *strfmt.DateTime, message string) {
	expectedUtc := time.Time(*expected).UTC()
	actualUtc := time.Time(*actual).UTC()
	if !expectedUtc.Equal(actualUtc) {
		log.Panicf("Unexpected data (%v != %v): %s", expectedUtc.String(), actualUtc.String(), message)
	}
}
func (self *PerfSession) CheckVersion(before PVersion, after PVersion) bool {
	return true
}
func (self *PerfSession) Finish(before PVersion, after PVersion) {

}
