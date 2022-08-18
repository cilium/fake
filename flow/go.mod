module github.com/cilium/fake/flow

go 1.17

require (
	github.com/cilium/cilium v1.12.1
	github.com/cilium/fake v0.1.0
	github.com/google/go-cmp v0.5.8
	github.com/stretchr/testify v1.8.0
	golang.org/x/net v0.0.0-20220812174116-3211cb980234
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220817070843-5a390386f1f2 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220815135757-37a418bb8959 // indirect
	google.golang.org/grpc v1.48.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20220713202445-9066eee9e0be
	k8s.io/client-go => github.com/cilium/client-go v0.0.0-20220808211112-706606a736d5
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)
