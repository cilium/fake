module github.com/cilium/fake/flow

go 1.18

require (
	github.com/cilium/cilium v1.13.0-rc1.0.20221007133919-8b1666ad3fd1
	github.com/cilium/fake v0.3.0
	github.com/google/go-cmp v0.5.9
	github.com/stretchr/testify v1.8.0
	golang.org/x/net v0.0.0-20220812174116-3211cb980234
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220815135757-37a418bb8959 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.51-0.20220729113855-5b94b11b46fc
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20220829170633-5d7dfb1129f7
	k8s.io/client-go => github.com/cilium/client-go v0.0.0-20220824093223-b6557c021e53
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)

replace github.com/cilium/fake => ../
