package security

import (
)

type FloodChain struct {
	NextChain Security
}

//Check if the user is flooding commands
func (f *FloodChain) Next(r Request) {
	//@TODO
	f.NextChain.Next(r)
}