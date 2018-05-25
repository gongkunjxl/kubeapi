package main

import (
    "fmt"
    "time"

    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/rest"
)
var (
    kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

func main() {
    nc := new(v1.Namespace)
    ncTypeMeta := unversioned.TypeMeta{Kind: "NameSpace", APIVersion: "v1"}
    nc.TypeMeta = ncTypeMeta

    nc.ObjectMeta = v1.ObjectMeta{
        Name: "k8s-test",
    }

    nc.Spec = v1.NamespaceSpec{}
}
