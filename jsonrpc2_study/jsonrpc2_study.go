package main

import (
	"context"
	"os"

	"golang.org/x/tools/internal/fakenet"
	"golang.org/x/tools/internal/jsonrpc2"
	"golang.org/x/tools/internal/lsp/cache"
	"golang.org/x/tools/internal/lsp/lsprpc"
)

func main() {
	ctx := context.Background()
	stream := jsonrpc2.NewHeaderStream(fakenet.NewConn("stdio", os.Stdin, os.Stdout))
	ss := lsprpc.NewStreamServer(cache.New(ctx, s.app.options), isDaemon)
	stream.Read()
}

// fakenet: これはstdioをReadWriteできるだけのnet.Conn
// conn := fakenet.NewConn("stdio", os.Stdin, os.Stdout)
