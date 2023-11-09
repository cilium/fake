module github.com/cilium/fake/flow

go 1.20

require (
	github.com/cilium/cilium v1.14.0
	github.com/cilium/fake v0.5.0
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.8.4
	golang.org/x/net v0.18.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230731193218-e0aa005b6bdf // indirect
	google.golang.org/grpc v1.57.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.51-0.20220729113855-5b94b11b46fc
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20220829170633-5d7dfb1129f7
	k8s.io/client-go => github.com/cilium/client-go v0.27.2-fix
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)

replace github.com/cilium/fake => ../
