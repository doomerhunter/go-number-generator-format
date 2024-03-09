package main

import (
    "flag"
    "fmt"
    "math/rand"
//    "os"
    "time"
)

func main() {
    var from, to, minInt, maxInt, minFrac, maxFrac, n int
    var floatFlag bool
    var mode string

    flag.IntVar(&from, "from", 0, "Start of range. Default = 0")
    flag.IntVar(&to, "to", 0, "End of range. Default = 0")
    flag.IntVar(&minInt, "minint", 0, "Minimum integer digits. Default = 0")
    flag.IntVar(&maxInt, "maxint", 0, "Maximum integer digits. Default = 0")
    flag.IntVar(&minFrac, "minfrac", 0, "Minimum fraction digits. Default = 0")
    flag.IntVar(&maxFrac, "maxfrac", 0, "Maximum fraction digits. Default = 0")
    flag.BoolVar(&floatFlag, "float", false, "Include floating point numbers. Default = false")
    flag.StringVar(&mode, "mode", "sequential", "Generation mode (sequential or random). Default = sequential")
    flag.IntVar(&n, "n", 0, "Number of random numbers to generate. Default = 0")

    flag.Parse()

    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

//    fmt.Println("Matching numbers:")
    if mode == "random" {
        generateRandomNumbers(from, to, n, minInt, maxInt, minFrac, maxFrac, floatFlag)
    } else {
        generateSequentialNumbers(from, to, minInt, maxInt, minFrac, maxFrac, floatFlag)
    }
}

func generateRandomNumbers(from, to, n, minInt, maxInt, minFrac, maxFrac int, floatFlag bool) {
    for i := 0; i < n; i++ {
        if floatFlag {
            // Generate random floating-point numbers
	    //number := rand.Float64()*(float64(to)-float64(from)) + float64(from)
            fracFactor := rand.Intn(maxFrac-minFrac+1) + minFrac
            fmt.Printf("%0*d.%0*d\n", minInt, rand.Intn(to-from+1)+from, fracFactor, rand.Intn(int(pow10(fracFactor))))
        } else {
            // Generate random integers with leading zeros
            number := rand.Intn(to-from+1) + from
            fmt.Printf("%0*d\n", minInt, number)
        }
    }
}

func generateSequentialNumbers(from, to, minInt, maxInt, minFrac, maxFrac int, floatFlag bool) {
    for i := from; i <= to; i++ {
        if floatFlag {
            for frac := minFrac; frac <= maxFrac; frac++ {
                for j := 0; j < int(pow10(frac)); j++ {
                    fmt.Printf("%0*d.%0*d\n", minInt, i, frac, j)
                }
            }
        } else {
            // Print integers with leading zeros
            fmt.Printf("%0*d\n", minInt, i)
        }
    }
}

func pow10(n int) float64 {
    return float64(int(pow10Int(n)))
}

func pow10Int(n int) int {
    result := 1
    for i := 0; i < n; i++ {
        result *= 10
    }
    return result
}

