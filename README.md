# kempctl

[![Go Report Card](https://goreportcard.com/badge/github.com/giantswarm/kempctl)](https://goreportcard.com/report/github.com/giantswarm/kempctl) [![Godoc](https://godoc.org/github.com/giantswarm/kempctl?status.svg)](http://godoc.org/github.com/giantswarm/kempctl)

A CLI to manage Kemp load balancers.

## Getting `kempctl`

Download the latest release: https://github.com/giantswarm/kempctl/releases/latest

Clone the git repository: https://github.com/giantswarm/kempctl.git

#### Dependencies

- [github.com/giantswarm/kemp-client](https://github.com/giantswarm/kemp-client)

## Running `kempctl`

```
kempctl --endpoint=https://1.2.3.4/access/ --user=foo --password=bar get cachesize

# list all virtual services
kempctl --endpoint=https://1.2.3.4/access/ --user=foo --password=bar virtual list

# show details and real servers of a specific virtual service
kempctl --endpoint=https://1.2.3.4/access/ --user=foo --password=bar virtual show 5
```

## Contact

- Mailing list: [giantswarm](https://groups.google.com/forum/!forum/giantswarm)
- Bugs: [issues](https://github.com/giantswarm/kempctl/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches, the contribution workflow as well as reporting bugs.

## License

`kempctl` is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
