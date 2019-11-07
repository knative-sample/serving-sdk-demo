package app

import (
	"github.com/spf13/cobra"
)

type Options struct {
	Version    bool
	MasterURL  string
	Kubeconfig string
}

func (s *Options) SetOps(ac *cobra.Command) {
	ac.Flags().StringVar(&s.MasterURL, "master", s.MasterURL, "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	ac.Flags().StringVar(&s.Kubeconfig, "kubeconfig", s.Kubeconfig, "Path to a kubeconfig. Only required if out-of-cluster.")
}
