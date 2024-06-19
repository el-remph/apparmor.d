// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2024 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package aa

import (
	"fmt"
)

const CAPABILITY Kind = "capability"

func init() {
	requirements[CAPABILITY] = requirement{
		"name": {
			"audit_control", "audit_read", "audit_write", "block_suspend", "bpf",
			"checkpoint_restore", "chown", "dac_override", "dac_read_search",
			"fowner", "fsetid", "ipc_lock", "ipc_owner", "kill", "lease",
			"linux_immutable", "mac_admin", "mac_override", "mknod", "net_admin",
			"net_bind_service", "net_broadcast", "net_raw", "perfmon", "setfcap",
			"setgid", "setpcap", "setuid", "sys_admin", "sys_boot", "sys_chroot",
			"sys_module", "sys_nice", "sys_pacct", "sys_ptrace", "sys_rawio",
			"sys_resource", "sys_time", "sys_tty_config", "syslog", "wake_alarm",
		},
	}
}

type Capability struct {
	RuleBase
	Qualifier
	Names []string
}

func newCapability(q Qualifier, rule rule) (Rule, error) {
	names, err := toValues(CAPABILITY, "name", rule.GetString())
	if err != nil {
		return nil, err
	}
	return &Capability{
		RuleBase:  newBase(rule),
		Qualifier: q,
		Names:     names,
	}, nil
}

func newCapabilityFromLog(log map[string]string) Rule {
	return &Capability{
		RuleBase:  newBaseFromLog(log),
		Qualifier: newQualifierFromLog(log),
		Names:     Must(toValues(CAPABILITY, "name", log["capname"])),
	}
}

func (r *Capability) Validate() error {
	if err := validateValues(r.Kind(), "name", r.Names); err != nil {
		return fmt.Errorf("%s: %w", r, err)
	}
	return nil
}

func (r *Capability) Compare(other Rule) int {
	o, _ := other.(*Capability)
	if res := compare(r.Names, o.Names); res != 0 {
		return res
	}
	return r.Qualifier.Compare(o.Qualifier)
}

func (r *Capability) String() string {
	return renderTemplate(r.Kind(), r)
}

func (r *Capability) Constraint() constraint {
	return blockKind
}

func (r *Capability) Kind() Kind {
	return CAPABILITY
}
