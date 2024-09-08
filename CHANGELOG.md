## 0.1.5, 0.1.6

Fix execution error for list-style invokes.

## 0.1.4

Remove unnecessary envelope properties from `get*` and `list*` functions.

## 0.1.3

Upgrade pulumi-provider-framework to fix a bug with validating response codes for DELETE calls.

## 0.1.0-0.1.2

The following resources and functions were renamed to better fit the
semantics of the REST endpoints they use.

Resources:

`DNSPreferences` -> `DNSPreferencesConfig`
`NameServers` -> `NameServersConfig`
`Routes` -> `RoutesConfig`

Functions:

`getDNSPreference` -> `getDNSPreferencesConfig`
`getNameServer` -> `getNameServersConfig`
`getRoute` -> `getRoutesConfig`

## 0.0.14

- Update pulumi-provider-framework to the latest version
  - This fixes an issue with resource refreshes that was incorrectly updating
    the stashed inputs in the outputs state with read-only properties from the
    provider which would result in diffs with the original inputs.

## 0.0.13

Fix bug with resource reads that causes path param inputs to get removed

## 0.0.12

Upgrade pulumi-provider-framework to the latest

## 0.0.11

Fix return type of `GET /tailnet/{tailnet}/devices`

## 0.0.10

- Regen schema with the latest pulschema to fix an issue with output properties.
- Regen SDKs.

## 0.0.8-0.0.9

Add support for OAuth2 client credentials auth flow in addition to API key-based auth.

## 0.0.7

Prepare for public release.

## 0.0.1-0.0.6

Initial release
