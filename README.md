# fah-sidecar

Sidecar app that works with [fah-collector](https://github.com/ebiiim/fah-sidecar) by exporting FAH data.

## Usage

### Version 1

Requirements: `fah-collector` version 1

#### Usage

`./fah-sidecar [-host] [-port] [-interval] [-livenessport] [-insecure] COLLECTOR_ENDPOINT_URL`

- `COLLECTOR_ENDPOINT_URL`: HTTP(S)://{COLLECTOR_HOST}/{SIDECAR_IDENTIFIER}
  - SIDECAR_IDENTIFIER: E.g., hostname, Pod name.
- `host`: FAH Telnet address. Default is `localhost` that works fine in a K8s Pod.
- `port`: FAH Telnet port. Default is `36330`.
- and so on, see `./fah-sidecar -h`.
