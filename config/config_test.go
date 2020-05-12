package config

import "testing"

func TestConfigDevelopment(t *testing.T) {
	Init("development")
	config := GetConfig()

	debug := config.GetBool("app.debug")

	if debug != true {
		t.Errorf("app.debug was incorrect, got: %v, want: %v.", debug, true)
	}
}

func TestConfigProduction(t *testing.T) {
	Init("production")
	config := GetConfig()

	debug := config.GetBool("app.debug")

	if debug != false {
		t.Errorf("app.debug was incorrect, got: %v, want: %v.", debug, false)
	}
}
