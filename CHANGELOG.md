# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

### Changed

- Correct wording for deprecation of storage field in Catalog CRD

## [0.4.0] - 2022-06-08

### Added

- Extend `Catalog` CRD with repository list.

## [0.3.1] - 2022-02-25

### Fixed

- Remove compatible providers validation for `AppCatalogEntry` as its overly strict.

## [0.3.0] - 2021-12-21

### Added

- Add note to `.spec.namespace` for `App` CRD that this field cannot be changed.

## [0.2.0] - 2021-12-15

### Added

- Add `Created At` printer column for `App` CRD.
- Add `Installed Version` printer column for `App` CRD for `-o wide` output.

## [0.1.0] - 2021-11-23

### Added

- Move existing APIs from `giantswarm/apiextensions` to this repository.


[Unreleased]: https://github.com/giantswarm/apiextensions-application/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/giantswarm/apiextensions-application/compare/v0.3.1...v0.4.0
[0.3.1]: https://github.com/giantswarm/apiextensions-application/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/apiextensions-application/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/apiextensions-application/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/apiextensions-application/releases/tag/v0.1.0
