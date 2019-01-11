kubectl get nodes -o json \
| jq ".items[]|{name: .metadata.name, externalIP: .status.addresses[1].address}"
