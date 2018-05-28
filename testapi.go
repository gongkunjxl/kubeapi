package main

import (
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
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
	/*services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(services.Items); i++ {
		fmt.Printf("%s \n", services.Items[i].ObjectMeta.Name)
	}*/

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
	ncresult, err := clientset.Core().Namespaces().Create(nc)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Create Namespace: %s success\n", ncresult.ObjectMeta.Name)*/
	//create new pod
	/*var privileged bool = false
	security := apiv1.SecurityContext{Privileged: &privileged}
	pod := new(apiv1.Pod)
	pod.TypeMeta = metav1.TypeMeta{Kind: "Pod", APIVersion: "v1"}
	objMeta := metav1.ObjectMeta{Name: "hello-pod", Namespace: "k8s-test", Labels: map[string]string{"name": "hello-pod"}}
	pod.ObjectMeta = objMeta
	spec := apiv1.PodSpec{
		NodeName: "node1.example.com",
		Containers: []apiv1.Container{
			apiv1.Container{
				Name:  "hello-pod",
				Image: "openshift/hello-openshift",
				Ports: []apiv1.ContainerPort{
					apiv1.ContainerPort{
						Name:          "tcp1",
						ContainerPort: 8080,
						Protocol:      apiv1.ProtocolTCP,
					},
				},
				Resources: apiv1.ResourceRequirements{
					Requests: map[apiv1.ResourceName]resource.Quantity{apiv1.ResourceCPU: resource.MustParse("500m"), apiv1.ResourceMemory: resource.MustParse("1Gi")},
				},
				VolumeMounts: []apiv1.VolumeMount{
					apiv1.VolumeMount{
						Name:      "tmp",
						MountPath: "/tmp",
					},
				},
				TerminationMessagePath: "/dev/termination-log",
				ImagePullPolicy:        "IfNotPresent",
				SecurityContext:        &security,
			},
		},
		Volumes: []apiv1.Volume{
			apiv1.Volume{
				Name: "tmp",
			},
		},
		RestartPolicy:      "Always",
		ServiceAccountName: "",
	}
	pod.Spec = spec
	podresult, err := clientset.Core().Pods("k8s-test").Create(pod)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Create a new Pod: %s\n", podresult.ObjectMeta.Name)*/
	//create service
	service := new(apiv1.Service)
	service.TypeMeta = metav1.TypeMeta{Kind: "Service", APIVersion: "v1"}
	service.ObjectMeta = metav1.ObjectMeta{Name: "hello-service", Namespace: "k8s-test", Labels: map[string]string{"name": "hello-service"}}
	svcSpec := apiv1.ServiceSpec{
		Ports: []apiv1.ServicePort{
			apiv1.ServicePort{
				Name: "app1",
				Port: 22,
				TargetPort: 22,
				Protocol: apiv1.ProtocolTCP
			},
			apiv1.ServicePort{
				Name: "app2",
				Port: 9000,
				TargetPort: 9000,
				Protocol: apiv1.ProtocolTCP
			}
		},
		Selector: map[string]string{"name": "hello-pod"},
		Type: apiv1.ServiceTypeClusterIP,
		SessionAffinity: apiv1.ServiceAffinityNone
	}
	service.Spec = svcSpec
	svcresult, err := clientset.Servicess("k8s-test").Create(service)
	if err != nil{
		panic(err.Error())
	}else{
		fmt.Printf("Create service %s \n",svcresult.ObjectMeta.Name)
	}
}

/*





















 */
