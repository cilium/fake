module github.com/cilium/fake/flow

go 1.20

require (
	github.com/cilium/cilium v1.14.5
	github.com/cilium/fake v0.6.1
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/net v0.24.0
	google.golang.org/protobuf v1.34.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.59.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20220829170633-5d7dfb1129f7
	k8s.io/client-go => github.com/cilium/client-go v0.27.6-fix
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)

replace github.com/cilium/fake => ../
