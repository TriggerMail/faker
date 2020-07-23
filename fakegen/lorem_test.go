package fakegen

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestDataFaker(t *testing.T) {
	SetDataFaker(Lorem{})
}

func TestWord(t *testing.T) {
	word, err := GetLorem().Word(context.Background(), reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	if !Contains(wordList, word.(string)) {
		t.Error("Expected word from slice wordList")
	}
}

func TestSentence(t *testing.T) {
	res, err := GetLorem().Sentence(context.Background(), reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	s := res.(string)
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected sentence")
	}
}

func TestParagraph(t *testing.T) {
	res, err := GetLorem().Paragraph(context.Background(), reflect.Value{})
	if err != nil {
		t.Error("Expected  not error, got err", err)
	}
	s := res.(string)
	fmt.Println(s)
	if s == "" || !strings.HasSuffix(s, ".") {
		t.Error("Expected paragraph")
	}
}

func TestFakeWord(t *testing.T) {
	word := Word()
	if !Contains(wordList, word) {
		t.Error("Expected word from slice wordList")
	}
}

func TestFakeSentence(t *testing.T) {
	res := Sentence()
	if res == "" || !strings.HasSuffix(res, ".") {
		t.Error("Expected sentence")
	}
}

func TestFakeParagraph(t *testing.T) {
	res := Paragraph()
	if res == "" || !strings.HasSuffix(res, ".") {
		t.Error("Expected paragraph")
	}
}
