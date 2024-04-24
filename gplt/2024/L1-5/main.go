package main

import (
	"bufio"
	"fmt"
	"os"
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
	strs := make([]string, n)
	for i := 0; i < n; i++ {
		readLine, _, _ := in.ReadLine()
		strs[i] = string(readLine)
	}

	var k int
	fmt.Fscanf(in, "%d\n", &k)
	line := ""
	// 这题输入很迷，题目描述是只有一行字符串，如果直接用readLine, _, err := in.ReadLine()这么读取，没有加for循环，就过不去
	for {
		readLine, _, err := in.ReadLine()
		if err != nil {
			break
		}
		line += string(readLine)
	}

	count := 0
	builder := strings.Builder{}
	for i := 0; i < n; i++ {
		begin := 0
		for begin < len(line) {
			index := strings.Index(line[begin:], strs[i])
			if index == -1 {
				builder.Write([]byte(line[begin : begin+1]))
				begin++
			} else {
				count++
				builder.Write([]byte(line[begin : index+begin]))
				builder.Write([]byte("***"))
				begin = index + begin + len(strs[i])
			}
		}
		line = builder.String()
		builder.Reset()
	}
	line = strings.ReplaceAll(line, "***", "<censored>")
	if count >= k {
		fmt.Fprintln(out, count)
		fmt.Fprintln(out, "He Xie Ni Quan Jia!")
	} else {
		fmt.Fprintln(out, line)
	}

}
