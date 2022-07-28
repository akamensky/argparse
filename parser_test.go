package argparse

import "testing"

func TestNew(t *testing.T) {
	p := New("test", `Test description`)
	if p == nil {
		t.Errorf("parser is nil")
	}
}

func TestParser_Parse(t *testing.T) {
	p := New("test", `Test description`)
	if p == nil {
		t.Errorf("parser is nil")
	}

	err := p.Parse([]string{})
	if err != nil {
		t.Errorf("Failed to parse empty argument list")
	}
}
