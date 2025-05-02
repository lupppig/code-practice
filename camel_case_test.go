package main

import "testing"

func TestToCamelCase(t *testing.T) {
	var str, expected, result string

	t.Run("test camel case 1", func(t *testing.T) {
		str = "the-stealth-warrior"
		expected = "theStealthWarrior"
		result = ToCamelCase(str)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}

	})
	t.Run("test camel case 2", func(t *testing.T) {
		str = "The_Stealth_Warrior"
		expected = "TheStealthWarrior"
		result = ToCamelCase(str)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}

	})

	t.Run("test camel case 3", func(t *testing.T) {
		str = "the-stealth-warrior"
		expected = "theStealthWarrior"
		result = ToCamelCase(str)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}

	})
	t.Run("test camel case 4", func(t *testing.T) {
		str = "the_stealth-warrior"
		expected = "theStealthWarrior"
		result = ToCamelCase(str)
		if result != expected {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	})

}
