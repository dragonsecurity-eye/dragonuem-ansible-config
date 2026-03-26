# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Go library that programmatically generates Ansible playbook YAML structures. Part of the DragonEye/DragonUEM platform (`eye.dragonsecurity.io/dragonuem-ansible-config`). This is a library package — it has no `main` package or binary output.

## Build & Test Commands

```bash
go build ./...       # Build all packages
go test ./...        # Run all tests
go vet ./...         # Static analysis
```

No tests exist yet. When adding tests, place them in the `ansible` package alongside the source files.

## Architecture

All code lives in the `ansible/` package. The library provides a builder pattern for constructing Ansible playbook structures that serialize to YAML.

- **types.go** — All struct definitions with `yaml` struct tags. Covers: playbooks, builtin modules (group, user, shell), and community modules (flatpak, homebrew/homebrew_cask). Uses `Present`/`Absent` constants for state fields.
- **ansible.go** — `NewAnsiblePlaybook()` constructor (defaults: localhost, local connection, no fact gathering) and `AddAnsibleTask()` to append tasks.
- **Module files** — Each file exposes Install/Uninstall (or Add/Remove) functions that return typed task structs:
  - `flatpak.go` — Flatpak packages
  - `homebrew.go` — Homebrew formulae and casks (install/upgrade/uninstall for both)
  - `local_group.go` — Local OS groups
  - `local_user.go` — Local OS users (linux vs macos password hashing distinction)
  - `shell.go` — Shell script execution

**Key pattern:** Each builder function validates required fields, sets the appropriate state, and returns `(*TypedStruct, error)`. Tasks are added to a playbook via `AddAnsibleTask(task any)` which accepts any task type.