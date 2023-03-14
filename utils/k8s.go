package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/yangkaiyue/gin-exp/global"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8SCli struct {
	Client *kubernetes.Clientset
}

func NewK8SCli() (*K8SCli, error) {

	var (
		conf *rest.Config
		err  error
	)

	if viper.GetString("k8s.mode") != "out" {
		conf, err = rest.InClusterConfig()
	} else {
		conf, err = clientcmd.BuildConfigFromFlags("", fmt.Sprintf("%v/%v", global.ProjectPath, "conf/kube.config"))
	}
	if err != nil {
		return nil, err
	}

	cli, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	return &K8SCli{
		Client: cli,
	}, nil
}
