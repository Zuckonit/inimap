package inimap

import (
	"testing"
)

func TestRead(t *testing.T) {
	testf0 := "nonexist.ini"
	testf1 := "test.ini"

	_, err0 := Read(testf0)
	if err0 == nil {
		t.Errorf("file %s not existed", testf0)
	}

	cfg, err1 := Read(testf1)
	if err1 != nil {
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
}
