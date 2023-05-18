package node

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type NodeStats struct {
	Node              coreV1.Node
	PodList           []coreV1.Pod
	NodeWorkloadStats struct {
		Cpu    string
		Memory string
	}
}

func GetNodeStats() []NodeStats {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, _ := kubernetes.NewForConfig(config)

	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metaV1.ListOptions{})
	if err != nil {
		panic(err)
	}

	var nodeStatsList []NodeStats

	for _, n := range nodes.Items {
		pods, err := clientset.CoreV1().Pods("").List(context.Background(), metaV1.ListOptions{
			FieldSelector: "spec.nodeName=" + n.Name,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}

		ns := NodeStats{
			Node:    n,
			PodList: pods.Items,
		}
		nodeStatsList = append(nodeStatsList, ns)
	}

	//data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/nodes").DoRaw(context.Background())

	return nodeStatsList
}

func GetPods(nodeName string) {

}
