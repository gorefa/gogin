package model

import (
	"fmt"

	"github.com/gorefa/log"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv1beta1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

func Monitoring(ns string) {

	RestConfig, err = GetrestConfig("ifeng-sre")
	c, _ := metricsv1beta1.NewForConfig(RestConfig)

	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	podmetris, _ := c.PodMetricses(ns).Get("test", metav1.GetOptions{})

	for _, value := range podmetris.Containers {
		fmt.Println(value.Usage.Memory())
		fmt.Println(value.Usage.Cpu())
	}
}

type PodStat struct {
	Name         string
	Status       bool
	RestartCount int32
	Message      string
}

type Ingress struct {
	NS   string
	Name string
}

func PodList(cluster, ns string) []PodStat {
	Clientset, err = NewInitK8S(cluster)
	pods, err := Clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	podstatus := []PodStat{}

	for i := 0; i < len(pods.Items); i++ {
		var item PodStat
		item.Name = pods.Items[i].Name
		//item.Status = string(pods.Items[i].Status.Phase)
		//if len(pods.Items[i].Status.ContainerStatuses) > 1 {
		//
		//	continue
		//}

		item.RestartCount = pods.Items[i].Status.ContainerStatuses[0].RestartCount
		item.Status = pods.Items[i].Status.ContainerStatuses[0].Ready
		if !item.Status {
			item.Message = pods.Items[i].Status.ContainerStatuses[0].State.Waiting.Reason
		}

		podstatus = append(podstatus, item)
	}
	//for i, pod := range pods.Items {
	//	var item PodStat
	//	item.Name = pod.Name
	//	item.Status = string(pod.Status.Phase)
	//	item.RestartCount = pod.Status.ContainerStatuses[i].RestartCount
	//	podstatus = append(podstatus, item)
	//}
	return podstatus
}

func IngressList(cluster, ns string) []string {
	Clientset, err = NewInitK8S(cluster)
	ingress, err := Clientset.ExtensionsV1beta1().Ingresses(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf(err, "list %s ingress err", ns)
	}
	var ingresses []string
	for _, ing := range ingress.Items {
		ingresses = append(ingresses, ing.Name)
	}

	return ingresses
}

func DeploymentCreate(cluster, ns string) (string, error) {
	Clientset, err = NewInitK8S(cluster)
	deploymentsClient := Clientset.AppsV1().Deployments(ns)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	log.Info("Create deployment....")
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		log.Fatalf(err, "create deployment failed")
		return "", err
	}
	return result.GetObjectMeta().GetName(), nil
}

func int32Ptr(i int32) *int32 { return &i }

func NodeList(cluster string) []string {
	Clientset, err = NewInitK8S(cluster)
	nodelist, err := Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf(err, "list node err")
	}

	var nodes []string
	for _, node := range nodelist.Items {
		nodes = append(nodes, node.GetName())
	}
	return nodes
}
