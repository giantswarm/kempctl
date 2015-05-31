# Manage kemp loadmaster via api

A CLI to manage kemp loadbalancers.

See `kempctl help` for more information

Example:
```
kempctl --endpoint=https://1.2.3.4/access/ --user=foo --password=bar get cachesize

# list all virtual services
kempctl --endpoint=https://1.2.3.4/access/ --user=foo --password=bar virtual list

# show details and real servers of a specific virtual service
kempctl --endpoint=https://1.2.3.4/access/ --user=foo --password=bar virtual show 5
```


