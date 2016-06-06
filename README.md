# kempctl

A CLI to manage Kemp load balancers.

## Getting `kempctl`

Download the latest release: https://github.com/giantswarm/kempctl/releases/latest

Clone the git repository: https://github.com/giantswarm/kempctl.git

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
- IRC: #[giantswarm](irc://irc.freenode.org:6667/#giantswarm) on freenode.org
- Bugs: [issues](https://github.com/giantswarm/kempctl/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches, the contribution workflow as well as reporting bugs.

## License

PROJECT is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.