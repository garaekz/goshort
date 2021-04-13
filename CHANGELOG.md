# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.2] - 2021-04-13

### Changed

- Minor change to set title/meta when embeded

## 2.0.1 - 2021-04-13

### Added

- `migrations/goshort.sql` file to import DB tables
- `config/local.yml` file with example data

## 2.0.0 - 2021-04-12

### Changed

- All of the backend has changed, we used a starter kit and followed some guidelines to keep this code clean.
- We re-structured our VueJS frontend, we give it some flavor using TailwindCSS and got a new design by a great designer.

### Removed

- We completely removed TLD check as many people use local URL's and some TLD's are pretety long, we expect this as a good change, we'll keep an eye on this to prevent abuse.

### Added

- We now keep our list of shorted URL's in the browser using the Local Storage (hope this helps!).
-

## [1.2.4] - 2020-01-21

### Changed

- Rollbacks 1.2.1 patch and replace it with server side formatting, turn base URL to lowercase mantaining case sensitive path and query parameters

## [1.2.3] - 2020-01-19

### Added

- Babel polyfill to old browsers compatibility

Thanks to **Yan Edy Chota Castillo** who first encountered this bug in production!

## [1.2.2] - 2020-01-19

### Added

- Implemented Travis CI to workflow

### Changed

- Refactor regex
- Made changes to be `golangci-lint` compliant

## 1.2.1 - 2020-01-19

### Changed

- Now accepts uppercase URL and parsed it to lowercase

## [1.2.0] - 2020-01-18

### Added

- Logo
- Temporal shortened list
- QR Codes
- Github link
- Copy link of temporal list

### Changed

- Visuals
- Sass breakpoints

## [1.1.0] - 2020-01-15

### Added

- Changelog file
- MIT License
- VueJS Frontend
- URL Redirection
- 404 Page
- Implemented regexp to help with valid URL's

### Changed

- URL recognition algorithm

## [1.0.0] - 2020-01-07

### Added

- API Version
- Create and find code
- Initial file structure

[2.0.2]: https://github.com/garaekz/goshort/compare/v2.0.1...v2.0.2
[2.0.1]: https://github.com/garaekz/goshort/compare/v2.0.0...v2.0.1
[1.2.4]: https://github.com/garaekz/goshort/compare/v1.2.3...v1.2.4
[1.2.3]: https://github.com/garaekz/goshort/compare/v1.2.2...v1.2.3
[1.2.2]: https://github.com/garaekz/goshort/compare/v1.2.0...v1.2.2
[1.2.0]: https://github.com/garaekz/goshort/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/garaekz/goshort/compare/v1.0...v1.1.0
[1.0.0]: https://github.com/garaekz/goshort/releases/tag/v1.0
