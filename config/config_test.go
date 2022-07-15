package config

import (
	"reflect"
	"testing"
)

func TestDefaultConfig(t *testing.T) {

	got := GetConfig()
	want := &Config{
		Url:      "https://www.test.com",
		LogLevel: "DEBUG",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
