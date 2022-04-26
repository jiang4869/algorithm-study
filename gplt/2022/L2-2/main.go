package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in *bufio.Reader = bufio.NewReader(os.Stdin)
var out *bufio.Writer = bufio.NewWriter(os.Stdout)

type Time struct {
	Begin string
	End   string
}

func main() {
	defer out.Flush()
	var n int
	fmt.Fscanf(in, "%d\n", &n)
	var str1, str2 string
	arr := make([]Time, 0)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%s - %s\n", &str1, &str2)
		arr = append(arr, Time{Begin: str1, End: str2})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Begin < arr[j].Begin
	})

	cur := "00:00:00"
	for i := 0; i < n; i++ {
		if arr[i].Begin != cur {
			fmt.Fprintln(out, cur, "-", arr[i].Begin)
		}
		cur = arr[i].End

	}
	if cur != "23:59:59" {
		fmt.Fprintln(out, cur, "-", "23:59:59")
	}

}
