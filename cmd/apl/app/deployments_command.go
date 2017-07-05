package app

import (
	"fmt"
	//"time"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
	"k8s.io/client-go/pkg/api/v1"
	"strconv"
	//"k8s.io/kubernetes/pkg/printers"
	"time"
)

var deploymentParams apl.DeploymentParams

// NewDeploymentsCommand Creates a cobra command for Deployments
func NewDeploymentsCommand() *cobra.Command {

	cmd := createListCommand(cmdListDeployments, "deployments", "")
	getCmd := createGetCommand(cmdGetDeployments, "deployment", "")
	createCmd := createCreateCommand(cmdCreateDeployments, "deployment", "")
	updateCmd := createUpdateCommand(cmdUpdateDeployments, "deployment", "")
	deleteCmd := createDeleteCommand(cmdDeleteDeployments, "deployment", "")
	podsCmd := &cobra.Command{
		Use:   "pods [deployment-id]",
		Short: fmt.Sprintf("get a list of pods"),
		Long:  "",
		Run:   cmdGetDeploymentPods,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return checkCommandHasIDInArgs(args, "id")
		},
	}

	// command flags
	//cmd.Flags().IntVar(&deploymentParams.ReleaseVersion, "release-version", 0, "Filter deployments by release_version")

	cmd.Flags().StringVar(&deploymentParams.Name, "name", "", "Filter deployments by name")
	cmd.Flags().StringVar(&deploymentParams.StackVersionID, "stack-version-id", "", "Filter deployments by stack_version_id")
	cmd.Flags().StringVar(&deploymentParams.ProjectID, "project-id", "", "Filter deployments by project_id")
	cmd.Flags().StringVar(&deploymentParams.WorkloadID, "workload-id", "", "Filter deployments by workload_id")
	cmd.Flags().StringVar(&deploymentParams.LeaseType, "lease-type", "", "Filter deployments by lease_type")
	cmd.Flags().StringVar(&deploymentParams.WorkloadName, "workload-name", "", "Filter deployments by workload_name")
	cmd.Flags().StringVar(&deploymentParams.LeaseExpiration, "lease-expiration", "", "Filter deployments by lease_expiration")
	cmd.Flags().StringVar(&deploymentParams.QosLevel, "qos-level", "", "Filter deployments by qos_level")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)
	cmd.AddCommand(podsCmd)

	return cmd
}

// cmdListDeployments returns a list of deployments
func cmdListDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&deploymentParams, aplSvc.Deployments.List)

	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Deployment), fields)
	}
}

// cmdGetDeployments gets a specified deployment by deployment-id
func cmdGetDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Deployments.Get)

	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.(apl.Deployment), fields)
	}
}

func cmdCreateDeployments(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.DeploymentCreateInput{}
	runCreateCommand(in, aplSvs.Deployments.Create)
}

func cmdUpdateDeployments(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.DeploymentUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Deployments.Update)
}

func cmdDeleteDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Deployments.Delete)
}

func cmdGetDeploymentPods(ccmd *cobra.Command, args []string) {

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

}

// Copied from
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
