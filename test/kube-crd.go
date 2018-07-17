package main

import (
	"flag"
	restClient "k8s.io/client-go/rest"
	"github.com/k8s.io/client-go/tools/clientcmd"
	"github.com/xieydd/kubenetes-crd/client"
	apiextcs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"github.com/xieydd/kubenetes-crd/pkg/apis/v1alpha"
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
	"fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/cache"
)
var (
	kubeconf   = flag.String("kubeconfig","admin.conf","Path to a kube config. Only required if out-of-cluster.")
	kubeMaster = flag.String("master","","The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()
	
	// Get ~/.kube/config
	config, err := GetClientConfig(*kubeMaster,*kubeconf)
	Must(err)
	
	// create clientset and create our QueueJob, this only need to run once
	clientset, err := apiextcs.NewForConfig(config)
	Must(err)
	
	// note: if the QueueJob CRD exist our CreateQueueJob function is set to exit without an error
	err = v1alpha.CreateQueueJob(clientset)
	Must(err)
	
	// Wait for the CRD to be created before we use it (only needed if its a new one)
	time.Sleep(3 * time.Second)
	
	queuejobcs, scheme, err := v1alpha.NewClient(config)
	Must(err)
	
	// Create a QueueJob CRD client interface
	crdclient := client.CrdClient(queuejobcs,scheme,"default")
	
	
	nodeselect := make(map[string]string)
	nodeselect["nodeSelector"] = "node1"
	
	matchlabel := make(map[string]string)
	matchlabel["xieydd"] = "test"
	
	// Create a new Object and write to k8s
	example := &v1alpha.QueueJob{
		TypeMeta: metav1.TypeMeta{
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:	"xieydd",
			Labels:  matchlabel,
		},
		Spec: v1alpha.QueueJobSpec{
			SchedSpec: v1alpha.SchedulingSpecTemplate{
				nodeselect,
				2,
			},
			TaskSpecs: []v1alpha.TaskSpec{
				v1alpha.TaskSpec{
					Replicas: 1,
					Template: v1.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name: "busybox",
							Labels: matchlabel,
						},
						Spec: v1.PodSpec{
							SchedulerName:   "kar-scheduler",
							RestartPolicy: v1.RestartPolicyNever,
							Containers: []v1.Container {
									v1.Container{
										Image: "busybox",
										Name:  "busybox",
										ImagePullPolicy: v1.PullIfNotPresent,
									},
							},
						},
					},
				},
			},
		},
		Status: v1alpha.QueueJobStatus{
			Running: 1,
		},
	}
	
	
	result, err := crdclient.Create(example)
	if err == nil {
		fmt.Printf("CREATED: %#v\n",result)
	}else if (apierrors.IsAlreadyExists(err)) {
		fmt.Printf("ALREADY EXEISTS:: %#v\n",result)
	}else {
		Must(err)
	}
	
	items,err := crdclient.List(metav1.ListOptions{})
	Must(err)
	fmt.Printf("List: \n%s\n",items)
	
	// Example Controller
	// Watch for changes in Example objects and fire Add, Delete, Update callbacks
	_, controller := cache.NewInformer(
		crdclient.NewListWatch(),
		&v1alpha.QueueJob{},
		time.Minute*10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("add: %s \n", obj)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Printf("delete: %s \n", obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Printf("Update old: %s \n      New: %s\n", oldObj, newObj)
			},
		},
	)
	
	stop := make(chan struct{})
	go controller.Run(stop)
	
	// Wait forever
	select {}
	
	
}

func Must(err error) {
	if err != nil {
		panic(err.Error())
	}
}


func GetClientConfig(master string,kubeconfig string) (*restClient.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags(master,kubeconfig)
	}
	return restClient.InClusterConfig()
}


