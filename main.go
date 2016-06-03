package main

import (
    "time"
    "os"
    "fmt"
    "strconv"
)

func main() {
    args := os.Args[1:] // exclude program"s name

    if len(args) == 0 {
        fmt.Printf("Usage: %s COMMAND\n", os.Args[0])
        fmt.Printf("\n")
        fmt.Printf("Command can be only:\n")
        fmt.Printf("\tfutility [numHashes]\n")
            fmt.Printf("\n")
        return
    }

    cmd := args[0]
    t := time.Now()
    switch cmd {
        case "futility":
            numHashes := 16
            // This can print around 16 hashes or otherwise tends to get stuck
            if len(args) > 1 {
                n, err := strconv.Atoi(args[1])
                if err != nil {
                    fmt.Printf("Error parsing number of hashes: %v\n", err)
                    return
                }
                numHashes = n
            }
            hashless(numHashes)
        default:
            fmt.Printf("Wrong command: '%s'. Invoke without any parameters to see help.\n", cmd)
            return
    }
    fmt.Printf("Took %v\n", time.Since(t))
}
