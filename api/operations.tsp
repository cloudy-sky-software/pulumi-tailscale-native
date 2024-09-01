import "@typespec/rest";

import "./models/acl.tsp";
import "./models/authKeys.tsp";
import "./models/device.tsp";
import "./models/deviceRoutes.tsp";
import "./models/dns.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

model TailscaleOAuthFlow {
    type: OAuth2FlowType.clientCredentials;
    tokenUrl: "https://api.tailscale.com/api/v2/oauth/token";

    // The goal is to map these scopes to individual operations
    // but these are less useful in the context of the Pulumi
    // provider.
    // TODO: This needs to be a map but is currently blocked.
    // https://github.com/microsoft/typespec/issues/1113
    scopes: [
        "all",
        "all:read",
        "acl",
        "acl:read",
        "devices",
        "devices:read",
        "dns",
        "dns:read",
        "routes",
        "routes:read"
    ];
}

@service({
    title: "Tailscale API",
    version: "2023-01-01",
})
@server("https://api.tailscale.com/api/v2", "Tailscale API")
@useAuth(BasicAuth | OAuth2Auth<[TailscaleOAuthFlow]>)
@doc("This namespace contains operations that can be performed on a Tailscale-connected device.")
namespace Operations {
    @error
    model Error {
        message: string;
    }

    @route("/tailnet/{tailnet}")
    namespace TailnetEndpoints {
        model CommonParameters {
            @path
            @segment("tailnet")
            @doc("For paid plans, your domain is your tailnet. For solo plans, the tailnet is the email you signed up with. So `alice@gmail.com` has the tailnet `alice@gmail.com` since `@gmail.com` is a shared email host. Alternatively, you can specify the value \"-\" to refer to the default tailnet of the authenticated user making the API call.")
            tailnet: "-" | string;
        }

        @route("/acl")
        @get
        op getAcl(...CommonParameters, @header accept: "application/json"): Acl;

        @route("/acl")
        @post
        op createAcl(
            ...CommonParameters,
            @header accept: "application/json",
            @body acl: Acl,
        ): Acl;

        @route("/devices")
        @get
        op listDevices(...CommonParameters): {
            devices: Device[];
        };

        @route("/keys")
        @get
        op listKeys(...CommonParameters): AuthKey[];

        @route("/keys/{id}")
        @get
        op getKey(...CommonParameters, @path id: string): AuthKey;

        @route("/keys")
        @post
        op createKey(...CommonParameters, @body authKey: AuthKey): AuthKey;

        @route("/keys/{id}")
        @delete
        op deleteKey(...CommonParameters, @path id: string): Response<200>;

        @route("/dns/nameservers")
        @get
        op getNameServersConfig(...CommonParameters): NameServers;

        @route("/dns/nameservers")
        @post
        @doc("Replaces the list of DNS nameservers for a tailnet.")
        op setNameServersConfig(
            ...CommonParameters,
            @body nameServer: NameServers,
        ): NameServers;

        @route("/dns/preferences")
        @get
        op getDNSPreferencesConfig(...CommonParameters): NameServersPreference;

        @route("/dns/preferences")
        @post
        @doc("Replaces the DNS preferences for a tailnet.")
        op setDNSPreferencesConfig(
            ...CommonParameters,
            @body prefs: NameServersPreference,
        ): NameServersPreference;

        @route("/dns/searchpaths")
        @get
        op listSearchPaths(...CommonParameters): DnsSearchPaths;

        @route("/dns/searchpaths")
        @post
        op replaceSearchPaths(
            ...CommonParameters,
            @body searchPaths: DnsSearchPaths,
        ): DnsSearchPaths;
    }

    @route("/device/{id}")
    namespace DeviceEndpoints {
        @get op getDevice(@path id: string): Device;

        @delete op deleteDevice(@path id: string): Response<200> | Error;

        @route("/authorized")
        @post
        op authorizeDevice(
            @path id: string,
            @body setAuthorized: SetAuthorized,
        ): Response<200>;

        @route("/tags")
        @post
        op setTags(@path id: string, @body setTags: SetTags): Response<200>;

        @route("/routes")
        @get
        op getRoutesConfig(@path id: string): DeviceRoutes;

        @route("/routes")
        @post
        op setRoutesConfig(
            @path id: string,
            @body setRoutes: SetDeviceRoutes,
        ): DeviceRoutes;

        @route("/key")
        @post
        op setKeyExpiry(
            @path id: string,
            @body setKeyExpiry: SetKeyExpiry,
        ): Response<200>;
    }
}
