module github.com/cilium/fake/flow

go 1.22.0

toolchain go1.23.1

require (
	github.com/cilium/cilium v1.16.6
	github.com/cilium/fake v0.7.0
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.10.0
	golang.org/x/net v0.34.0
	google.golang.org/protobuf v1.36.4
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241206012308-a4fef0638583 // indirect
	google.golang.org/grpc v1.68.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20220829170633-5d7dfb1129f7
	// Using private fork of controller-tools. See commit msg for more context
	// as to why we are using a private fork.
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.8.0-2
)
