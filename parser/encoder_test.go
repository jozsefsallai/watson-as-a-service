package parser_test

import (
	"testing"

	. "github.com/jozsefsallai/watson-as-a-service/parser"
)

type testcase struct {
	Title string
	Input string
	DT    DataType
	Want  string
}

func tc(title, input string, dataType DataType, want string) testcase {
	return testcase{
		Title: title,
		Input: input,
		DT:    dataType,
		Want:  want,
	}
}

func TestEncodeAs(t *testing.T) {
	testcases := []testcase{
		tc("should encode as integer", "123", DTInt, "Bububububbubu"),
		tc("should encode as float", "3.14", DTFloat, "Bubbbbbbbbbbbubbbubbbbububububbubbubububbbbbubbubbbbububububbubbubububbbbbubbubbbbububububui"),
		tc("should encode as string", "hi", DTString, "?Shahaahaaa-Shahaahaaah-"),
		tc("should encode as boolean", "true", DTBool, "zo"),
		tc("should encode as object", "{\"hello\": \"world\"}", DTObject, "~?Shahaahaaa-Shahaaahaah-Shahaahahaa-Shahaahahaa-Shahaahahahah-$Bubububbububu!Bububbubububu!Bubububbbub!Bububbububb!Bububbbubb!M"),
		tc("should encode as array", "[ 1, 2, \"hello\" ]", DTArray, "@BububububububububububbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbisBubbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbis?Shahaahaaa-Shahaaahaah-Shahaahahaa-Shahaahahaa-Shahaahahahah-?"),

		tc("should return 0 for non-integer", "a", DTInt, "B"),
		tc("should return 0.0 for non-float", "a", DTFloat, "Bi"),
		tc("should return false for non-boolean", "a", DTBool, "z"),
		tc("should return empty object for invalid object", "a", DTObject, "~"),
		tc("should return null for invalid array", "a", DTArray, "."),

		tc("should try to parse raw data on unknown data type", "a", 100, "@Bububbbbbu's"), // [97]
	}

	for _, tt := range testcases {
		t.Run(tt.Title, func(t *testing.T) {
			buf, err := EncodeAs([]byte(tt.Input), tt.DT)
			if err != nil {
				t.Fatalf(err.Error())
			}

			if string(buf) != tt.Want {
				t.Fatalf("EncodeAs(%s, <%s>), want %s, got %s", tt.Input, tt.DT.String(), tt.Want, string(buf))
			}
		})
	}
}
