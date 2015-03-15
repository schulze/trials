// A basic sanity check of Go's rand.Prime() and rsa.GenerateKey() functions.
// The program generates lots of random primes and checks for duplicates.
// Because generating big primes takes a long time the program, by default,
// only uses keys of 256 bits.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"flag"
	"log"
	"math/big"
	"runtime"
	
	"time"
)

var bits = flag.Int("bits", 256, "size of RSA key in bits")

type primesTable map[uint64] *big.Int

func check(table primesTable, p *big.Int) error {
	if q, ok := table[p.Uint64()]; ok {
 		if p.Cmp(q) == 0 {
 			log.Printf("same primes TODO \n")
 			return errors.New("duplicate prime")
 		}
 		log.Printf("same uint64, but not same prime TODO \n")
	}
	table[p.Uint64()] = p
	return nil	
}

func generate(keys chan *rsa.PrivateKey) {
	for {
		priv, err := rsa.GenerateKey(rand.Reader, *bits)
		if err != nil {
			log.Printf("error generating key: %s\n", err)
			continue
		}
		if err = priv.Validate(); err != nil {
			log.Printf("invalid key: %s\n", err)
			continue
		}
		keys <- priv
	}
}

func main() {
	flag.Parse()

	table := make(map[uint64] *big.Int)	
	keys := make(chan *rsa.PrivateKey)
	
	go func() {
		d := time.Tick(time.Minute)
		for {
			<- d
			log.Printf("checked %v keypairs\n", len(table))
		}
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go generate(keys)
	}
	
	for {
		priv := <- keys
		check(table, priv.Primes[0])
		check(table, priv.Primes[1])
	}
}