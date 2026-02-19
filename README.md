# Watchtower

A lightweight fork of [containrrr/watchtower](https://github.com/containrrr/watchtower) for automating Docker container base image updates.

## Quick Start

```bash
docker run --detach \
    --name watchtower \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    ghcr.io/naiba-forks/watchtower
```

## Features

- **Auto Update** — Detects new images and automatically restarts containers
- **Notifications** — Via [Shoutrrr](https://github.com/containrrr/shoutrrr) (Telegram, Discord, etc.)
- **Filtering** — Whitelist/blacklist by container name, label, or scope
- **Lifecycle Hooks** — Execute commands before/after container updates
- **Scheduling** — Cron expressions or polling intervals

## Configuration

### Environment Variables

| Variable | Description |
|---|---|
| `WATCHTOWER_POLL_INTERVAL` | Poll interval in seconds (default: 86400) |
| `WATCHTOWER_SCHEDULE` | Cron expression for update schedule |
| `WATCHTOWER_CLEANUP` | Remove old images after update |
| `WATCHTOWER_LABEL_ENABLE` | Only update containers with enable label |
| `WATCHTOWER_DISABLE_CONTAINERS` | Comma-separated list of containers to exclude |
| `WATCHTOWER_DISABLE_IMAGES` | Comma-separated list of image names to exclude |
| `WATCHTOWER_SCOPE` | Monitoring scope for multi-instance setups |
| `WATCHTOWER_NOTIFICATION_URL` | Shoutrrr notification URL |
| `WATCHTOWER_LIFECYCLE_HOOKS` | Enable pre/post update lifecycle hooks |
| `WATCHTOWER_MONITOR_ONLY` | Only monitor, don't update |
| `WATCHTOWER_NO_PULL` | Don't pull new images |
| `WATCHTOWER_NO_RESTART` | Don't restart containers |
| `WATCHTOWER_RUN_ONCE` | Run once and exit |
| `WATCHTOWER_INCLUDE_STOPPED` | Include stopped containers |
| `WATCHTOWER_REVIVE_STOPPED` | Restart stopped containers after update |
| `WATCHTOWER_ROLLING_RESTART` | Restart containers one at a time |
| `WATCHTOWER_LABEL_TAKE_PRECEDENCE` | Container labels override arguments |
| `WATCHTOWER_TIMEOUT` | Timeout before force-stopping a container |

### Telegram Notification Example

```bash
docker run --detach \
    --name watchtower \
    --volume /var/run/docker.sock:/var/run/docker.sock \
    -e WATCHTOWER_NOTIFICATION_URL="telegram://token@telegram?chats=@channel" \
    -e WATCHTOWER_POLL_INTERVAL=3600 \
    -e WATCHTOWER_CLEANUP=true \
    ghcr.io/naiba-forks/watchtower
```

### Container Labels

| Label | Description |
|---|---|
| `com.centurylinklabs.watchtower.enable` | `true`/`false` to include/exclude |
| `com.centurylinklabs.watchtower.monitor-only` | `true` to only monitor |
| `com.centurylinklabs.watchtower.scope` | Scope for multi-instance setups |
| `com.centurylinklabs.watchtower.lifecycle.pre-check` | Command to run before checking |
| `com.centurylinklabs.watchtower.lifecycle.post-check` | Command to run after checking |
| `com.centurylinklabs.watchtower.lifecycle.pre-update` | Command to run before updating |
| `com.centurylinklabs.watchtower.lifecycle.post-update` | Command to run after updating |

## Changes from Upstream

This fork removes features not needed for simple homelab use:

- Removed HTTP API and Prometheus metrics
- Removed legacy notification backends (Email, Slack, MS Teams, Gotify)
- Removed documentation site, Grafana dashboards, GoReleaser config
- Migrated from `github.com/docker/docker` to `github.com/moby/moby` sub-modules
- Added `WATCHTOWER_DISABLE_IMAGES` for filtering by image name
- Kept: core auto-update, Shoutrrr notifications, filtering, lifecycle hooks

## License

[Apache-2.0](LICENSE.md)
