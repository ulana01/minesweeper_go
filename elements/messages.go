package main

import "github.com/alem-platform/ap"

func println(s string) {
	for _, r := range s {
		ap.PutRune(r)
	}
	ap.PutRune('\n')
}

func print(s string) {
	for _, r := range s {
		ap.PutRune(r)
	}
}
