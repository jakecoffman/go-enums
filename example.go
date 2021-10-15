package main

import (
	"encoding/json"
	"fmt"
	"github.com/jakecoffman/go-enums/color"
	"strings"
)

func main() {
	var theColor color.Color = color.Red

	// Printing just works since we implemented the const as strings.
	// We didn't have to define Stringer on each color.
	fmt.Println("Color is:", theColor)

	// We don't have pattern matching in Go, so to make is so all enums must be handled,
	// add a func to the interface to do what you need. It will force all Color types
	// to implement the func and thus achieve what we want with pattern matching.
	theColor.DoThing()

	// Now we get to the hard part: how can we use this "enum" in a struct when decoding
	// JSON? This is a pretty common problem. One solution is to make a type wrapper
	// around the enum and give it a custom Marshaler/Unmarshaler. See color.Wrapper.
	const data = `{"color":"blue"}`
	var thing myThing
	if err := json.NewDecoder(strings.NewReader(data)).Decode(&thing); err != nil {
		panic(err)
	}

	// This looks good because fmt.Stringer is implemented on color.Wrapper.
	fmt.Println(thing)
	// Since Color here is an interface it could be nil, however the json package's Decode would
	// have errored so it can't be nil. It is unfortunate the compiler can't guarantee that, but
	// this is the best we can get!
	if thing.Color.Color != color.Blue {
		panic(thing)
	}
	thing.Color.DoThing()
}

type myThing struct {
	Color color.Wrapper `json:"color"`
}
