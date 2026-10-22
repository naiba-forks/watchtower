# Watchtower 镜像说明

本仓库仅用于说明镜像来源与维护位置，不存放 Watchtower 源码，也不运行独立构建工作流。

- **统一维护仓库：** [naiba/Dockerfiles](https://github.com/naiba/Dockerfiles)
- **上游：** [nicholas-fedor/watchtower](https://github.com/nicholas-fedor/watchtower)
- **镜像地址：** `ghcr.io/naiba-forks/watchtower`
- **架构：** `linux/amd64`、`linux/arm64`

## 当前状态

首次基线静态恶意行为审查已通过，固定上游 `v1.23.1` / `520f95dea6e3c99b3e0e656e4a2950af11d6397b`。批准范围仅限中央固定构建与 scratch 运行镜像，不是漏洞、功能或历史 fork 兼容性保证。

- 中央构建提交：[`f353ab7`](https://github.com/naiba/Dockerfiles/commit/f353ab7)
- 发布 CI：[39740434494](https://github.com/naiba/Dockerfiles/actions/runs/39740434494)，amd64、arm64 及 publish 全部成功。
- 已验证标签：`latest`、`v1.23.1`、`v1.23.1-520f95d`。
- 三个标签的双架构清单摘要：`sha256:ef913ef732cc4755e5ceeb19439bbd619c5f0ff14cf279b015cb9c8637fbe395`。
- **包为 public，已验证匿名拉取可用。**
- [构建与审查记录](https://github.com/naiba/Dockerfiles/tree/master/watchtower)

镜像只含静态程序、固定 builder 的 CA 和内嵌时区数据，不包含 shell、Docker CLI 或凭据辅助程序。中央工作流没有修改、部署或重启服务器。

## 与历史 fork 的关系

这是说明仓库，不是历史 fork 的恢复。不恢复或承诺历史 fork 的额外功能，包括 `WATCHTOWER_DISABLE_IMAGES`；使用前请按所选上游版本的文档核对配置。

Watchtower 能通过 Docker socket 更新容器，并可执行用户配置的生命周期钩子和通知。镜像发布不代表已经获准修改、部署或重启任何服务器。
