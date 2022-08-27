package security

import (
	"log"
)

type FloodChain struct {
	NextChain Security
}

func (f *FloodChain) Next(r Request) {
	log.Println("FloodChain")
	f.NextChain.Next(r)
}