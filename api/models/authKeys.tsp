model CreateKey {
    reusable: boolean;
    ephemeral: boolean;
    preauthorized: boolean;
    tags: string[];
}

model DeviceKeyCapabilities {
    create: CreateKey;
}

model KeyCapabilities {
    devices: DeviceKeyCapabilities;
}

model AuthKey {
    @visibility("read") key: string;
    @visibility("read") created?: string;
    @visibility("read") expires: string;
    @visibility("create") capabilities: KeyCapabilities;
    @visibility("create") expirySeconds: int64;
}
