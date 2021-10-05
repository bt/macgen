package main

import (
    "fmt"
    "flag"
)

func main() {
    prefix := flag.String("prefix", "", "prefix for mac address")
    quantity := flag.Int("quantity", 1, "number of mac addresses to generate")
    
    flag.Parse()
    
    toGenerate := 6
    if *prefix != "" {
        toGenerate = 6 - len(*prefix) / 2
        fmt.Printf("%s", *prefix)
    }
    
    for i := 1; i <= toGenerate; i++ {
        fmt.Printf("%s", "AA")
    }
}