<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [v0.1.0] - 2024-03-22

### Features

- (chain) [#9](https://github.com/zkMeLabs/mechain/pull/9) Set prefix to mc and denom to zkme chain name to mechain


### Improvement

- (chore) [#33](https://github.com/zkMeLabs/mechain/pull/33) Fix test after remove recovery/incentives/revenue/vesting/inflation/claims module and remove upgrades.
- (dev) [#38](https://github.com/zkMeLabs/mechain/pull/38) Add dev.js script for development and testing.
- (dev) [#40](https://github.com/zkMeLabs/mechain/pull/40) Add four quick command and fix stop node bug.

### Bug Fixes



### State Machine Breaking

- (recovery) [#27](https://github.com/zkMeLabs/mechain/pull/27) Remove `x/recovery` module.
- (incentives) [#28](https://github.com/zkMeLabs/mechain/pull/28) Remove `x/incentives` module.
- (revenue) [#29](https://github.com/zkMeLabs/mechain/pull/29) Remove `x/revenue` module.
- (vesting) [#30](https://github.com/zkMeLabs/mechain/pull/30) Remove `x/vesting` module.
- (inflation) [#31](https://github.com/zkMeLabs/mechain/pull/31) Remove `x/inflation` module.
- (claims) [#32](https://github.com/zkMeLabs/mechain/pull/32) Remove `x/claims` module.
- (evm) [#35](https://github.com/zkMeLabs/mechain/pull/35) Enable EIP 3855 for solidity push0 instruction.
- (deps) [#43](https://github.com/zkMeLabs/mechain/pull/43) Bump Cosmos-SDK to v0.47.2 and ibc-go to v7.2.0.


### API Breaking

