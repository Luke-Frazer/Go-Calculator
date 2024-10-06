package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetInput_Success(t *testing.T) {
	input := "hello\n"
	reader := strings.NewReader(input)
	result, err := getInput(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result != "hello" {
		t.Fatalf("Expected 'hello', got %v", result)
	}
}

func TestGetInput_NoInput(t *testing.T) {
	input := "\n"
	reader := strings.NewReader(input)
	result, err := getInput(reader)

	if err != errNoChar {
		t.Fatalf("Expected errNoChar, got %v", err)
	}
	if result != "" {
		t.Fatalf("Expected empty result, got %v", result)
	}
}

func TestGetInput_MultipleInputs(t *testing.T) {
	input := "hello world\n"
	reader := strings.NewReader(input)
	result, err := getInput(reader)

	if err != errSingleChar {
		t.Fatalf("Expected errSingleChar, got %v", err)
	}
	if result != "" {
		t.Fatalf("Expected empty result, got %v", result)
	}
}

func TestGetInput_FailedReading(t *testing.T) {
	input := ""
	reader := strings.NewReader(input)
	result, err := getInput(reader)

	if err != errFailedRead {
		t.Fatalf("Expected errFailedRead, got %v", err)
	}
	if result != "" {
		t.Fatalf("Expected empty result, got %v", result)
	}
}

func TestGetInputs_SuccessSingle(t *testing.T) {
	input := "hello\n"
	reader := strings.NewReader(input)
	result, err := getInputs(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := []string{"hello"}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestGetInputs_SuccessMulti(t *testing.T) {
	input := "hello world\n"
	reader := strings.NewReader(input)
	result, err := getInputs(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := []string{"hello", "world"}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestGetInputs_NoInput(t *testing.T) {
	input := "\n"
	reader := strings.NewReader(input)
	result, err := getInputs(reader)

	if err != errNoChar {
		t.Fatalf("Expected errNoChar, got %v", err)
	}
	if result != nil {
		t.Fatalf("Expected nil, got %v", result)
	}
}

func TestGetInputs_FailedReading(t *testing.T) {
	input := ""
	reader := strings.NewReader(input)
	result, err := getInputs(reader)

	if err != errFailedRead {
		t.Fatalf("Expected errFailedRead, got %v", err)
	}
	if result != nil {
		t.Fatalf("Expected nil result, got %v", result)
	}
}

func TestConvertToNumbers_SuccessMulti(t *testing.T) {
	result, err := convertToNumbers("1", "2", "3", "4")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestConvertToNumbers_SuccessSingle(t *testing.T) {
	result, err := convertToNumbers("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expected := []int{1}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Expected %v, got %v", expected, result)
	}
}

func TestConvertToNumbers_Failed(t *testing.T) {
	result, err := convertToNumbers("dfa$(#@)")
	if err == nil {
		t.Fatalf("Expected invalid syntax error, got %v", err)
	}
	if result != nil {
		t.Fatalf("Expected nil, got %v", result)
	}
}

func TestConvertToNumbers_SuccessNoInput(t *testing.T) {
	result, err := convertToNumbers()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatalf("Expected nil result, got %v", result)
	}
}

func TestAddValues_SuccessTwo(t *testing.T) {
	result := addValues(1, 2)
	expected := int(3)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestAddValues_SuccessOne(t *testing.T) {
	result := addValues(1)
	expected := int(1)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestAddValues_SuccessNone(t *testing.T) {
	result := addValues()
	expected := int(0)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestAddValues_SuccessNegativeNumbers(t *testing.T) {
	result := addValues(-1, -2, -3)
	expected := int(-6)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestAddValues_SuccessMixedNumbers(t *testing.T) {
	result := addValues(-1, 2, -3, 4)
	expected := int(2)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestAddValues_SuccessWithZero(t *testing.T) {
	result := addValues(0, 1, 2, 3)
	expected := int(6)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestSubValues_SuccessTwo(t *testing.T) {
	result := subValues(1, 2)
	expected := int(-1)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestSubValues_SuccessOne(t *testing.T) {
	result := subValues(1)
	expected := int(1)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestSubValues_SuccessNone(t *testing.T) {
	result := subValues()
	expected := int(0)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestSubValues_SuccessNegativeNumbers(t *testing.T) {
	result := subValues(-1, -2, -3)
	expected := int(4)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestSubValues_SuccessMixedNumbers(t *testing.T) {
	result := subValues(-1, 2, -3, 4)
	expected := int(-4)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestSubValues_SuccessWithZero(t *testing.T) {
	result := subValues(0, 1, 2, 3)
	expected := int(-6)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestMultValues_SuccessTwo(t *testing.T) {
	result := multValues(1, 2)
	expected := int(2)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestMultValues_SuccessOne(t *testing.T) {
	result := multValues(1)
	expected := int(1)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestMultValues_SuccessNone(t *testing.T) {
	result := multValues()
	expected := int(0)
	if result != expected {
		t.Fatalf("Expected %d, got %v", expected, result)
	}
}

func TestMultValues_SuccessNegativeNumbers(t *testing.T) {
	result := multValues(-1, -2, -3)
	expected := int(-6)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestMultValues_SuccessMixedNumbers(t *testing.T) {
	result := multValues(-1, 2, -3, 4)
	expected := int(24)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}

func TestMultValues_SuccessWithZero(t *testing.T) {
	result := multValues(0, 1, 2, 3)
	expected := int(0)
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
}
