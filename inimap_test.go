package inimap

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	testf0 := "nonexist.ini"
	testf1 := "test.ini"

	_, err := ReadFile(testf0)
	if err == nil {
		t.Errorf("file %s not existed", testf0)
	}

	cfg, err := ReadFile(testf1)
	if err != nil {
		t.Errorf("read %s ini failed", testf1)
	}

	subcfg, ok := (*cfg)["test"]
	if !ok {
		t.Errorf("section test existed")
	}
	if subcfg["a"] != "1" || subcfg["b"] != "2" {
		t.Errorf("a, b existed")
	}

	subcfg, ok = (*cfg)["test2"]
	if !ok {
		t.Errorf("section test2 existed")
	}
	if subcfg["a"] != "1" || subcfg["b"] != "2" {
		t.Errorf("a, b existed")
	}

	_, ok = subcfg["c"]
	if ok {
		t.Errorf("c is a comment")
	}

	cfg, err = ReadFile("test2.ini")
	if err == nil || cfg != nil {
		t.Errorf("test2.ini is invalid")
	}
}
