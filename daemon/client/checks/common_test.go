package checks

import (
	"io/ioutil"
	"log"
	"testing"
)

func init() {
	// Omit standard log output when running tests to allow one to focus on
	// actual test results.
	log.SetOutput(ioutil.Discard)
}

func TestIsVgSizeOk(t *testing.T) {
	tests := []struct {
		line   string
		okSize int
		want   bool
	}{
		{"invalid input", 99, false},
		{"5.37 26.84 vg_slow", 5, true},
		{"5.37 26.84 vg_slow", 25, false},
		{"      0 511.03 fedora", 10, false},
		{"\t25\t250 test", 10, true},
		{"10G 50G test", 15, true},
		{"10G 50G test", 25, false},
	}
	for _, tt := range tests {
		if got := isVgSizeOk(tt.line, tt.okSize); got != tt.want {
			t.Errorf("isVgSizeOk(%q, %v) = %v, want %v", tt.line, tt.okSize, got, tt.want)
		}
	}
}
