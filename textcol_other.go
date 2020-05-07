// +build !windows

package textcol

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTermWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err1 := cmd.Output()
	check(err1)
	numsStr := strings.Trim(string(out), "\n ")
	width, err2 := strconv.Atoi(strings.Split(numsStr, " ")[1])
	check(err2)
	return width
}
