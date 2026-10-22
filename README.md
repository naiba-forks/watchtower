# Watchtower 镜像说明

本仓库仅用于说明镜像来源与维护位置，不存放 Watchtower 源码，也不运行独立构建工作流。

- **统一维护仓库：** [naiba/Dockerfiles](https://github.com/naiba/Dockerfiles)
- **上游：** [nicholas-fedor/watchtower](https://github.com/nicholas-fedor/watchtower)
- **计划镜像地址：** `ghcr.io/naiba-forks/watchtower`
- **计划架构：** `linux/amd64`、`linux/arm64`

## 当前状态

首次基线静态审查进行中，尚未批准或发布本仓库对应的镜像。候选版本为上游 `v1.23.1`，完整提交为 `520f95dea6e3c99b3e0e656e4a2950af11d6397b`。

审查仅针对恶意行为与依赖来源，不是漏洞、正确性或安全保证。通过后，固定提交构建、审查记录和后续更新均由 `naiba/Dockerfiles` 统一维护；镜像是否已发布及对应摘要，以该仓库的记录为准。

## 与历史 fork 的关系

这是说明仓库，不是历史 fork 的恢复。不恢复或承诺历史 fork 的额外功能，包括 `WATCHTOWER_DISABLE_IMAGES`；使用前请按所选上游版本的文档核对配置。

Watchtower 能通过 Docker socket 更新容器，并可执行用户配置的生命周期钩子和通知。镜像发布不代表已经获准修改、部署或重启任何服务器。
