module github.com/isovalent/fake

go 1.15

require (
	github.com/cilium/cilium v1.9.3
	github.com/cilium/hubble v0.7.1
	github.com/golang/protobuf v1.4.3
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/protobuf v1.25.0
)

// cilium replace directives, keep in sync
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.8
	k8s.io/client-go => github.com/cilium/client-go v0.0.0-20201223004022-d3a5d921c1db
	sigs.k8s.io/controller-tools => github.com/christarazi/controller-tools v0.3.1-0.20200911184030-7e668c1fb4c2
	sigs.k8s.io/structured-merge-diff/v4 => github.com/christarazi/structured-merge-diff/v4 v4.0.2-0.20200917183246-1cc601931628
)
