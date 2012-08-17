# mpr

Tired of passing the same values to `fmt.Printf` over and over again?

    permutations := `
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    %v%v%v
    `
    fmt.Printf(permutations, 1, 2, 3, 1, 3, 2, 2, 1, 3, 2, 3, 1,
        3, 1, 2, 3, 2, 1)

Feels redundant and hard to read, yes?  How about this instead?

    permutations := `
    %(one)v%(two)v%(three)v
    %(one)v%(three)v%(two)v
    %(two)v%(one)v%(three)v
    %(two)v%(three)v%(one)v
    %(three)v%(one)v%(two)v
    %(three)v%(two)v%(one)v
    `

    nums := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
    }

    mpr.Printf(permutations, nums)


## TODO

* Use reflection to support structs in addition to maps

* Seriously consider making the syntax less Pythonic and more
  Go-esque (e.g., `{{name}}` instead of `%(name)s`)


## Inspiration

Inspired by Python's dictionary-passing syntax and a strong desire to
make things better.

--Steve Phillips / [@elimisteve](http://twitter.com/elimisteve)
