package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dotabuff/yasha"
	"github.com/dotabuff/yasha/dota"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected a .dem file as argument")
	}
	//added this to see what would happen if i tried to print args
	//i wonder if there's a default way to print any tuype, or if
	//when you define a type you have to define a way for it to print itself?
	fmt.Println(os.Args)

	//this for loop just loops over the args array discarding the first arg
	//that's what the _ does, this lets you parse multiple games
	for _, path := range os.Args[1:] {
		//initializes parser variable for ease of use := is initialize whereas = updates
		parser := yasha.ParserFromFile(path)
		//parser.OnSayText2 is defined in parser.go
		//i'm not really sure what's going on here, why are we redifining a function
		//ahh it's not defined in the struct but it's TYPE is i'm guessing?
		parser.OnSayText2 = func(tick int, obj *dota.CUserMsg_SayText2) {
			//ok so i'm guessing what's going on here is that when the parser runs any defined bits of it
			//are then passed the relevant information tick is probably gametime as an int?
			//star means pointer, so then it's passed a pointer to the dota.CUserMsg_SayText2 value?
			//for the given tick? dota.xxxx is the prefix for protobuf generated crap i think?o
			fmt.Printf("%s - %07d | %s: %s\n", filepath.Base(path), tick, obj.GetPrefix(), obj.GetText())
			fmt.Println(obj)
		}
		parser.OnLocationPing = func(tick int, obj *dota.CDOTAUserMsg_LocationPing) {
			fmt.Println(obj)
		}

		parser.Parse()
	}
}
