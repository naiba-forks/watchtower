# Watchtower 镜像说明

本仓库仅用于说明镜像来源与维护位置，不存放 Watchtower 源码，也不运行独立构建工作流。

- **统一维护仓库：** [naiba/Dockerfiles](https://github.com/naiba/Dockerfiles)
- **上游：** [nicholas-fedor/watchtower](https://github.com/nicholas-fedor/watchtower)
- **镜像地址：** `ghcr.io/naiba-forks/watchtower`
- **架构：** `linux/amd64`、`linux/arm64`

## 当前状态

当前发布为 `v1.23.2`，固定上游 `cd5946be205ea0a60dd31235a11f04ac9851ebf6`。增量静态恶意行为审查仅在既有中央源码白名单与 scratch 交付边界内通过；整版上游及无限制 CI/开发/测试工具链仍为暂缓，不是漏洞、功能或历史 fork 兼容性保证。

- 中央构建提交：[`de135b4`](https://github.com/naiba/Dockerfiles/commit/de135b48c55a0beea6de09ceac53a8018900efd1)。
- 发布 CI：[39879617105](https://github.com/naiba/Dockerfiles/actions/runs/39879617105)，amd64、arm64 及 publish 全部 `completed/success`。
- 已匿名验证标签：`latest`、`v1.23.2`、`v1.23.2-cd5946b`。
- 三个标签的双架构清单摘要：`sha256:332cd48ea2dba05d99f5359aaa2abffdeb7fb84969b26442eb290fe669d0d41a`。
- `linux/amd64` 子清单：`sha256:c95523b2cd402030a04238c489660d96df994332955460255b4a034010cf972d`。
- `linux/arm64` 子清单：`sha256:a2b1f3f590b46b6d1c39f9be66841599071760f316aeea228c570f4bb490093e`。
- **镜像公开可访问；已独立匿名读取三个标签及两个子镜像的 manifest/config，校验摘要、真实架构、版本、上游源码与中央构建提交。** 验证时间：`2026-10-23T00:44:25Z`；未下载镜像层或运行镜像。
- [构建与审查记录](https://github.com/naiba/Dockerfiles/tree/de135b48c55a0beea6de09ceac53a8018900efd1/watchtower)，[增量证据](https://github.com/naiba/Dockerfiles/blob/de135b48c55a0beea6de09ceac53a8018900efd1/watchtower/incremental-review-evidence.json)。

覆盖 474 个项目差异、23 个模块升级、1501 个直接输入分片；1381 项上下文请求中 1287 项关闭，94 项按基线排除且仍未核验，中央材料阻断为 0。排除项包括不由中央入口执行的上游工作流、开发/测试/安装工具及未核验外部制品，不代表这些内容获批；完整范围见审查记录，原始基线证据保持保留。

Go 1.27.3 原生 linux/amd64 builder 固定为 `sha256:8aa2821993207ae840ce09e4223217efcd14e01f71525db734faf4232bd59999`，源码白名单为 `go.mod go.sum main.go cmd/ internal/ pkg/`。scratch 镜像只含静态程序、同一 builder 的 CA 和内嵌时区，不含 shell、Docker CLI 或凭据辅助程序。

中央工作流没有修改、部署或重启服务器。发布 `latest` 后，已有自动更新服务可能独立拉取新版本。

## 与历史 fork 的关系

这是说明仓库，不是历史 fork 的恢复。不恢复或承诺历史 fork 的额外功能，包括 `WATCHTOWER_DISABLE_IMAGES`；使用前请按所选上游版本的文档核对配置。

Watchtower 能通过 Docker socket 更新容器，并可执行用户配置的生命周期钩子和通知。镜像发布不代表已经获准修改、部署或重启任何服务器。
