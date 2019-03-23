 easygen config easygen | gofmt > ../../config.go
 easygen flags easygen | sed 's/^package easygen/package main/' | gofmt > flags.go
