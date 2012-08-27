// Steve Phillips / elimisteve
// 2012.08.17

package main

import (
	"fmt"
	"github.com/elimisteve/mapr"
)

func main() {
	// The Old Way
    permutations := `
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
`
    one, two, three := 1, 2, 3

    fmt.Printf(permutations, one, two, three, one, three, two, two,
        one, three, two, three, one, three, one, two, three, two, one)


	// The New Way
    permutations = `
    {{one}}{{two}}{{three}}
    {{one}}{{three}}{{two}}
    {{two}}{{one}}{{three}}
    {{two}}{{three}}{{one}}
    {{three}}{{one}}{{two}}
    {{three}}{{two}}{{one}}
`
    nums := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
    }
    mapr.Printf(permutations, nums)
}
