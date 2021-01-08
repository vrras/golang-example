package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Firas Luthfi Dwiyansyah", "Firas"))
	fmt.Println(strings.Contains("Firas Luthfi Dwiyansyah", "Putri"))
	fmt.Println(strings.Split("Firas Luthfi Dwiyansyah", " "))
	fmt.Println(strings.ToLower("Firas Luthfi Dwiyansyah"))
	fmt.Println(strings.ToUpper("Firas Luthfi Dwiyansyah"))
	fmt.Println(strings.ToTitle("Firas Luthfi Dwiyansyah"))
	fmt.Println(strings.Trim("    Firas Luthfi Dwiyansyah  ", " "))
	fmt.Println(strings.ReplaceAll("Firas Firas Firas Firas", "Firas", "Luthfi"))
}
