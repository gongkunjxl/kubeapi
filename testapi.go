package main

import (
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
	fmt.Println("compete the main function")
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("https://master.example.com:8443", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//get all nodes
	/*nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(nodes.Items); i++ {
		fmt.Printf("%s \n", nodes.Items[i].ObjectMeta.Name)
	}*/

	//get all services
	services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(services.Items); i++ {
		fmt.Printf("%s \n", services.Items[i].ObjectMeta.Name)
	}

	//get all Namespaces
	/*namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(namespaces.Items); i++ {
		fmt.Printf("%s \n", namespaces.Items[i].ObjectMeta.Name)
	}*/

	//get all pods
	/*pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for i := 0; i < len(pods.Items); i++ {
		fmt.Printf("%s %s ", pods.Items[i].ObjectMeta.ClusterName, pods.Items[i].ObjectMeta.Namespace)
		fmt.Println(pods.Items[i].ObjectMeta.Name)
	}*/

	//create namesapce
	/*nc := new(apiv1.Namespace)
	ncTypeMeta := metav1.TypeMeta{Kind: "NameSpace", APIVersion: "v1"}
	nc.TypeMeta = ncTypeMeta
	nc.ObjectMeta = metav1.ObjectMeta{
		Name: "k8s-test",
	}
	nc.Spec = apiv1.NamespaceSpec{}
	clientset.Core().Namespaces().Create(nc)
	*/

}
