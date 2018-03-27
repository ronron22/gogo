package  main

import (
	"os"
	//"bufio"
	//"io/ioutil"
	"fmt"
)

var  inputfile  string = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/*myfile, err := ioutil.ReadFile( inputfile)
	check(err)
	fmt.Print(string(myfile))*/


	f, err := os.Open(inputfile)
	check(err)

	defer f.Close() 

	//o2, err := f.Seek(6, 0)
    //check(err)
    buffer := make([]byte, 50)
    n2, err := f.Read(buffer)
    check(err)
    fmt.Printf("%d bytes @  %s\n", n2, string(buffer))

	newoffset, err := f.Seek(50, 0)
	check(err)
	
	n2, err = f.Read(buffer)

	 fmt.Println("glagla", string(buffer[:n2]))

         fmt.Println("New off set : ", newoffset)

	

}
	
	

