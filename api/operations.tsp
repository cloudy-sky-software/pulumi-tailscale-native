import "@typespec/rest";

import "./models/acl.tsp";
import "./models/authKeys.tsp";
import "./models/device.tsp";
import "./models/deviceRoutes.tsp";
import "./models/dns.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;

@service({
    title: "Tailscale API",
    version: "2023-01-01",
})
@server("https://api.tailscale.com/api/v2", "Tailscale API")
@useAuth(BasicAuth)
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
        op getAcl(
            @path tailnet: "-" | string,
            @header accept: "application/json"
        ): Acl;

        @route("/acl")
        @post
        op createAcl(
            @path tailnet: "-" | string,
            @header accept: "application/json",
            @body acl: Acl
        ): Acl;

        @route("/devices")
        @get
        op listDevices(@path tailnet: "-" | string): Device[];

        @route("/keys")
        @get
        op listKeys(@path tailnet: "-" | string): AuthKey[];

        @route("/keys/{id}")
        @get
        op getKey(@path tailnet: "-" | string, @path id: string): AuthKey;

        @route("/keys")
        @post
        op createKey(
            @path tailnet: "-" | string,
            @body authKey: AuthKey
        ): AuthKey;

        @route("/keys/{id}")
        @delete
        op deleteKey(
            @path tailnet: "-" | string,
            @path id: string
        ): Response<200>;

        @route("/dns/nameservers")
        @get
        op getNameServers(@path tailnet: "-" | string): NameServers;

        @route("/dns/nameservers")
        @post
        op replaceNameServers(
            @path tailnet: "-" | string,
            @body nameServer: NameServers
        ): NameServers;

        @route("/dns/preferences")
        @get
        op getDNSPreferences(
            @path tailnet: "-" | string
        ): NameServersPreference;

        @route("/dns/preferences")
        @post
        op replaceDNSPreferences(
            @path tailnet: "-" | string,
            @body prefs: NameServersPreference
        ): NameServersPreference;

        @route("/dns/searchpaths")
        @get
        op listSearchPaths(@path tailnet: "-" | string): DnsSearchPaths;

        @route("/dns/searchpaths")
        @post
        op replaceSearchPaths(
            @path tailnet: "-" | string,
            @body searchPaths: DnsSearchPaths
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
            @body setAuthorized: SetAuthorized
        ): Response<200>;

        @route("/tags")
        @post
        op setTags(@path id: string, @body setTags: SetTags): Response<200>;

        @route("/routes")
        @get
        op getRoutes(@path id: string): DeviceRoutes;

        @route("/routes")
        @post
        op setRoutes(
            @path id: string,
            @body setRoutes: SetDeviceRoutes
        ): DeviceRoutes;

        @route("/key")
        @post
        op setKeyExpiry(
            @path id: string,
            @body setKeyExpiry: SetKeyExpiry
        ): Response<200>;
    }
}