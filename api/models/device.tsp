model Latency {
    latencyMs: float;
    preferred: boolean;
}

model ClientSupports {
    hairPinning: boolean;
    ipv6: boolean;
    pcp: boolean;
    pmp: boolean;
    udp: boolean;
    upnp: boolean;
}

model ClientConnectivity {
    endpoints: string;
    derp: string;
    mappingVariesByDestIP: boolean;
    latency: Record<Latency>;
    clientSupports: ClientSupports;
}

model Device {
    addresses: string[];
    id: string;
    user: string;
    name: string;
    hostname: string;
    clientVersion: string;
    updateAvailable: boolean;
    os: string;
    // TODO: Should these use the CADL date/time types?
    created: string;
    lastSeen: string;
    keyExpiryDisabled: boolean;
    // TODO: Same here.
    expires: string;
    authorized: boolean;
    isExternal: boolean;
    machineKey: string;
    nodeKey: string;
    blocksIncomingConnections: boolean;
    enabledRoutes: string[];
    advertisedRoutes: string[];
    clientConnectivity: ClientConnectivity;
}
