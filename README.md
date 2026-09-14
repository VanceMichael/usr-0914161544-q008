# 冠名列车文旅权益核销总线

项目接收旅程相关事件并维护车票与两地文旅权益的核销投影。`internal/domain` 保存事件契约，Gin 提供 HTTP 边界，NATS JetStream 用于消息传递，PostgreSQL 保存事件、检查点和投影。

执行 `docker compose up --build` 启动完整环境，`GET /health` 用于探活。领域测试可使用 `go test ./...` 运行。
