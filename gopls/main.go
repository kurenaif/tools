// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gopls command is an LSP server for Go.
// The Language Server Protocol allows any text editor
// to be extended with IDE-like features;
// see https://langserver.org/ for details.
//
// See https://github.com/golang/tools/tree/master/gopls
// for the most up-to-date information on the gopls status.
package main // import "golang.org/x/tools/gopls"

import (
	"context"
	"io"
	"log"
	"os"

	"golang.org/x/tools/gopls/internal/hooks"
	"golang.org/x/tools/internal/lsp/cmd"
	"golang.org/x/tools/internal/tool"
)

func main() {
	ctx := context.Background()
	// RDWRはreadとwrite。パーミッションで0666は読み書きができるユーザーその他。

	logfile, _ := os.OpenFile("/home/kurenaif/Desktop/tools/gopls/test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)

	// io.MultiWriteで、
	// 標準出力とファイルの両方を束ねて、
	// logの出力先に設定する
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.Println("LSP Start")
	tool.Main(ctx, cmd.New("gopls", "otu", nil, hooks.Options), os.Args[1:])
	log.Println("LSP End")
}

func test_function(){
}
