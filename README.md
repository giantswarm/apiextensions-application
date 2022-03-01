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

## Release Process

After creating a release the following steps are needed.

### Update apiextensions

Update the reference in the [apiextensions] then run `make generate` in apiextensions.
`opsct ensure crds` uses the master branch of the [crds-common] chart.

### Update docs

Update the reference in the [docs config].

### Sync CRDs

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
[apiextensions]: https://github.com/giantswarm/apiextensions/blob/ae4226f518e758a88c9d05edb54fbf22810b93b4/hack/assets.go#L88
[apptestctl]: https://github.com/giantswarm/apptestctl
[crds-common]: https://github.com/giantswarm/apiextensions/tree/master/helm/crds-common/templates
[docs config]: https://github.com/giantswarm/docs/blob/806b6383c51fd6f1c4b78ca32203cbb8071fb935/scripts/update-crd-reference/config.yaml#L8-L11
