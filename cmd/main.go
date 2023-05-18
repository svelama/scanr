package main

import (
	"github.com/svelama/scanr/internal/k8s"
)

func main() {
	s := k8s.NewService()

	s.GetNodes()
}
