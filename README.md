# fake - generate random networking data for testing and benchmarks

[![Go Reference](https://pkg.go.dev/badge/github.com/cilium/fake.svg)](https://pkg.go.dev/github.com/cilium/fake)
[![CI](https://github.com/cilium/fake/actions/workflows/tests.yml/badge.svg)](https://github.com/cilium/fake/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/cilium/fake)](https://goreportcard.com/report/github.com/cilium/fake)

Generate realistic random IPs, MACs, ports, Kubernetes names, and [Hubble] network
flows — with fine-grained control over every field.

## Quick start

```go
import "github.com/cilium/fake"

f := fake.New()

f.IP()                              // "10.244.3.87" or "fd00::a1b2:..."
f.IP(fake.WithIPv4())               // always IPv4
f.IP(fake.WithIPCIDR("10.0.0.0/8")) // always in 10.0.0.0/8

f.MAC()                             // "3e:a7:01:cc:9f:12"
f.Port()                            // 1–65535
f.Port(fake.WithPortDynamic())      // 49152–65535
```

## Kubernetes generators

```go
f := fake.New()

f.K8sNamespace() // "kube-system" or "grafana-prod"
f.K8sPodName()   // "relay-x8f2k"
f.K8sNodeName()  // "brave-falcon"
f.K8sLabels()    // ["app=hubble", "team=relay"]
```

## Generating Hubble flows

The `flow` sub-package builds complete [`flowpb.Flow`][flowpb] protobufs suitable
for testing Hubble pipelines, exporters, or UIs.

```go
import "github.com/cilium/fake/flow"

ff := flow.NewFaker()

// A fully random flow
f := ff.NewFlow()

// A dropped flow on a specific node
f = ff.NewFlow(
    flow.WithFlowVerdict(flowpb.Verdict_DROPPED),
    flow.WithFlowDropReason(flowpb.DropReason_POLICY_DENIED),
    flow.WithFlowNodeName("worker-03"),
)
```

## Deterministic output

Pass your own source to get reproducible sequences — useful for golden tests
or benchmarks.

```go
import randv2 "math/rand/v2"

seed := [32]byte{} // fixed seed
f := fake.NewWithSource(randv2.NewChaCha8(seed))
fmt.Println(f.IP()) // always the same
```

## Performance

Each `Faker` instance is independent and lock-free. For parallel workloads,
create one per goroutine:

```go
for range runtime.GOMAXPROCS(0) {
    go func() {
        f := fake.New()
        for {
            process(f.IP())
        }
    }()
}
```

Package-level convenience functions (`fake.IP()`, `fake.Name()`, ...) use a
global instance behind a mutex and are slower under contention.

## CLI

A small CLI wraps the library for shell scripts and quick iteration.

```
$ fake ip --ipv4 --count 3
192.0.2.41
10.128.7.200
172.21.0.14

$ fake flow | jq .verdict
"FORWARDED"
```

Install with:

    make install                    # → /usr/local/bin/fake
    BINDIR=~/.local/bin make install

## Modules

| Module | Purpose |
|--------|---------|
| [`github.com/cilium/fake`][1] | Generic generators (IPs, MACs, ports, names, K8s resources) |
| [`github.com/cilium/fake/flow`][2] | [Hubble] flow protobuf generation |
| `github.com/cilium/fake/cmd` | CLI binary |

## Contributing

Contributions are welcome! The options pattern makes it straightforward to add
new constraints to existing generators or to introduce new ones.

[1]: https://pkg.go.dev/github.com/cilium/fake
[2]: https://pkg.go.dev/github.com/cilium/fake/flow
[flowpb]: https://pkg.go.dev/github.com/cilium/cilium/api/v1/flow#Flow
[Cilium]: https://github.com/cilium/cilium
[Hubble]: https://github.com/cilium/hubble
