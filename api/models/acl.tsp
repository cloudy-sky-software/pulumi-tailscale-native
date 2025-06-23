/**
 * Tailscale ACL rules are "default deny".
 * So the only possible value for an ACL
 * rule is `accept`.
 */
enum AclRuleAction {
  accept,
}

model AclRule {
  action: AclRuleAction;
  ports: string[];
  users: string[];
}

enum SshRuleAction {
  check,
  accept,
}

model SshRule {
  action: SshRuleAction;
  src: string[];
  dst: string[];
  users: string[];
  checkPeriod: string;
}

model NodeAttrs {
  target: string[];
  attr: string[];
}

model Acl {
  acls: AclRule[];
  autoApprovers: Record<string[]>;
  groups: Record<string[]>;
  hosts: Record<string>;
  nodeAttrs: NodeAttrs[];
  ssh: SshRule[];
  tagOwners: Record<string[]>;
  tests: string;
}
