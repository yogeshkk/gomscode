package main

import (
	"fmt"
	"gomscode/src/services"
	"net/http"
)

func main() {
	fmt.Println("hello main")
	//	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	//	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	//	flag.Parse()
	//	kaau.GetKubeClient(kubeconfig)
	r := services.NewRouter()
	if err := http.ListenAndServe(":8000", r); err != nil {
		panic(err)
	}
}
