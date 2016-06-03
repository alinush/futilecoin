package main

import (
    "fmt"
    "crypto/sha256"
    "crypto/rand"
    "math/big"
    "encoding/hex"
)

// generates r -> h(r) -> h(h(r)) -> ... -> h^max(r) such that
// if i < j, then h^i(r) > h^j(r)
// (nonces are used to adjust the hash values)
func hashless(max int) {
    // Generate some random bytes
    bytes := make([]byte, 32)
    rand.Read(bytes)

    // Print said random bytes
    fmt.Printf("Generating %d hashes, each one less than its predecessor.\n", max)
    prevHash := sha256.Sum256(bytes)
    fmt.Printf("Initial random data (%d bytes): %s\nHash of data is: %s\n",
        len(bytes), hex.EncodeToString(bytes[:]),
        hex.EncodeToString(prevHash[:]))

    found := 0
    for found < max {
        // Store the previous hash as a number
        prevNum := big.NewInt(0)
        prevNum.SetBytes(prevHash[:])

        // Create a nonce, initialize it to 0
        nonce := big.NewInt(0)
        var iter uint64 = 1 // Used as a progress indicator only
        for {
            // Adjust the nonce for the next iteration
            nonce.Add(nonce, big.NewInt(1))

            // data := prevHash | nonce
            data := append(prevHash[:], nonce.Bytes()...)
            // Compute SHA256(data)
            hash := sha256.Sum256(data)

            // Store the current hash as a number
            curNum := big.NewInt(0)
            curNum.SetBytes(hash[:])

            // Indicate some progress in the UI
            if iter % 1000000000 == 0 {
                fmt.Printf("+")
                //fmt.Printf("Attempt #%d w/ nonce %v...\n", iter, nonce)
                //fmt.Printf("Hash (%d bytes): %s\n",
                //    len(hash),
                //    hex.EncodeToString(hash[:]))
            }
            iter++

            // See if the current hash is less than the previous hash
            less := curNum.Cmp(prevNum) == -1

            if less {
                fmt.Println()
                fmt.Printf("Found SHA256(%v|%v) hash:\n%v < %v\n", 
                    hex.EncodeToString(prevHash[:]),
                    hex.EncodeToString(nonce.Bytes()),
                    hex.EncodeToString(hash[:]),
                    hex.EncodeToString(prevHash[:]))
                prevHash = hash
                break
            }
        }

        found++
    }
}
