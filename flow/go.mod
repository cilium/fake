module github.com/cilium/fake/flow

go 1.17

require (
	github.com/cilium/cilium v1.11.5
	github.com/cilium/fake v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.6
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20220107192237-5cfca573fb4d
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/viper v1.10.0 // indirect
	golang.org/x/sys v0.0.0-20220110181412-a018aaa089fe // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.43.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)

replace github.com/cilium/fake => ../
