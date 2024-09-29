package main

import (
	"reflect"
	"testing"

	"github.com/MadhavKrishanGoswami/Cash-CLI/currency"
	"github.com/MadhavKrishanGoswami/Cash-CLI/models"
)

func TestConvertor(t *testing.T) {
	inp := models.Input{
		"INR",
		"INR",
		"1",
	}
	out_check := models.Output{
		1,
		1,
	}
	out := currency.Convertor(inp)
	if !reflect.DeepEqual(out, out_check) {
		t.Error("Convertor Is not working")
	}
}
