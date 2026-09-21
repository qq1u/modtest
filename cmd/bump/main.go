package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/mod/semver"
)

func main() {
	raw, err := os.ReadFile("VERSION.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := strings.TrimSpace(string(raw))
	cur := semver.Canonical(s)
	if !semver.IsValid(cur) {
		log.Fatalf("invalid version %q", s)
	}

	var major, minor, patch uint64
	if _, err := fmt.Sscanf(cur, "v%d.%d.%d", &major, &minor, &patch); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("v%d.%d.%d\n", major, minor, patch+1)
}
