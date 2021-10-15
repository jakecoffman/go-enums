# go-enums

Type-safe enums in Go

Well, as close as we can get with the language currently.

I fully support adding proper Enums and pattern matching to the language. In the meantime though I'll do this.

## problem

The issue with using `iota` as enums is you can accidentally assign other values to the type and Go is ok with it.

```go
type Color int

const (
	Red = Color(iota+1)
	Blue
)
```

Even though we've defined the colors here as a type `Color` Go won't stop us from doing this:

```go
var color Color = Red
color = 7 // not a valid color, compiles ok
```

## solution

The solution here is to use an interface to hide the types behind a layer of abstraction...

Well just go look at [example.go](example.go) and you'll see!
