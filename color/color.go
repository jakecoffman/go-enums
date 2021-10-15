package color

import (
	"encoding/json"
	"fmt"
)

// Color represents a color "enum" type. Use the consts like Red and Blue
// when the value is known.
type Color interface {
	// DoThing is an example of how to avoid another switch statement. Adding
	// a new method here forces all of the color types below to have to implement
	// it as well. This safely achieves what pattern matching can do.
	DoThing()
}

// If red and blue were struct{} we wouldn't be able to make this const.
// The downside of making them a string (or int, etc) is you could still
// do `red := red("green")` but we're not exporting the type to help avoid this.
const (
	Red  = red("red")
	Blue = blue("blue")
)

// Error at compile time if these types don't implement the interface
var (
	_ = Color(Red)
	_ = Color(Blue)
)

// FromString takes a string value and returns the Color type. If the string value is
// not a valid Color then an error is returned.
func FromString(str string) (Color, error) {
	// A switch statement is unavoidable in this case, so unfortunately this
	// has to be kept up-to-date with the const above, but I've positioned it
	// here as close to the const as possible to help with that. There should
	// be no reason to ever switch on this value outside of this function.
	switch str {
	case string(Blue):
		return Blue, nil
	case string(Red):
		return Red, nil
	default:
		return nil, fmt.Errorf("invalid color " + str)
	}
}

type red string

func (red) DoThing() {
	fmt.Println("red DoThing")
}

type blue string

func (blue) DoThing() {
	fmt.Println("blue DoThing")
}

// Wrapper provides a way to add the Color to another struct and have it
// marshal/unmarshal as expected. If the string in the JSON is not a valid
// color, the json package's Decode or json.Unmarshal will return an error.
type Wrapper struct {
	// Color will hold the underlying "enum"
	Color
}

// MarshalJSON satisfies json.Marshaler
func (w *Wrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.Color)
}

// UnmarshalJSON satisfies json.Unmarshaler
func (w *Wrapper) UnmarshalJSON(data []byte) (err error) {
	var c string
	if err = json.Unmarshal(data, &c); err != nil {
		// this error would represent that the data sent was not a string
		return
	}
	w.Color, err = FromString(c)
	// this error would mean that the string wasn't a valid Color
	return
}

// String implements fmt.Stringer
func (w Wrapper) String() string {
	return fmt.Sprint(w.Color)
}
