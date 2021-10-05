package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var src = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandHex(n int) string {
    b := make([]byte, n+1)

    if _, err := src.Read(b); err != nil {
            panic(err)
    }

    return hex.EncodeToString(b)[:n*2]
}

func main() {
    prefix := flag.String("prefix", "", "prefix for mac address")
    quantity := flag.Int("quantity", 1, "number of mac addresses to generate")
    
    flag.Parse()
    
    toGenerate := 6

    for q := 0; q < *quantity; q++ {
        mac := ""

        if *prefix != "" {
            if len(*prefix) % 2 != 0 {
                panic("invalid prefix, length must be multiple of 2")
            }

            toGenerate = 6 - len(*prefix) / 2
            mac += *prefix
        }

        mac += RandHex(toGenerate)
        fmt.Println(strings.ToLower(mac))
    }
}