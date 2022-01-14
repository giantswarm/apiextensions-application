[![CircleCI](https://circleci.com/gh/giantswarm/apiextensions-application.svg?style=shield)](https://circleci.com/gh/giantswarm/apiextensions-application)

# apiextensions-application

This repository defines the Giant Swarm Kubernetes APIs in the `application.giantswarm.io` group including the following:

- `App`
- `AppCatalog` (deprecated in favor of `Catalog`)
- `AppCatalogEntry`
- `Catalog`
- `Chart`

Examples CRs for these can be found in the `docs/cr` directory.

**Note: These APIs were originally defined in `giantswarm/apiextensions` and were moved here to simplify dependency graphs.**

## Code Generation

Code generation is used to generate:

- `DeepCopy` methods for each Go type to satisfy the `runtime.Object` interface (`zz_generated.deepcopy.go`).
- CRDs for each API in YAML format to be applied to a Kubernetes cluster before creating CRs using that API (`config/crd`).

Regenerate these files using `make generate`. Other options can be viewed in `Makefile.custom.mk`.

Additionally, example CRs are generated from code in Go tests using `go test ./... -update`.

## Update docs

When a new release is created the reference needs to be updated in the [docs config].

## Sync CRDs

All CRDs are embedded in [apptestctl] and the Chart CRD is embedded in [app-operator].

### apptestctl

```sh
$ make sync-crds
```

### app-operator

```sh
$ make sync-chart-crd
```

[app-operator]: https://github.com/giantswarm/app-operator
[apptestctl]: https://github.com/giantswarm/apptestctl
[docs config]: https://github.com/giantswarm/docs/blob/806b6383c51fd6f1c4b78ca32203cbb8071fb935/scripts/update-crd-reference/config.yaml#L8-L11
