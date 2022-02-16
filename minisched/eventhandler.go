package minisched

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func addAllEventHandlers(
	sched *Scheduler,
	informerFactory informers.SharedInformerFactory,
) {
	// unscheduled pod
	informerFactory.Core().V1().Pods().Informer().AddEventHandler(
		cache.FilteringResourceEventHandler{
			FilterFunc: func(obj interface{}) bool {
				switch t := obj.(type) {
				case *v1.Pod:
					return !assignedPod(t)
				default:
					return false
				}
			},
			Handler: cache.ResourceEventHandlerFuncs{
				// Consider only adding.
				AddFunc: sched.addPodToSchedulingQueue,
			},
		},
	)
}

// assignedPod selects pods that are assigned (scheduled and running).
func assignedPod(pod *v1.Pod) bool {
	return len(pod.Spec.NodeName) != 0
}

func (sched *Scheduler) addPodToSchedulingQueue(obj interface{}) {
	pod := obj.(*v1.Pod)

	if err := sched.SchedulingQueue.Add(pod); err != nil {
		utilruntime.HandleError(fmt.Errorf("unable to queue %T: %v", obj, err))
	}
}
