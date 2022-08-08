package security

import (
	"log"
)

type FloodChain struct {
	NextChain Security
}

func (f *FloodChain) Next(r Requisition) {
	log.Println("FloodChain")
	f.NextChain.Next(r)
}