package main

import "os"

func main() {
	f, _ := os.Create("./tmp1.txt")
	f.WriteString("you are very handsome\n")
	f.WriteString("you are ugly")
}
