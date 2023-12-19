package main
    
import (
	"fmt"
	"math"
)
    
func main() {
	
    fmt.Println("Hello, Calc!")
	
	
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // io.EOF
			break
		}
		println(line)
	}
}