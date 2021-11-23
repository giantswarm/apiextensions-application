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
