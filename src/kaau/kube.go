package kaau

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func kube() {

	var ns, label, field, maxClaims string
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	flag.StringVar(&ns, "namespace", "", "namespace")
	flag.StringVar(&label, "l", "", "Label selector")
	flag.StringVar(&field, "f", "", "Field selector")
	flag.StringVar(&maxClaims, "max-claims", "100Gi", "Maximum total claims to watch")
	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()

	// total resource quantities
	//	var totalClaimedQuant resource.Quantity
	//	maxClaimedQuant := resource.MustParse(maxClaims)

	// bootstrap config
	fmt.Println()
	fmt.Println("Using kubeconfig: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	fmt.Println(reflect.TypeOf(clientset))
	if err != nil {
		log.Fatal(err)
	}
	api := clientset.CoreV1()
	fmt.Println(reflect.TypeOf(api))
	// initial list
	listOptions := metav1.ListOptions{LabelSelector: label, FieldSelector: field}

	namespace, err := api.Namespaces().List(listOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Name\t\t Status\n")
	for _, v := range namespace.Items {
		fmt.Printf("%s\t\t  %s\n", v.Name, v.Status)
		//		getpods(clientset, v.Name)
		//		getsvc(clientset, v.Name)
	}
}
