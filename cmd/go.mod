module github.com/cilium/fake/cmd

go 1.17

require (
	github.com/cilium/cilium v1.11.5
	github.com/cilium/fake v0.0.0-00010101000000-000000000000
	github.com/cilium/fake/flow v0.0.0-00010101000000-000000000000
	github.com/cilium/hubble v0.9.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.6-0.20200504143853-81378bbcd8a1
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/fatih/color v1.13.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.0.0-20220107192237-5cfca573fb4d // indirect
	golang.org/x/sys v0.0.0-20220110181412-a018aaa089fe // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 // indirect
	google.golang.org/grpc v1.43.0 // indirect
)

// Replace directives from github.com/cilium/cilium. Keep in sync when updating Cilium!
replace (
	github.com/miekg/dns => github.com/cilium/dns v1.1.4-0.20190417235132-8e25ec9a0ff3
	github.com/optiopay/kafka => github.com/cilium/kafka v0.0.0-20180809090225-01ce283b732b
	go.universe.tf/metallb => github.com/cilium/metallb v0.1.1-0.20210831235406-48667b93284d
	sigs.k8s.io/controller-tools => github.com/cilium/controller-tools v0.6.2
)

replace github.com/cilium/fake => ../

replace github.com/cilium/fake/flow => ../flow/
