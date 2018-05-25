package main

import (
    "fmt"
    "flag"
    apiv1 "k8s.io/client-go/pkg/api/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
)
var (
    kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
    fmt.Println("come main")
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
    //create namesapce
    nc := new(apiv1.Namespace)
    ncTypeMeta := metav1.TypeMeta{Kind: "NameSpace", APIVersion: "v1"}
    nc.TypeMeta = ncTypeMeta

    nc.ObjectMeta = metav1.ObjectMeta{
        Name: "k8s-test",
    }
    nc.Spec = apiv1.NamespaceSpec{}
    clientset.Core().Namespaces().Create(nc)
    
}
