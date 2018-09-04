package client

import (
	"github.com/xieydd/kubenetes-crd/pkg/apis/v1alpha"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

//This file implement all the (CRUD) client methods we need to access our CRD object
/*type QueueJobGetter interface {
	QueueJobs(namespaces string) QueueJobInterface
}*/

/*type QueueJobInterface interface {
	Create(list crd.QueueJob) (crd.QueueJob, error)
	Update(crd.QueueJob) (crd.QueueJob, error)
	UpdateStatus(crd.QueueJob) (crd.QueueJob, error)
	Delete(name string, options metav1.DeleteOptions) error
	Get(name string) (crd.QueueJob, error)
	List(opts metav1.ListOptions) (crd.QueueJobList, error)
}*/

func CrdClient(cl *rest.RESTClient, scheme *runtime.Scheme, namespace string) *queuejobs {
	return &queuejobs{
		client: cl,
		ns:     namespace,
		plural: v1alpha.QueueJobPlural,
		codec:  runtime.NewParameterCodec(scheme),
	}
}

// queuejobs implements QueueJobInterface
type queuejobs struct {
	client *rest.RESTClient
	ns     string
	plural string
	codec  runtime.ParameterCodec
}

func (q *queuejobs) Create(obj *v1alpha.QueueJob) (*v1alpha.QueueJob, error) {
	var result v1alpha.QueueJob
	err := q.client.Post().Namespace(q.ns).Resource(q.plural).Body(obj).Do().Into(&result)
	return &result, err
}

func (q *queuejobs) Update(obj *v1alpha.QueueJob) (*v1alpha.QueueJob, error) {
	var result v1alpha.QueueJob
	err := q.client.Put().Namespace(q.ns).Resource(q.plural).Body(obj).Do().Into(&result)
	return &result, err
}

func (q *queuejobs) Delete(name string, options *metav1.DeleteOptions) error {
	return q.client.Delete().Namespace(q.ns).Resource(q.plural).Name(name).Body(options).Do().Error()
}

func (q *queuejobs) Get(name string) (*v1alpha.QueueJob, error) {
	var result v1alpha.QueueJob
	error := q.client.Get().Namespace(q.ns).Name(name).Resource(q.plural).Do().Into(&result)
	return &result, error
}

func (q *queuejobs) List(opts metav1.ListOptions) (*v1alpha.QueueJob, error) {
	var result v1alpha.QueueJob
	err := q.client.Get().Namespace(q.ns).Resource(q.plural).VersionedParams(&opts, q.codec).Do().Into(&result)
	return &result, err
}

// Create a new List watch for our TPR
func (q *queuejobs) NewListWatch() *cache.ListWatch {
	return cache.NewListWatchFromClient(q.client, q.plural, q.ns, fields.Everything())
}
