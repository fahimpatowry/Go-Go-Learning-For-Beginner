package main

import fmt "fmt"

// import "rsc.io/quote"
import "github.com/google/shlex"

func main() {
	// fmt.Println("Hello, World!")

	// form quote package
	// fmt.Println(quote.Go())
	// fmt.Println(quote.Glass())
	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Opt())

	z := "one \"two\" \"three\" four"
	numberS, err := shlex.Split(z)
	if err != nil {
        fmt.Println("Error:", err)
        return
    }

	fmt.Println(numberS)

	//for lop
	if len(numberS) > 0{
		for i := 0; i < len(numberS); i++ {
			fmt.Println(numberS[i])
		}
	}

}