package jsondb

import (
	"fmt"
	"testing"
)

type FakeThing struct {
	Number int `json:"name"`
}

func (f *FakeThing) String() string {
	return fmt.Sprintf("%d", f.Number)
}

func TestGetTableFile(t *testing.T) {
	fname := getTableFile("products")

	t.Log(fname)
}

func TestWriteTable(t *testing.T) {
	WriteTable("fakething", []*FakeThing{
		&FakeThing{1},
		&FakeThing{2},
		&FakeThing{3},
		&FakeThing{4},
		&FakeThing{5},
		&FakeThing{6},
		&FakeThing{7},
		&FakeThing{8},
		&FakeThing{9},
		&FakeThing{10},
		&FakeThing{11},
		&FakeThing{12},
		&FakeThing{13},
	})
}

func TestLoadTable(t *testing.T) {
	things := make([]*FakeThing, 0, 100)
	err := LoadTable("fakething", &things)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(things)
}
