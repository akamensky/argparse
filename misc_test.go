package argparse

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestIsNilFile(t *testing.T) {
	var b bool
	b = IsNilFile(nil)
	// nil != &{nil}
	if b {
		t.Errorf("Expected [%v], got [%v]", "false", b)
	}

	var fd os.File
	b = IsNilFile(&fd)
	// &{nil} != &{nil}
	if !b {
		t.Errorf("Expected [%v], got [%v]", "true", !b)
	}

	fdp, _ := ioutil.TempFile(os.TempDir(), "test")
	b = IsNilFile(fdp)
	// *os.File != &{nil}
	if b {
		t.Errorf("Expected [%v], got [%v]", "false", !b)
	}
}
