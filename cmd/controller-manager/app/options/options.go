/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// The controller is responsible for running control loops that reconcile
// the state of service catalog API resources with service brokers, service
// classes, service instances, and service instance credentials.

package options

import (
	"time"

	"github.com/spf13/pflag"
	utilfeature "k8s.io/apiserver/pkg/util/feature"

	osb "sigs.k8s.io/go-open-service-broker-client/v2"
	"github.com/kubernetes-sigs/service-catalog/pkg/apis/componentconfig"
	"github.com/kubernetes-sigs/service-catalog/pkg/controller"
	k8scomponentconfig "github.com/kubernetes-sigs/service-catalog/pkg/kubernetes/pkg/apis/componentconfig"
	"github.com/kubernetes-sigs/service-catalog/pkg/kubernetes/pkg/client/leaderelectionconfig"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

const (
	// Use the same SSL configuration as we use in Catalog API Server.
	// Store generated SSL certificates in a place that won't collide with the
	// k8s core API server.
	certDirectory = "/var/run/kubernetes-service-catalog"
)

// ControllerManagerServer is the main context object for the controller
// manager.
type ControllerManagerServer struct {
	componentconfig.ControllerManagerConfiguration
}

const (
	defaultResyncInterval                         = 5 * time.Minute
	defaultServiceBrokerRelistInterval            = 24 * time.Hour
	defaultContentType                            = "application/json"
	defaultBindAddress                            = "0.0.0.0"
	defaultPort                                   = 8444
	defaultK8sKubeconfigPath                      = "./kubeconfig"
	defaultServiceCatalogKubeconfigPath           = "./service-catalog-kubeconfig"
	defaultOSBAPIContextProfile                   = true
	defaultConcurrentSyncs                        = 5
	defaultLeaderElectionNamespace                = "kube-system"
	defaultReconciliationRetryDuration            = 7 * 24 * time.Hour
	defaultOperationPollingMaximumBackoffDuration = 20 * time.Minute
	defaultOSBAPITimeOut                          = 60 * time.Second
)

var defaultOSBAPIPreferredVersion = osb.LatestAPIVersion().HeaderValue()

// NewControllerManagerServer creates a new ControllerManagerServer with a
// default config.
func NewControllerManagerServer() *ControllerManagerServer {
	s := ControllerManagerServer{
		ControllerManagerConfiguration: componentconfig.ControllerManagerConfiguration{
			Address:                                defaultBindAddress,
			Port:                                   0,
			ContentType:                            defaultContentType,
			K8sKubeconfigPath:                      defaultK8sKubeconfigPath,
			ServiceCatalogKubeconfigPath:           defaultServiceCatalogKubeconfigPath,
			ResyncInterval:                         defaultResyncInterval,
			ServiceBrokerRelistInterval:            defaultServiceBrokerRelistInterval,
			OSBAPIContextProfile:                   defaultOSBAPIContextProfile,
			OSBAPIPreferredVersion:                 defaultOSBAPIPreferredVersion,
			OSBAPITimeOut:                          defaultOSBAPITimeOut,
			ConcurrentSyncs:                        defaultConcurrentSyncs,
			LeaderElection:                         leaderelectionconfig.DefaultLeaderElectionConfiguration(),
			LeaderElectionNamespace:                defaultLeaderElectionNamespace,
			EnableProfiling:                        true,
			EnableContentionProfiling:              false,
			ReconciliationRetryDuration:            defaultReconciliationRetryDuration,
			OperationPollingMaximumBackoffDuration: defaultOperationPollingMaximumBackoffDuration,
			SecureServingOptions:                   genericoptions.NewSecureServingOptions(),
		},
	}
	// set defaults, these will be overridden by user specified flags
	s.SecureServingOptions.BindPort = defaultPort
	s.SecureServingOptions.ServerCert.CertDirectory = certDirectory
	s.LeaderElection.LeaderElect = true
	return &s
}

// AddFlags adds flags for a ControllerManagerServer to the specified FlagSet.
func (s *ControllerManagerServer) AddFlags(fs *pflag.FlagSet) {
	fs.Var(k8scomponentconfig.IPVar{Val: &s.Address}, "address", "DEPRECATED: see --bind-address instead")
	fs.MarkDeprecated("address", "see --bind-address instead")
	fs.Int32Var(&s.Port, "port", 0, "DEPRECATED: see --secure-port instead")
	fs.IntVar(&s.ConcurrentSyncs, "concurrent-syncs", defaultConcurrentSyncs, "Number of concurrent syncs")
	fs.MarkDeprecated("port", "see --secure-port instead")
	fs.StringVar(&s.ContentType, "api-content-type", s.ContentType, "Content type of requests sent to API servers")
	fs.StringVar(&s.K8sAPIServerURL, "k8s-api-server-url", "", "The URL for the k8s API server")
	fs.StringVar(&s.K8sKubeconfigPath, "k8s-kubeconfig", "", "Path to k8s core kubeconfig")
	fs.StringVar(&s.ServiceCatalogAPIServerURL, "service-catalog-api-server-url", "", "The URL for the service-catalog API server")
	fs.StringVar(&s.ServiceCatalogKubeconfigPath, "service-catalog-kubeconfig", "", "Path to service-catalog kubeconfig")
	fs.BoolVar(&s.ServiceCatalogInsecureSkipVerify, "service-catalog-insecure-skip-verify", s.ServiceCatalogInsecureSkipVerify, "Skip verification of the TLS certificate for the service-catalog API server")
	fs.DurationVar(&s.ResyncInterval, "resync-interval", s.ResyncInterval, "The interval on which the controller will resync its informers")
	fs.DurationVar(&s.ServiceBrokerRelistInterval, "broker-relist-interval", s.ServiceBrokerRelistInterval, "The interval on which a broker's catalog is relisted after the broker becomes ready")
	fs.BoolVar(&s.OSBAPIContextProfile, "enable-osb-api-context-profile", s.OSBAPIContextProfile, "This does nothing.")
	fs.MarkHidden("enable-osb-api-context-profile")
	fs.StringVar(&s.OSBAPIPreferredVersion, "osb-api-preferred-version", s.OSBAPIPreferredVersion, "The string to send as the version header.")
	fs.BoolVar(&s.EnableProfiling, "profiling", s.EnableProfiling, "Enable profiling via web interface host:port/debug/pprof/")
	fs.BoolVar(&s.EnableContentionProfiling, "contention-profiling", s.EnableContentionProfiling, "Enable lock contention profiling, if profiling is enabled")
	leaderelectionconfig.BindFlags(&s.LeaderElection, fs)
	fs.StringVar(&s.LeaderElectionNamespace, "leader-election-namespace", s.LeaderElectionNamespace, "Namespace to use for leader election lock")
	fs.DurationVar(&s.ReconciliationRetryDuration, "reconciliation-retry-duration", s.ReconciliationRetryDuration, "The maximum amount of time to retry reconciliations on a resource before failing")
	fs.DurationVar(&s.OperationPollingMaximumBackoffDuration, "operation-polling-maximum-backoff-duration", s.OperationPollingMaximumBackoffDuration, "The maximum amount of time to back-off while polling an OSB API operation")
	fs.DurationVar(&s.OSBAPITimeOut, "osb-api-request-timeout", s.OSBAPITimeOut, "The maximum amount of timeout to any request to the broker.")
	s.SecureServingOptions.AddFlags(fs)
	utilfeature.DefaultMutableFeatureGate.AddFlag(fs)
	fs.StringVar(&s.ClusterIDConfigMapName, "cluster-id-configmap-name", controller.DefaultClusterIDConfigMapName, "k8s name for clusterid configmap")
	fs.StringVar(&s.ClusterIDConfigMapNamespace, "cluster-id-configmap-namespace", controller.DefaultClusterIDConfigMapNamespace, "k8s namespace for clusterid configmap")
}
