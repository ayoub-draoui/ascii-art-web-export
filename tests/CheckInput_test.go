package tests

import (
	"testing"

	"functions/functions"
)

func TestCheckInput(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{"Valid Case", "hello world", "hello world"},
		{"Valid Case", "L'été en Provence est rempli de parfums envoûtants et de couleurs éclatantes.", "L't en Provence est rempli de parfums envotants et de couleurs clatantes."},
		{"Valid Case", "héllè", "hll"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.CheckInput(tt.text)
			if result != tt.want {
				t.Errorf("correct result value is %v but your result was %v", result, tt.want)
			}
		})
	}
}
