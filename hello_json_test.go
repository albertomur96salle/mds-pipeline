package main

import (
	"testing"
)

func TestGetHelloJSON(t *testing.T) {
	t.Run("This test checks if the hello JSON class returns the proper JSON", func(t *testing.T) {
		assertEquals(`{"message": "Hola Mundo"}`, GetHelloJSON(), t)
	})
}

func assertEquals(expected string, result string, t *testing.T) {
	if result != expected {
		t.Errorf("wrong greeting. expected greeting => '%s', current greeting '%s'", expected, result)
	}
}
