# mapr: map[] printer for Go/golang

"Map printer?  Huh?"

See the example below.

## The Example Below

Tired of passing the same values to `fmt.Printf` over and over again?

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


Feel redundant yet?  Were the variables passed in in the correct
order?  Hard to say...  Try this instead:

    permutations := `
    %(one)v%(two)v%(three)v
    %(one)v%(three)v%(two)v
    %(two)v%(one)v%(three)v
    %(two)v%(three)v%(one)v
    %(three)v%(one)v%(two)v
    %(three)v%(two)v%(one)v
    `
    nums := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
    }
    mapr.Printf(permutations, nums)


## TODO

* Make syntax less Pythonic and more Go-esque (e.g., `{{name}}`
  instead of `%(name)s`)

* Use reflection to support structs in addition to maps

* Make the example more compelling by making it more similar to the
  example that actually inspired `mapr`

* Add unit tests


## Inspiration

Inspired by Python's dictionary formatting syntax, certain experiences
with Go + much printing values contained in maps, and a strong desire
to make things better.

--Steve Phillips / [@elimisteve](http://twitter.com/elimisteve)
