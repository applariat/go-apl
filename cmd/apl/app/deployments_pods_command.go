package app

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
	"k8s.io/client-go/pkg/api/v1"
	"strconv"
	"time"
)

// NewDeploymentsPodsCommand
func NewDeploymentsPodsCommand() *cobra.Command {

	// Pods
	cmd := &cobra.Command{
		Use:   "pods [deployment-id]",
		Short: fmt.Sprintf("get a list of pods"),
		Long:  "",

		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkCommandHasIDInArgs(args, "deployment-id")
		},

		Run: func(ccmd *cobra.Command, args []string) {

			aplSvc := apl.NewClient()

			output := runGetCommand(args, aplSvc.Deployments.Get)

			if output != nil {
				deploy := output.(apl.Deployment)

				namespace := deploy.Status.(map[string]interface{})["namespace"].(string)
				//namespace := status["namespace"]
				fmt.Println(namespace)

				//fmt.Println(deploy.Status)

				kubeSvc, err := apl.GetKubeClient()
				if err != nil {
					fmt.Println(err)
					return
				}
				pods, err := kubeSvc.Core().Pods(namespace).List(v1.ListOptions{})
				if err != nil {
					panic(err.Error())
				}

				data := make([][]string, len(pods.Items))
				for _, p := range pods.Items {
					podInfo := printPod(&p)
					data = append(data, podInfo)
				}
				header := []string{"Name", "Ready", "Status", "Restarts", "Age"}
				//header := []string{"Name", "Ready", "Status", "Restarts"}
				printTableResults(data, header)

			}

		},
	}

	return cmd

}

//
func printPod(pod *v1.Pod) []string {

	restarts := 0
	totalContainers := len(pod.Spec.Containers)
	readyContainers := 0

	reason := string(pod.Status.Phase)
	if pod.Status.Reason != "" {
		reason = pod.Status.Reason
	}

	initializing := false
	for i := range pod.Status.InitContainerStatuses {
		container := pod.Status.InitContainerStatuses[i]
		restarts += int(container.RestartCount)
		switch {
		case container.State.Terminated != nil && container.State.Terminated.ExitCode == 0:
			continue
		case container.State.Terminated != nil:
			// initialization is failed
			if len(container.State.Terminated.Reason) == 0 {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Init:Signal:%d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("Init:ExitCode:%d", container.State.Terminated.ExitCode)
				}
			} else {
				reason = "Init:" + container.State.Terminated.Reason
			}
			initializing = true
		case container.State.Waiting != nil && len(container.State.Waiting.Reason) > 0 && container.State.Waiting.Reason != "PodInitializing":
			reason = "Init:" + container.State.Waiting.Reason
			initializing = true
		default:
			reason = fmt.Sprintf("Init:%d/%d", i, len(pod.Spec.InitContainers))
			initializing = true
		}
		break
	}

	if !initializing {
		restarts = 0
		for i := len(pod.Status.ContainerStatuses) - 1; i >= 0; i-- {
			container := pod.Status.ContainerStatuses[i]

			restarts += int(container.RestartCount)
			if container.State.Waiting != nil && container.State.Waiting.Reason != "" {
				reason = container.State.Waiting.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason != "" {
				reason = container.State.Terminated.Reason
			} else if container.State.Terminated != nil && container.State.Terminated.Reason == "" {
				if container.State.Terminated.Signal != 0 {
					reason = fmt.Sprintf("Signal:%d", container.State.Terminated.Signal)
				} else {
					reason = fmt.Sprintf("ExitCode:%d", container.State.Terminated.ExitCode)
				}
			} else if container.Ready && container.State.Running != nil {
				readyContainers++
			}
		}
	}

	if pod.DeletionTimestamp != nil && pod.Status.Reason == "NodeLost" {
		reason = "Unknown"
	} else if pod.DeletionTimestamp != nil {
		reason = "Terminating"
	}

	age := "??"
	if !pod.CreationTimestamp.IsZero() {
		age = timeDuration(time.Now().Sub(pod.CreationTimestamp.Time))
	}

	//for _, refs := range pod.GetOwnerReferences() {
	//	fmt.Println(refs.Kind)
	//}

	//annotations := pod.GetAnnotations()
	//fmt.Println(annotations)
	//fmt.Println(annotations["resource"])
	//fmt.Println(annotations["resource"])

	return []string{
		pod.GetName(),
		fmt.Sprintf("%d/%d", readyContainers, totalContainers),
		reason,
		strconv.Itoa(restarts),
		age,
	}

}
