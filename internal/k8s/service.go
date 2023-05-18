package k8s

import (
	"fmt"

	"github.com/svelama/scanr/internal/k8s/node"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetNodes() {
	ns := node.GetNodeStats()

	for _, n := range ns {
		fmt.Println("name: " + n.Node.Name)
		fmt.Printf("allocatable cpu: %vm \n", n.Node.Status.Allocatable.Cpu().MilliValue())
		fmt.Printf("allocatable memory: %v\n", n.Node.Status.Allocatable.Memory())
		fmt.Printf("allocatable pods: %v\n", n.Node.Status.Allocatable.Pods())
		fmt.Printf("node info: %v\n", n.Node.Status.NodeInfo)
		fmt.Println()
		// fmt.Printf("list of images on the node: %v\n", n.Status.Images)

		fmt.Println("pods: ")
		for _, p := range n.PodList {
			fmt.Println(p.Name)
		}
	}
}
