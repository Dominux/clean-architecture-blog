package main

import "fmt"


type Kek struct {
	Ahahh string
}

type Lol struct {
	Kek *Kek
	Cheburek string
}

func main() {
	lol := Lol{ Kek: nil }
	fmt.Printf("%v", lol.Kek)
}
