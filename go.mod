module github.com/kubernetes-sigs/service-catalog

go 1.13

require (
	github.com/go-openapi/spec v0.20.4
	github.com/google/gofuzz v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/peterbourgon/mergemap v0.0.0-20130613134717-e21c03b7a721
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.0
	github.com/vrischmann/envconfig v1.3.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/net v0.0.0-20211216030914-fe4d6282115f
	gomodules.xyz/jsonpatch/v2 v2.2.0
	k8s.io/api v0.23.1
	k8s.io/apiextensions-apiserver v0.23.1
	k8s.io/apimachinery v0.23.1
	k8s.io/apiserver v0.23.1
	k8s.io/cli-runtime v0.23.1
	k8s.io/client-go v0.23.1
	k8s.io/code-generator v0.23.1
	k8s.io/component-base v0.23.1
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20220104150652-0d48a347b383
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/go-open-service-broker-client/v2 v2.0.0-20210928082444-d912fc53f8bd
	sigs.k8s.io/yaml v1.3.0
)
