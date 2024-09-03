## v0.1.0, v0.1.1

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

## v0.0.14

- Update pulumi-provider-framework to the latest version
  - This fixes an issue with resource refreshes that was incorrectly updating
    the stashed inputs in the outputs state with read-only properties from the
    provider which would result in diffs with the original inputs.

## v0.0.13

Fix bug with resource reads that causes path param inputs to get removed

## v0.0.12

Upgrade pulumi-provider-framework to the latest

## v0.0.11

Fix return type of `GET /tailnet/{tailnet}/devices`

## v0.0.10

- Regen schema with the latest pulschema to fix an issue with output properties.
- Regen SDKs.

## v0.0.8-v0.0.9

Add support for OAuth2 client credentials auth flow in addition to API key-based auth.

## v0.0.7

Prepare for public release.

## v0.0.1-v0.0.6

Initial release
