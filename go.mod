module github.com/isovalent/fake

go 1.17

require (
	github.com/cilium/cilium v1.11.0
	github.com/cilium/hubble v0.9.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.6-0.20200504143853-81378bbcd8a1
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20220105145211-5b0dc2dfae98
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb // indirect
	google.golang.org/grpc v1.43.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d
	sigs.k8s.io/controller-tools => github.com/christarazi/controller-tools v0.6.2
)
