package argparse

import "strings"

func getLastLine(input string) string {
	slice := strings.Split(input, "\n")
	return slice[len(slice)-1]
}

func addToLastLine(base string, add string, width int, padding int, canSplit bool) string {
	// Split the multiline string to add
	if strings.Contains(add, "\n") {
		lines := strings.Split(add, "\n")
		firstLine := true
		for _, v := range lines {
			if firstLine {
				base = addToLastLine(base, v, width, padding, true)
				firstLine = false
			} else {
				base = base + "\n" + strings.Repeat(" ", padding)
				base = addToLastLine(base, v, width, padding, true)
			}
		}
		return base
	}
	// If last line has less than 10% space left, do not try to fill in by splitting else just try to split
	hasTen := (width - len(getLastLine(base))) > width/10
	if len(getLastLine(base)+" "+add) >= width {
		if hasTen && canSplit {
			adds := strings.Split(add, " ")
			for _, v := range adds {
				base = addToLastLine(base, v, width, padding, false)
			}
			return base
		}
		base = base + "\n" + strings.Repeat(" ", padding)
	}
	base = base + " " + add
	return base
}
