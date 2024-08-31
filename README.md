# govectorize
This is a Go lib that helps converting a slice of strings into a slice of vectors.

Function `Generate([]string)` will collect all tokens from all the provided strings and return a slice of vectors (i.e. `map[string]float64`). Order of the vectors and the same as the order of input strings.

Example of usage:
````
import (
    "fmt"

    "github.com/example/dataprovider"
    "github.com/dmitriitimoshenko/govectorize"
)

func main() {
    strings := dataprovider.GetStringSlice()
    vectors := govectorize.Generate(strings)
    
    // vectors.Map() returns map[string]float64
    for _, vector := range vectors.Map() {
        fmt.Printf("\n\tkey\tvalue\n\n")
        for key, value := range vector {
            fmt.Printf("\t%s\t%f\n", key, value)
        }
        fmt.Printf("\n***********************************\n")
    }
}
````