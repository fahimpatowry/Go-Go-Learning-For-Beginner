package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "this needs to go in a sfsf"

	file, err := os.Create("./mycogofile.txt")

	checkNillErr(err)

	length, err1 := io.WriteString(file, content)

	checkNillErr(err1)
	fmt.Println("length is: ", length)

	defer file.Close()

	readFile("./mycogofile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	checkNillErr(err)

	fmt.Println("length is: ", string(dataByte))

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
