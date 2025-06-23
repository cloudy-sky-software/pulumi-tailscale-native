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
  @visibility(Lifecycle.Read) key: string;
  @visibility(Lifecycle.Read) created?: string;
  @visibility(Lifecycle.Read) expires: string;
  @visibility(Lifecycle.Create) capabilities: KeyCapabilities;
  @visibility(Lifecycle.Create) expirySeconds: int64;
}
