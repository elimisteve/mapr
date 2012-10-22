# mapr: map[] printer for Go/golang

"Map printer?  Huh?"

See the example below.

## The Example Below

Tired of passing the same values to `fmt.Printf` over and over again?

    // WTF are all these %v's? Can't see the structure of the data

    permutations := `
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    `
    one, two, three := 1, 2, 3

    // Long function call, passing in the same variables several times

    fmt.Printf(permutations, one, two, three, one, three, two, two,
        one, three, two, three, one, three, one, two, three, two, one)


Hard to read? Feel redundant yet?  Were the variables passed in in the
correct order?  Hard to say...  Try this instead:

    // Easier to read

    permutations := `
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

    // One-line, easy-to-read function call. Much better!

    mapr.Printf(permutations, nums)


## TODO

* Use reflection to support structs in addition to maps

* Find way to create combine `mapr` with Go's HTML templates so Go web
  developers can benefit from this, too

* Make the example more compelling by making it more similar to the
  example that actually inspired `mapr`, with many lines and `%s`'s
  and `%v`'s strewn about

* Add unit tests


## Done

* Make syntax less Pythonic and more Go-esque (e.g., `{{name}}`
  instead of `%(name)s`)

* Add mapr.Sprintf


## Inspiration

Inspired by Python's dictionary formatting syntax, certain experiences
with Go + much printing values contained in maps, and a strong desire
to make things better.

--Steve Phillips / [@elimisteve](http://twitter.com/elimisteve)
