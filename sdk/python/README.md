# Pulumi Native Provider for Tailscale (Preview)

[Tailscale](https://tailscale.com/) connects your team's devices and development environments for easy access to remote resources.

:information_source: This provider uses Tailscale's API directly.

> This provider was generated using [`pulschema`](https://github.com/cloudy-sky-software/pulschema) and [`pulumi-provider-framework`](https://github.com/cloudy-sky-software/pulumi-provider-framework).

## Why Is This Called Tailscale Native?

Despite the fact that all Pulumi providers created by Cloudy Sky Software being "native" Pulumi providers, by default, there is a [prior Pulumi provider for Tailscale](https://github.com/pulumi/pulumi-tailscale), albeit a TF-bridged one. So this provider had to be renamed to avoid naming conflicts, specifically in language package registries such as PyPi and Nuget where the packages are not namespaced under an organization or user unlike npm.

## Package SDKs

- Node.js: https://www.npmjs.com/package/@cloudyskysoftware/pulumi-tailscale-native
- Python: https://pypi.org/project/pulumi_tailscale_native/
- .NET: https://www.nuget.org/packages/Pulumi.TailscaleNative
- Go: `import github.com/cloudy-sky-software/pulumi-tailscale-native/sdk/go/tailscale`

## Using The Provider

### OAuth2 Client

A quote from Tailscale's own docs:

> OAuth clients provide a framework for delegated and scoped access to the Tailscale API. An OAuth client creates access tokens for scoped API access, instead of using a fully-permitted access token which expires after 90 days (or less if the fully-permitted access token has a shorter expiry). OAuth clients support the principle of least privilege by allowing fine-grained control on the access granted to the client using scopes, unlike a fully-permitted access token which grants full access to the Tailscale API.

- Create an OAuth client on https://login.tailscale.com/admin/settings/oauth.
- Grant the appropriate scopes (permissions) to the client.
- Copy the client ID and secret.
- Set the client ID with `pulumi config set tailscale-native:clientId <clientID>`.
- Set the client secret with `pulumi config set --secret tailscale-native:clientSecret <clientSecret>`.
- Optional: Remove the API key config `tailscale-native:apiKey` from your stack config if you previously used that method of auth.

Read the full official docs at https://tailscale.com/kb/1215/oauth-clients/.

### API Key

You'll need an API access token (aka API key). Follow Tailscale's [docs](https://tailscale.com/kb/1101/api/?q=API%20access#authentication) for creating one or head straight to your tailnet [admin console](https://login.tailscale.com/admin/settings/keys) to create one.
Then set the API key as a secret with `pulumi config set --secret tailscale-native:apiKey`.

## Releasing A New Version

:info: Switch to the `main` branch first and get the latest `git pull origin main && git fetch`. Check what the last release tag was.

1. Regular releases should just increment the patch version unless a minor or a major (breaking changes) version bump is warranted.
1. Update the `CHANGELOG.md` with notes about what will be included in this release.
1. Commit the changelog with `git commit -am "vX.Y.Z"` or something similar and push `git push origin main`.
1. Tag the commit with the release version by running

   ```bash
   git tag vX.Y.Z
   git tag sdk/vX.Y.Z
   ```

1. Push the tags.

   ```bash
   git push --tags
   ```
