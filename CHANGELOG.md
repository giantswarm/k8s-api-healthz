# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [0.2.0] - 2023-03-06

### Changed

- Require TLS 1.2 or above.

# Fixed

- Fixed wiring issue that made API healthcheck probes to be sent to etcd endpoint instead.

## [0.1.1] - 2020-07-06

### Added

- Add missing `aliyun` CI configuration.

## [0.1.0] - 2020-06-30

### Changed

- Switch from `dep` to go modules.
- Use latest `architect-orb` `0.9.0`.

[Unreleased]: https://github.com/giantswarm/k8s-api-healthz/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/giantswarm/k8s-api-healthz/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/giantswarm/k8s-api-healthz/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/k8s-api-healthz/releases/tag/v0.1.0
