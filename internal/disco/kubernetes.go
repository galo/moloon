package disco

import (
	"github.com/galo/moloon/internal/logging"
	"github.com/galo/moloon/pkg/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// KubernetesDiscovery the K8s discovery service
type KubernetesDiscovery struct {
	k8sServer string
	ns        string
}

// NewKubernetesDiscoveryService Creates a new Kubernetes discovery service, The url points ot the K8s cluster,
// use nil to use in cluster discovery.
func NewKubernetesDiscoveryService(url string, ns string) *KubernetesDiscovery {
	return &KubernetesDiscovery{k8sServer: url, ns: ns}
}

// GetAll Get all agents running on a cluster
func (k *KubernetesDiscovery) GetAll() ([]*models.Agent, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var agents []*models.Agent

	// Look for agents in the agents Namespace
	pods, err := clientSet.CoreV1().Pods("agents").List(metav1.ListOptions{})
	if err != nil {
		logging.Logger.Errorf("There are %d pods in the cluster", len(pods.Items))
		return nil, err
	}

	logging.Logger.Infof("There are %d pods in the cluster", len(pods.Items))

	//

	// TODO: create the agents slice

	return agents, nil
}
