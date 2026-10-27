# Go Context: timeouts, cancellation and graceful shutdown

Code for the video [Golang Context Explained: Timeouts, Cancellation and Graceful Shutdown](https://youtu.be/fCYZprTvNVg).

The search server answers after a delay set in the URL (`/search?mirror=N&ms=D`) and logs when a client gives up first.

```bash
go run ./server &        # localhost:9091
```

| Step | Run | What it shows |
|---|---|---|
| no time limit | `go run ./hang` | plain `http.Get` waits for a 10-minute answer (stop it with Ctrl+C) |
| `context.WithTimeout` | `go run ./timeout` | `context deadline exceeded` after 2 s; the server log says the client is gone |
| `context.WithCancel` | `go run ./mirrors` | three mirrors, the fastest wins, the other two stop at ~300 ms |
| your own loops | `go run ./loop` | stops after 5 of 20 images; `errors.Is(err, context.DeadlineExceeded)` is true |
| `context.WithCancelCause` | `go run ./cause` | `ctx.Err()` says canceled, `context.Cause(ctx)` says why |
| graceful shutdown | `go run ./shutdown`, then `curl localhost:9092/report` and Ctrl+C in the server | `signal.NotifyContext` + `srv.Shutdown`: the open request still finishes |
| `context.WithValue` | `go run ./reqid`, then `curl "localhost:9190/search?q=go"` | a request ID from middleware in every log line |
| lost cancel | `go vet ./vetcheck` | vet reports a discarded cancel function |

`./vetcheck` is meant to fail `go vet`; it shows the warning.
