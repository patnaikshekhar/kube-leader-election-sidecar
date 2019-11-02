package main

import (
	"time"
	"os"
	"log"
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

func startElectionProcess(ctx context.Context, kubeconfig string, callback func(bool)) {

	config, err := buildConfig(kubeconfig)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	client := clientset.NewForConfigOrDie(config)

	lock := &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      "coordination-lock",
			Namespace: os.Getenv("POD_NAMESPACE"),
		},
		Client: client.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: os.Getenv("POD_NAME"),
		},
	}

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          lock,
		LeaseDuration: time.Second * 15,
		RenewDeadline: time.Second * 10,
		RetryPeriod:   time.Second * 2,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				callback(true)
			},
			OnStoppedLeading: func() {
				callback(false)
			},
		},
	})
}

func buildConfig(kubeconfig string) (*rest.Config, error) {

	if kubeconfig != "" {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		return config, err
	}

	config, err := rest.InClusterConfig()
	return config, err
}