package apl

import (
	k "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func GetKubeClient() (*k.Clientset, error) {

	configFile := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return nil, err
	}

	settings, err := clientcmd.LoadFromFile(configFile)
	if err != nil {
		return nil, err
	}

	config, err := clientcmd.NewDefaultClientConfig(*settings, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}

	// creates the clientset
	clientset, err := k.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func GetKubeClientREST() {

}
