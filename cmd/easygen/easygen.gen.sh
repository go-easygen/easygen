 easygen -tf config easygen | gofmt > ../../config.go
 easygen -tf flags easygen | sed 's/^package easygen/package main/' | gofmt > flags.go
