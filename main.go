package main

import (
	"fmt"

	"github.com/FcoManueel/Dinosaur/dino"
)

func main() {
	fmt.Println("Hello dinosaur! Enjoy your evolution. ")

	d := dino.New(100)

	d.Run(100)
	//  // I think that this way will make easier the communication with the front end
	//    i := 0
	//    for {
	//        state := d.Step()
	//        fmt.Printf(`\n----------------------------\n
	//                     Step %d: %+v\n`, i, state)
	//        SendToBrowser(state) // update clients
	//        i++
	//    }
}
