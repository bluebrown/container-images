# netutils

A small alpine image containing some useful network utilities. The main purpose of this image is debugging and troubleshooting. The container runs as root. So there are no security measurements in place. The container is not intended to be used for production.

## Installed Utilities

- telnet
- netcat
- socat
- dig
- nslookup
- curl

## Running the image

The below commands drops you into an interactive shell. Use your favorite networking tool from there.

### Docker

```bash
docker run --rm -it bluebrown/netutils
```

### Kubernetes

```bash
kubectl run debug --rm -ti --image=bluebrown/netutils
```
