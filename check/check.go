package check

import "fmt"

func check(e error) {
	if e != nil {
		fmt.Println("<cluster>/ yamls exists, please manually edit file to make changes or provide new cluster name")
		panic(e)
	}
}
