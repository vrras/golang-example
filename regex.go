package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("Fir([a-z])s")

	fmt.Println(regex.MatchString("Firas"))
	fmt.Println(regex.MatchString("Firus"))
	fmt.Println(regex.MatchString("FirOs"))

	fmt.Println(regex.FindAllString("Firas Firus Firis furus fetes foros", -1))
}
