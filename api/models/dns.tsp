model NameServersPreference {
    magicDNS: boolean;
}

model NameServers {
    dns: string[];

    // It's unfortunate we can't use model composition here because
    // pulschema doesn't support the format that CADL emits when
    // composition is used.
    // See https://github.com/microsoft/cadl/issues/826.
    magicDNS: boolean;
}

model DnsSearchPaths {
    searchPaths: string[];
}
