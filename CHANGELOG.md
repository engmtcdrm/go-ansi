# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added `Strip` function to replace `StripCodes` function.
- Added `Dim` constant to replace `Faint` constant.
- Added `Hide` constant to replace `Hidden` constant.
- Added `ResetHide` constant to replace `ResetHidden` constant.
- Added `ScrollUp` function to replace `ScrollUpN` function.
- Added `ScrollDown` function to replace `ScrollDownN` function.

### Deprecated

- `StripCodes` is deprecated and will be removed in the future. Use `Strip` instead.
- `Faint` is deprecated and will be removed in the future. Use `Dim` instead.
- `Hidden` is deprecated and will be removed in the future. Use `Hide` instead.
- `ResetHidden` is deprecated and will be removed in the future. Use `ResetHide` instead.
- `ScrollUpN` is deprecated and will be removed in the future. Use `ScrollUp` instead.
- `ScrollDownN` is deprecated and will be removed in the future. Use `ScrollDown` instead.

## [v1.0.1] - 2025-01-02

### Added

- Added additional csi for `StripCodes` func

### Changed

- Updated documentation in `README.md`
- Renamed `ScrollNext` to `ScrollUp1`
- Renamed `ScrollPrev` to `ScrollDown1`
- Moved `DoubleUnderline` under Text Formatting

## [v1.0.0] - 2024-12-30

Initial release testing versioning in go. This version is identical to v0.1.0 and was removed.

## [v0.1.0] - 2024-12-30

Initial release
