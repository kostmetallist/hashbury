package main

func Digest(key uint32, divisor uint32) uint32 {
	return key % divisor
}
