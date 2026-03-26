# dragonuem-ansible-config

Go library for programmatically generating Ansible playbook YAML structures. Part of the DragonEye/DragonUEM platform.

## Installation

```bash
go get eye.dragonsecurity.io/dragonuem-ansible-config
```

## Usage

```go
package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "eye.dragonsecurity.io/dragonuem-ansible-config/ansible"
)

func main() {
    pb := ansible.NewAnsiblePlaybook()
    pb.Name = "Setup workstation"

    task, _ := ansible.InstallHomeBrewFormula("Install git", "git", "", true, false)
    pb.AddAnsibleTask(task)

    group, _ := ansible.AddLocalGroup("Create dev group", "developers", 1001, false, false)
    pb.AddAnsibleTask(group)

    playbooks := []ansible.AnsiblePlaybook{*pb}
    out, _ := yaml.Marshal(playbooks)
    fmt.Println(string(out))
}
```

## Supported Modules

| Module | Functions |
|--------|-----------|
| `ansible.builtin.group` | `AddLocalGroup`, `RemoveLocalGroup` |
| `ansible.builtin.user` | `AddLocalUser`, `RemoveLocalUser` |
| `ansible.builtin.shell` | `ExecuteScript` |
| `community.general.flatpak` | `InstallFlatpakPackage`, `UninstallFlatpakPackage` |
| `community.general.homebrew` | `InstallHomeBrewFormula`, `UpgradeHomeBrewFormula`, `UninstallHomeBrewFormula` |
| `community.general.homebrew_cask` | `InstallHomeBrewCask`, `UpgradeHomeBrewCask`, `UninstallHomeBrewCask` |

## License

See [LICENSE](LICENSE) for details.
