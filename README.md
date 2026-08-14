# 带超时的抓取器

CLI → Fetcher → HTTPX；父 context 取消应中断下游请求。

## 测试

```bash
set GOTOOLCHAIN=local
go test ./... -count=1
```
