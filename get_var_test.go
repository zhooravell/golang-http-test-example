package main

import (
	"log"
	"os"
	"testing"
)

func TestGetVarDefaultValue(t *testing.T) {
	defaultValue := "--TEST--"
	variableName := "OS_ENV_TEST_VARIABLE"

	if getVar(variableName, defaultValue) != defaultValue {
		t.Fail()
	}
}

func TestGetVarValueFromENV(t *testing.T) {
	defaultValue := "--TEST--"
	variableName := "OS_ENV_TEST_VARIABLE"

	originValue := "SUPPER+TEST"
	if err := os.Setenv(variableName, originValue); err != nil {
		log.Fatal(err)
	}

	if getVar(variableName, defaultValue) != originValue {
		t.Fail()
	}
}
