# fah-sidecar

Sidecar app that works with [fah-collector](https://github.com/ebiiim/fah-collector) by exporting FAH data.

## Usage

### Version 1

Requirements: `fah-collector` versions in [1, 2]

#### Usage

`./fah-sidecar [-host] [-port] [-interval] [-livenessport] [-insecure] [-nodename] COLLECTOR_ENDPOINT_URL`

- `COLLECTOR_ENDPOINT_URL`
  - `fah-collector` v1: HTTP(S)://{COLLECTOR_HOST}/{SIDECAR_IDENTIFIER}
    - SIDECAR_IDENTIFIER: E.g., hostname, Pod name.
  - `fah-collector` v2: HTTP(S)://{COLLECTOR_HOST}
- `host`: FAH Telnet address. Default is `localhost` that works fine in a K8s Pod.
- `port`: FAH Telnet port. Default is `36330`.
- `nodename`: K8s nodename or any other identifier. Default is `""`.
- and so on, see `./fah-sidecar -h`.
