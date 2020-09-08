package server

import "fmt"

const version = "v0.1.0"

// Print showcases submodule selection.
func Print() {
	fmt.Printf("ocis-webdav version: %v", version)
}