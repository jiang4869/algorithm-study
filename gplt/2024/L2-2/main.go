package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var n int
	fmt.Fscanf(in, "%d\n", &n)
	mmid := make(map[string][]string)
	for i := 0; i < n; i++ {
		readLine, _, _ := in.ReadLine()
		line := string(readLine)
		splits := getSplits(line)
		str := ""
		for j := 0; j < len(splits); j++ {
			str += string(splits[j][0])
		}

		val, _ := mmid[str]
		val = append(val, line)
		mmid[str] = val
	}

	var m int
	fmt.Fscanf(in, "%d\n", &m)
	for i := 0; i < m; i++ {
		readLine, _, _ := in.ReadLine()
		line := string(readLine)
		splits := getSplits(line)
		str := ""
		for j := 0; j < len(splits); j++ {
			str += string(splits[j][0])
		}

		val, ok := mmid[str]
		sort.Strings(val)
		if ok {
			for j := 0; j < len(val); j++ {
				if j != 0 {
					fmt.Fprint(out, "|")
				}
				fmt.Fprint(out, val[j])
			}
			fmt.Fprintln(out, "")
		} else {
			fmt.Fprintln(out, line)
		}
	}

}

func getSplits(line string) []string {
	arr := make([]string, 0)
	begin := 0
	end := 0
	for end < len(line) {
		if line[end] == ' ' || end == len(line)-1 {
			if end == len(line)-1 {
				end++
			}
			str := line[begin:end]
			str = strings.TrimSpace(str)
			if len(str) > 0 {
				arr = append(arr, str)
			}
			begin = end + 1
		}
		end++
	}
	return arr
}
