module github.com/cilium/fake/flow

go 1.24.0

require (
	github.com/cilium/cilium v1.19.0
	github.com/cilium/fake v0.7.0
	github.com/google/go-cmp v0.7.0
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.11.1
	golang.org/x/net v0.50.0
	google.golang.org/protobuf v1.36.11
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/sys v0.40.0 // indirect
	golang.org/x/text v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250721164621-a45f3dfb1074 // indirect
	google.golang.org/grpc v1.74.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!

// Using private fork of controller-tools. See commit msg for more context
// as to why we are using a private fork.
replace sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.16.5-1

// Using private fork of gobgp. See commit msg for more context as to why we
// are using a private fork.
replace github.com/osrg/gobgp/v3 => github.com/cilium/gobgp/v3 2c4e905b1198

// Using private fork of k8s.io/apimachinery with additional patches to support
// `omitzero` json tags in convertors used by cilium.
replace k8s.io/apimachinery => github.com/cilium/apimachinery v0.33.3-1

replace github.com/cilium/fake => ../
