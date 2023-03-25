package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 1 {
// 		fmt.Println("Usage : " + os.Args[0] + " file name")
// 		os.Exit(1)
// 	}

// 	f, err := ioutil.ReadFile(os.Args[1])
// 	if err != nil {
// 		fmt.Println("Cannot read the file")
// 		os.Exit(1)
// 	}
// 	// do something with the file
// 	fmt.Print(string(f))
// }
