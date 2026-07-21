# Security Policy

## Supported versions

Security fixes are applied to the latest `main` branch release.

## Reporting a vulnerability

Please open a private security advisory on GitHub, or email the maintainer via the profile on https://github.com/Fracizz.

Do **not** open a public issue for secrets exposure or remote code execution.

## Threat model (short)

- Inventory passwords are encrypted at rest with AES-256-GCM using a **machine-derived** key. Anyone who can run `sshctl` as the same OS user can decrypt them.
- SSH host keys are verified against `~/.ssh/known_hosts` by default. `--insecure` disables this and is for labs only.
- Private key files referenced by `key_file` are read from disk as-is; protect them with OS permissions (and preferably a passphrase).
