package argparse

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestFlagSimple1(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1"}

	p := NewParser("", "description")
	flag1 := p.Flag("", "flag-arg1", nil)
	flag2 := p.Flag("", "flag-arg2", nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if flag1 == nil {
		t.Errorf("Test %s failed with flag1 being nil pointer", t.Name())
		return
	}

	if flag2 == nil {
		t.Errorf("Test %s failed with flag2 being nil pointer", t.Name())
		return
	}

	if *flag1 != true {
		t.Errorf("Test %s failed with flag1 being false", t.Name())
		return
	}

	if *flag2 != false {
		t.Errorf("Test %s failed with flag2 being true", t.Name())
		return
	}
}

func TestFlagSimple2(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1", "-f"}

	p := NewParser("", "description")
	flag1 := p.Flag("", "flag-arg1", nil)
	flag2 := p.Flag("", "flag-arg2", nil)
	flag3 := p.Flag("f", "flag-arg3", nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if flag1 == nil {
		t.Errorf("Test %s failed with flag1 being nil pointer", t.Name())
		return
	}

	if flag2 == nil {
		t.Errorf("Test %s failed with flag2 being nil pointer", t.Name())
		return
	}

	if flag3 == nil {
		t.Errorf("Test %s failed with flag5 being nil pointer", t.Name())
		return
	}

	if *flag1 != true {
		t.Errorf("Test %s failed with flag1 being false", t.Name())
		return
	}

	if *flag2 != false {
		t.Errorf("Test %s failed with flag2 being true", t.Name())
		return
	}

	if *flag3 != true {
		t.Errorf("Test %s failed with flag3 being false", t.Name())
		return
	}
}

func TestFlagMultiShorthand1(t *testing.T) {
	testArgs := []string{"progname", "-abcd", "-e"}

	p := NewParser("", "description")
	flag1 := p.Flag("a", "aa", nil)
	flag2 := p.Flag("b", "bb", nil)
	flag3 := p.Flag("c", "cc", nil)
	flag4 := p.Flag("d", "dd", nil)
	flag5 := p.Flag("e", "ee", nil)
	flag6 := p.Flag("f", "ff", nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if *flag1 != true {
		t.Errorf("Test %s failed with flag1 being false", t.Name())
	}

	if *flag2 != true {
		t.Errorf("Test %s failed with flag2 being false", t.Name())
	}

	if *flag3 != true {
		t.Errorf("Test %s failed with flag3 being false", t.Name())
	}

	if *flag4 != true {
		t.Errorf("Test %s failed with flag4 being false", t.Name())
	}

	if *flag5 != true {
		t.Errorf("Test %s failed with flag5 being false", t.Name())
	}

	if *flag6 != false {
		t.Errorf("Test %s failed with flag6 being true", t.Name())
	}
}

func TestFailDuplicate(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1", "-f"}

	p := NewParser("", "description")
	_ = p.Flag("f", "flag-arg1", nil)
	_ = p.Flag("", "flag-arg2", nil)

	err := p.Parse(testArgs)
	if err == nil {
		t.Errorf("Test %s failed with. Duplicate flag use not detected", t.Name())
		return
	}
}

func TestFailCaseSensitive(t *testing.T) {
	testArgs := []string{"progname", "-F"}

	p := NewParser("", "description")
	_ = p.Flag("f", "", &Options{Required: true})

	err := p.Parse(testArgs)
	if err == nil {
		t.Errorf("Test %s failed with. Sees -F as -f", t.Name())
		return
	}
}

func TestFailExcessiveArguments(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1", "whatever"}

	p := NewParser("", "description")
	_ = p.Flag("f", "flag-arg1", nil)
	_ = p.Flag("", "flag-arg2", nil)

	err := p.Parse(testArgs)
	if err == nil {
		t.Errorf("Test %s failed with. Excessive argument not detected", t.Name())
		return
	}
}

func TestStringSimple1(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1", "test"}

	p := NewParser("", "description")
	s1 := p.String("f", "flag-arg1", nil)
	s2 := p.String("", "flag-arg2", nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if s1 == nil {
		t.Errorf("Test %s failed with flag1 being nil pointer", t.Name())
		return
	}

	if s2 == nil {
		t.Errorf("Test %s failed with flag2 being nil pointer", t.Name())
		return
	}

	if *s1 != "test" {
		t.Errorf("Test %s failed. Want: [%s], got: [%s]", t.Name(), "test", *s1)
		return
	}

	if *s2 != "" {
		t.Errorf("Test %s failed. Want: [%s], got: [%s]", t.Name(), "\"\"", *s1)
		return
	}
}

func TestStringSimple2(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1", "test"}

	p := NewParser("", "description")
	s1 := p.String("f", "flag-arg1", nil)
	s2 := p.String("", "flag-arg2", nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if s1 == nil {
		t.Errorf("Test %s failed with flag1 being nil pointer", t.Name())
		return
	}

	if s2 == nil {
		t.Errorf("Test %s failed with flag2 being nil pointer", t.Name())
		return
	}

	if *s1 != "test" {
		t.Errorf("Test %s failed. Want: [%s], got: [%s]", t.Name(), "test", *s1)
		return
	}

	if *s2 != "" {
		t.Errorf("Test %s failed. Want: [%s], got: [%s]", t.Name(), "\"\"", *s1)
		return
	}
}

func TestFileSimple1(t *testing.T) {
	// Test file location
	fpath := "./test.tmp"
	// Create test file
	f, err := os.Create(fpath)
	if err != nil {
		t.Error(err)
		return
	}
	f.Close()
	defer os.Remove(fpath)

	testArgs := []string{"progname", "-f", fpath}

	p := NewParser("", "")

	file1 := p.File("f", "file", os.O_RDWR, 0666, nil)

	err = p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}
	defer file1.Close()

	if file1 == nil {
		t.Errorf("Test %s failed with file1 being nil pointer", t.Name())
		return
	}

	testString := "Test"
	recSlice := make([]byte, 4)
	_, err = file1.WriteString(testString)
	if err != nil {
		t.Errorf("Test %s write operation failed with error: %s", t.Name(), err.Error())
		return
	}
	file1.Seek(0, 0)
	n, err := file1.Read(recSlice)
	if err != nil {
		t.Errorf("Test %s read operation failed with error: %s", t.Name(), err.Error())
		return
	}
	if n != 4 || string(recSlice) != testString {
		t.Errorf("Test %s failed on read operation", t.Name())
		return
	}
}

func TestListSimple1(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1", "test1", "--flag-arg1", "test2"}
	list1Expect := []string{"test1", "test2"}
	list2Expect := make([]string, 0)

	p := NewParser("", "description")
	l1 := p.List("f", "flag-arg1", nil)
	l2 := p.List("", "flag-arg2", nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if l1 == nil {
		t.Errorf("Test %s failed with l1 being nil pointer", t.Name())
		return
	}

	if l2 == nil {
		t.Errorf("Test %s failed with l2 being nil pointer", t.Name())
		return
	}

	if !reflect.DeepEqual(*l1, list1Expect) {
		t.Errorf("Test %s failed. Want: %s, got: %s", t.Name(), list1Expect, *l1)
		return
	}

	if !reflect.DeepEqual(*l2, list2Expect) {
		t.Errorf("Test %s failed. Want: %s, got: %s", t.Name(), list2Expect, *l2)
		return
	}
}

func TestSelectorSimple1(t *testing.T) {
	flag1Expect := "test2"
	allowedValues := []string{"test1", flag1Expect}
	testArgs := []string{"progname", "--flag-arg1", flag1Expect}

	p := NewParser("", "")
	s1 := p.Selector("f", "flag-arg1", allowedValues, nil)

	err := p.Parse(testArgs)
	if err != nil {
		t.Errorf("Test %s failed with error: %s", t.Name(), err.Error())
		return
	}

	if s1 == nil {
		t.Errorf("Test %s failed with s1 being nil pointer", t.Name())
		return
	}

	if *s1 != flag1Expect {
		t.Errorf("Test %s failed. Want: %s, got: %s", t.Name(), flag1Expect, *s1)
		return
	}
}

func TestSelectorFailSimple1(t *testing.T) {
	allowedValues := []string{"test1", "test2"}
	testArgs := []string{"progname", "--flag-arg1", "test3"}

	p := NewParser("", "")
	_ = p.Selector("f", "flag-arg1", allowedValues, nil)

	err := p.Parse(testArgs)
	if err == nil {
		t.Errorf("Test %s failed. Expected error did not happen", t.Name())
		return
	}
}

func TestCommandSimple1(t *testing.T) {
	testArgsList := [][]string{
		{"progname", "cmd1", "--flag1", "--string-flag1", "test"},
		{"progname", "cmd2"},
	}

	for _, testArgs := range testArgsList {
		p := NewParser("progname", "description")

		cmd1 := p.NewCommand("cmd1", "cmd1 description")
		flag1 := cmd1.Flag("f", "flag1", nil)
		string1 := cmd1.String("s", "string-flag1", nil)

		cmd2 := p.NewCommand("cmd2", "cmd2 description")

		p.Parse(testArgs)

		if cmd1.Happened() {
			if *flag1 != true {
				t.Errorf("Test %s failed with %s: flag1: wanted [true], got [false]", t.Name(), testArgs[1])
				return
			}
			if *string1 != "test" {
				t.Errorf("Test %s failed with %s: string1: wanted [test], got [%s]", t.Name(), testArgs[1], *string1)
				return
			}
		}
		if cmd2.Happened() {
			if *flag1 != false {
				t.Errorf("Test %s failed with %s: flag1: wanted [false], got [true]", t.Name(), testArgs[1])
				return
			}
			if *string1 != "" {
				t.Errorf("Test %s failed with %s: string1: wanted [], got [%s]", t.Name(), testArgs[1], *string1)
				return
			}
		}
		if (cmd1.Happened() && cmd2.Happened()) || (!cmd1.Happened() && !cmd2.Happened()) {
			t.Errorf("Test %s failed, either cmd1 and cmd2 or neither of them Happened()", t.Name())
			return
		}
	}
}

func TestCommandMixedArgs1(t *testing.T) {
	testArgsList := [][]string{
		{"progname", "cmd1", "--flag1", "--string-flag1", "test", "--global-flag", "--global-string", "global test string"},
		{"progname", "cmd2", "--global-string", "global test string", "--global-flag"},
	}

	for _, testArgs := range testArgsList {
		p := NewParser("progname", "description")

		cmd1 := p.NewCommand("cmd1", "cmd1 description")
		cmd1flag1 := cmd1.Flag("f", "flag1", nil)
		cmd1string1 := cmd1.String("s", "string-flag1", nil)

		cmd2 := p.NewCommand("cmd2", "cmd2 description")

		pflag1 := p.Flag("", "global-flag", nil)
		pstring1 := p.String("", "global-string", nil)

		p.Parse(testArgs)

		// Check global flags
		if *pflag1 != true {
			t.Errorf("Test %s failed with %s: pflag1: wanted [true], got [false]", t.Name(), testArgs[1])
			return
		}
		if *pstring1 != "global test string" {
			t.Errorf("Test %s failed with %s: pstring1: wanted [global test string], got [%s]", t.Name(), testArgs[1], *pstring1)
			return
		}

		// Check commands
		if cmd1.Happened() {
			if *cmd1flag1 != true {
				t.Errorf("Test %s failed with %s: flag1: wanted [true], got [false]", t.Name(), testArgs[1])
				return
			}
			if *cmd1string1 != "test" {
				t.Errorf("Test %s failed with %s: string1: wanted [test], got [%s]", t.Name(), testArgs[1], *cmd1string1)
				return
			}
		}
		if cmd2.Happened() {
			if *cmd1flag1 != false {
				t.Errorf("Test %s failed with %s: flag1: wanted [false], got [true]", t.Name(), testArgs[1])
				return
			}
			if *cmd1string1 != "" {
				t.Errorf("Test %s failed with %s: string1: wanted [], got [%s]", t.Name(), testArgs[1], *cmd1string1)
				return
			}
		}
		if (cmd1.Happened() && cmd2.Happened()) || (!cmd1.Happened() && !cmd2.Happened()) {
			t.Errorf("Test %s failed, either cmd1 and cmd2 or neither of them Happened()", t.Name())
			return
		}
	}
}

func TestOptsRequired1(t *testing.T) {
	testArgs := []string{"progname", "--flag-arg1"}

	p := NewParser("", "description")
	_ = p.Flag("", "flag-arg1", nil)
	_ = p.String("", "flag-arg2", &Options{Required: true})

	err := p.Parse(testArgs)
	if err == nil {
		t.Errorf("Test %s failed to detect required argument", t.Name())
		return
	}
}

var opts = &Options{Validate: func(args []string) error {
	if len(args) > 0 {
		if args[0] != "pass" {
			return errors.New("failure")
		}
	}
	return nil
},
}

func TestOptsValidatePass1(t *testing.T) {
	testArgsList := [][]string{
		{"progname", "--string-flag1", "pass"},
		{"progname", "--string-flag1", "fail"},
	}

	for _, testArgs := range testArgsList {
		p := NewParser("progname", "")

		string1 := p.String("", "string-flag1", opts)

		err := p.Parse(testArgs)

		if testArgs[2] == "pass" {
			if err != nil {
				t.Errorf("Test %s failed on %s with err: %s", t.Name(), testArgs[2], err.Error())
				return
			}

			if *string1 != "pass" {
				t.Errorf("Test %s failed on %s; string1 expected [%s], got [%s]", t.Name(), testArgs[2], testArgs[2], *string1)
				return
			}
		} else {
			if err == nil {
				t.Errorf("Test %s failed to validate argument (should return error)", t.Name())
				return
			}
		}
	}
}

var pUsage = `usage: verylongprogname <command> [-h|--help] [-s|--verylongstring-flag1 "<value>"]

                        prog description

Commands:

   veryverylongcmd1    cmd1 description
   cmd2                cmd2 description

Arguments:

   -h   --help                    Print help information
   -s   --verylongstring-flag1    string1 description

`

var cmd1Usage = `usage: verylongprogname veryverylongcmd1 [-f|--verylongflag1] -a|--verylongflagA [-h|--help]
                        [-s|--verylongstring-flag1 "<value>"]

                        cmd1 description

Arguments:

   -f   --verylongflag1           flag1 description
   -a   --verylongflagA           flag1 description
   -h   --help                    Print help information
   -s   --verylongstring-flag1    string1 description

`

var cmd2Usage = `usage: verylongprogname cmd2 [-h|--help] [-s|--verylongstring-flag1 "<value>"]

                        cmd2 description

Arguments:

   -h   --help                    Print help information
   -s   --verylongstring-flag1    string1 description

`

func TestUsageSimple1(t *testing.T) {
	p := NewParser("verylongprogname", "prog description")

	cmd1 := p.NewCommand("veryverylongcmd1", "cmd1 description")
	_ = cmd1.Flag("f", "verylongflag1", &Options{Help: "flag1 description"})
	_ = cmd1.Flag("a", "verylongflagA", &Options{Required: true, Help: "flag1 description"})
	_ = p.String("s", "verylongstring-flag1", &Options{Help: "string1 description"})

	cmd2 := p.NewCommand("cmd2", "cmd2 description")

	p.Parse(os.Args)

	if pUsage != p.Usage(nil) {
		t.Errorf("%s", p.Usage(nil))
	}
	if cmd1Usage != cmd1.Usage(nil) {
		t.Errorf("%s", cmd1.Usage(nil))
	}
	if cmd2Usage != cmd2.Usage(nil) {
		t.Errorf("%s", cmd2.Usage(nil))
	}
}
