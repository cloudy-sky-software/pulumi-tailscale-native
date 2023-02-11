import * as tailscale from "@cloudyskysoftware/pulumi-tailscale";

const devices = tailscale.tailnet.listDevicesOutput({
  tailnet: "pronty.87@gmail.com",
});

devices.apply(console.log);
