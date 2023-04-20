import * as tailscale from "@cloudyskysoftware/pulumi-tailscale-native";

const devices = new tailscale.tailnet.listDevicesOutput({
  tailnet: "yourtailnetname",
});

devices.apply((d) => console.log(d.items));
