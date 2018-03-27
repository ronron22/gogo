package  main

import (
	"os"
	//"bufio"
	"io/ioutil"
	"fmt"
)

var  inputfile  string = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	myfile, err := ioutil.ReadFile( inputfile)
	check(err)
	fmt.Print(string(myfile))


	f, err := os.Open(inputfile)
	check(err)

	ffs, err := f.Seek(6, 0)
	check(err)
	bb := make([]byte, 2)
	n2, err := f.Read(bb)
	check(err)
	fmt.Printf(string(n2), offs)


}
	
	

