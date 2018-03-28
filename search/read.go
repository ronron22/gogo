01: package main
02:
03: import (
04: "fmt"
05: "os"
06: "io"
07: )
08:
09: func main() {
10: fic, err := os.Open("monFichier.txt")
11: defer fic.Close()
12:
13: if err != nil {
14: fmt.Println("Ouverture de fichier impossible")
15: fmt.Println(err)
16: os.Exit(1)
17: }
18:
19: data := make([]byte, 10)
20: for {
21: n, err := fic.Read(data)
22: if err == io.EOF {
23: fmt.Println("Fin du fichier")
24: break
25: }
26: for i := 0; i < n; i++ {
27: fmt.Print(string(data[i]))
28: }
29: }
30: }
Nous 
