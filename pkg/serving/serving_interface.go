package serving

import (
	"time"

	v1 "k8s.io/api/core/v1"

	"github.com/kubeflow/arena/pkg/apis/types"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServingJob defines a serving job
type ServingJob interface {
	// GetName returns the job name
	Name() string
	// GetNamespace returns the namespace
	Namespace() string
	// Type returns the type
	Type() types.ServingJobType
	// Version returns the job version
	Version() string
	// Pods returns the job pods
	Pods() []*v1.Pod
	// Deployment returns the deployment
	Deployment() *appv1.Deployment
	// Service returns the job services
	Services() []*v1.Service
	// Age returns the job age
	Age() time.Duration
	// Get start time
	StartTime() *metav1.Time
	// Endpoints return the endpoints
	Endpoints() []types.Endpoint
	// RequestGPUs returns the gpus which serving job owned
	RequestGPUs() int
	// RequestGPUMemory returns the gpu memory,only for gpushare
	RequestGPUMemory() int
	// DesiredInstances return the desired instances count
	DesiredInstances() int
	// AvailableInstances returns the available instances
	AvailableInstances() int
	// Convert2JobInfo convert to ServingJobInfo
	Convert2JobInfo() types.ServingJobInfo
}

// Processer is used to process serving jobs
type Processer interface {
	// Type returns the processer type
	Type() types.ServingJobType
	// IsSupported is used to check the processer support the serving job or not
	IsSupported(namespace, name, version string) bool
	// IsEnabled returns the processer is enabled or not
	IsEnabled() bool
	// ListServingJobs is used to list serving jobs
	ListServingJobs(namespace string, allNamespace bool) ([]ServingJob, error)
	// GetServingJob is used to get serving job
	GetServingJobs(namespace, name, version string) ([]ServingJob, error)
	// FilterServingJobs is used to filter serving jobs
	FilterServingJobs(namespace string, allNamespace bool, filter string) ([]ServingJob, error)
}