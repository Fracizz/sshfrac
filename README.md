# sshctl

High-performance cross-platform SSH/SCP CLI with an encrypted local server inventory.

## Features

- Single static binaries for Linux / Windows / macOS (amd64 + arm64)
- Remote command execution and interactive shells
- Fast file transfer via SFTP (`scp` subcommand, directories supported)
- AES-256-GCM encrypted passwords at rest (`enc:v1:`)
- Case-insensitive contains search on name / IP / description
- Host key verification via `~/.ssh/known_hosts` (optional `--insecure` for labs)

## Install

```bash
go install github.com/Fracizz/sshctl@latest
```

Or download a release binary / build from source:

```bash
git clone https://github.com/Fracizz/sshctl.git
cd sshctl
go build -o bin/sshctl .
```

## Quick start

```bash
# Create ~/.sshctl/servers.json (placeholder only — no real secrets)
sshctl init

# Add a host (password is encrypted on save)
sshctl add --name lab --host 192.0.2.10 --user root --password 'secret' --desc "lab box"

# Or key-based auth
sshctl add --name prod --host 192.0.2.11 --user root --key ~/.ssh/id_ed25519 --desc "prod"

sshctl list
sshctl search -s 192.0.2
sshctl exec lab -- hostname
sshctl shell lab
sshctl scp ./app.tar.gz lab:/tmp/app.tar.gz
```

Config path priority: `--config` > `$SSHCTL_CONFIG` > `~/.sshctl/servers.json`

First SSH to a host should populate OpenSSH `known_hosts` (or pass `--insecure` only in trusted labs).

## JSON format

```json
{
  "servers": [
    {
      "name": "example-host",
      "description": "demo entry",
      "host": "192.0.2.10",
      "port": 22,
      "user": "root",
      "password": "enc:v1:...",
      "os": "Linux"
    }
  ]
}
```

See `configs/servers.example.json` for placeholders. Never commit real passwords.

## Security

Read [SECURITY.md](SECURITY.md). Ciphertext is bound to the local machine user material; moving the file to another PC requires re-entering passwords.

## License

MIT — see [LICENSE](LICENSE).
