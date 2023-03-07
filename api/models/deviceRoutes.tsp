model DeviceRoutes {
    advertisedRoutes: string[];
    enabledRoutes: string[];
}

model SetAuthorized {
    authorized: boolean;
}

model SetDeviceRoutes {
    routes: string[];
}

model SetTags {
    tags: string[];
}

model SetKeyExpiry {
    keyExpiryDisabled: boolean;
}
