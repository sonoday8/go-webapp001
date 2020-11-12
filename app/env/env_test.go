package env

import (
	"os"
	"testing"
)

func TestEnvSuccess(t *testing.T) {
	os.Setenv("test", "true")
	if res := GetBool("test", false); res != true {
		t.Fatal("failed test")
	}
	if res := GetBool("dummy", false); res != false {
		t.Fatal("failed test")
	}
	if res := GetStr("test", "false"); res != "true" {
		t.Fatal("failed test")
	}
	if res := GetStr("dummy", "false"); res != "false" {
		t.Fatal("failed test")
	}
	os.Setenv("test", "1")
	if res := GetInt("test", 0); res != 1 {
		t.Fatal("failed test")
	}
	if res := GetInt("dummy", 0); res != 0 {
		t.Fatal("failed test")
	}
}
