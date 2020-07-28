package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type JSONRPC struct {
	Id int `json:"id"`
}

func addContentLength(s string) string {
	return fmt.Sprintf("Content-Length:%d\r\n\r\n%s", len(s), s)
}

func initializeRequest(stdin io.WriteCloser) {
	initialize := `{"id": 0, "jsonrpc": "2.0", "method": "initialize", "params":{"capabilities":{"textDocument":{"codeAction":{"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.inline","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":false},"completion":{"completionItem":{"insertReplaceSupport":false,"snippetSupport":false}},"declaration":{"linkSupport":true},"definition":{"linkSupport":true},"hover":{},"implementation":{"linkSupport":true},"publishDiagnostics":{"relatedInformation":true},"semanticHighlightingCapabilities":{"semanticHighlighting":true},"signatureHelp":{"signatureInformation":{"parameterInformation":{"labelOffsetSupport":true}}},"typeDefinition":{"linkSupport":true}},"workspace":{"applyEdit":true,"didChangeWatchedFiles":{"dynamicRegistration":true}}},"clientInfo":{"name":"LanguageClient-neovim","version":"0.1.158 "},"processId":%d,"rootPath":"/home/kurenaif/Desktop/tools/gopls","rootUri":"file:///home/kurenaif/Desktop/tools/gopls","trace":"off"}}`
	initialize = fmt.Sprintf(initialize, os.Getpid())
	initialize = addContentLength(initialize)
	stdin.Write([]byte(initialize))
}

func initializedRequest(stdin io.WriteCloser) {
	s := `Content-Length:53
	
	{"jsonrpc":"2.0","method":"initialized","params":{}}`
	stdin.Write([]byte(s))
}

func codeLens(stdin io.WriteCloser) {
	s := `Content-Length:144
	
	{"jsonrpc":"2.0","method":"textDocument/codeLens","params":{"textDocument":{"uri":"file:///home/kurenaif/Desktop/tools/gopls/main.go"}},"id":1}`
	stdin.Write([]byte(s))
}

func documentSymbol(stdin io.WriteCloser) {
	s := `Content-Length:150
	
	{"jsonrpc":"2.0","method":"textDocument/documentSymbol","params":{"textDocument":{"uri":"file:///home/kurenaif/Desktop/tools/gopls/main.go"}},"id":2}`
	stdin.Write([]byte(s))
}

func workspaceSymbol(stdin io.WriteCloser) {
	s := `{"jsonrpc":"2.0","method":"workspace/symbol","params":{"query": "hello"},"id":2}`
	s = addContentLength(s)
	stdin.Write([]byte(s))
}

func didOpen(stdin io.WriteCloser) {
	s := `{"jsonrpc":"2.0","method":"textDocument/didOpen","params":{"textDocument":{"languageId":"go","text":"// Copyright 2019 The Go Authors. All rights reserved.\n// Use of this source code is governed by a BSD-style\n// license that can be found in the LICENSE file.\n\n// The gopls command is an LSP server for Go.\n// The Language Server Protocol allows any text editor\n// to be extended with IDE-like features;\n// see https://langserver.org/ for details.\n//\n// See https://github.com/golang/tools/tree/master/gopls\n// for the most up-to-date information on the gopls status.\npackage main // import \"golang.org/x/tools/gopls\"\n\nimport (\n\t\"context\"\n\t\"io\"\n\t\"log\"\n\t\"os\"\n\n\t\"golang.org/x/tools/gopls/internal/hooks\"\n\t\"golang.org/x/tools/internal/lsp/cmd\"\n\t\"golang.org/x/tools/internal/tool\"\n)\n\nfunc main() {\n\tctx := context.Background()\n\t// RDWRはreadとwrite。パーミッションで0666は読み書きができるユーザーその他。\n\n\tlogfile, _ := os.OpenFile(\"/home/kurenaif/Desktop/tools/gopls/test.log\", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)\n\tmultiLogFile := io.MultiWriter(logfile)\n\tlog.SetFlags(log.Ldate | log.Ltime | log.Llongfile)\n\tlog.SetOutput(multiLogFile)\n\n\t// io.MultiWriteで、\n\t// 標準出力とファイルの両方を束ねて、\n\t// logの出力先に設定する\n\tlog.SetOutput(io.MultiWriter(logfile, os.Stdout))\n\tlog.Println(\"LSP Start\")\n\ttool.Main(ctx, cmd.New(\"gopls\", \"otu\", nil, hooks.Options), os.Args[1:])\n\tlog.Println(\"LSP End\")\n}\n\nfunc test_function(){\n}\n","uri":"file:///home/kurenaif/Desktop/tools/gopls/main.go","version":0}}}`
	s = addContentLength(s)
	stdin.Write([]byte(s))
}

func didChangeWatchedFiles(stdin io.WriteCloser) {
	s := `{"jsonrpc":"2.0","error":{"code":-32603,"message":"entity not found"},"id":1}`
	s = addContentLength(s)
	stdin.Write([]byte(s))
}

func waitMessage(bufReader *bufio.Reader) JSONRPC {
	for {
		buf, isPrefix, err := bufReader.ReadLine()
		s := string(buf)
		if err != nil {
			log.Fatalf("%q", err)
		}
		if isPrefix {
			log.Fatalf("isPrefix が trueのときはまだ実装してないよ!")
		}

		// Headerを発見
		if strings.Contains(s, "Content-Length:") {
			sp := strings.Split(s, ":")
			if len(sp) < 2 {
				log.Fatalf("Content-Length has not (req %s):", s)
			}
			num, err := strconv.Atoi(strings.TrimSpace(string(sp[1])))
			if err != nil {
				log.Fatalf("Content-Length(%s) Prase Error: %q", s, err)
			}
			fmt.Println(num)

			// 空行でHeaderを区別する
			for {
				buf, isPrefix, err := bufReader.ReadLine()
				if err != nil {
					log.Fatalf("%q", err)
				}
				s := string(buf)
				if isPrefix {
					log.Fatalf("isPrefix が trueのときはまだ実装してないよ!")
				}
				if strings.TrimSpace(s) == "" {
					break
				}
			}
			// Header終わり

			buf := make([]byte, num)
			readNum, err := bufReader.Read(buf)
			if readNum != num {
				log.Fatal("読み込んだ文字数が一致してないよ!")
			}
			if err != nil {
				log.Fatal("%q", err)
			}
			fmt.Println(string(buf))

			var rpc JSONRPC
			if err := json.Unmarshal(buf, &rpc); err != nil {
				log.Fatal(err)
			}
			return rpc
		}
	}
}

func main() {

	// cmd := exec.Command("/home/kurenaif/Desktop/tools/gopls/gopls", "-logfile", "/home/kurenaif/Desktop/tools/gopls/rpc.log", "-rpc.trace", "-debug", "localhost:9090")
	cmd := exec.Command("/home/kurenaif/go/bin/gopls", "-logfile", "/home/kurenaif/Desktop/tools/gopls/rpc.log", "-rpc.trace", "-debug", "localhost:9090")
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}
	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}
	reader := bufio.NewReader(stdout)

	go func(reader io.Reader) {
		initializeRequest(stdin) // initialize
		bufReader := bufio.NewReader(reader)
		waitMessage(bufReader)    // initialize response
		initializedRequest(stdin) // initialized
		//		waitMessage(bufReader)    // initialize response
		//		waitMessage(bufReader)    // initialize response
		//		waitMessage(bufReader)    // initialize response
		waitMessage(bufReader) // initialize response
		didChangeWatchedFiles(stdin)
		didOpen(stdin)
		codeLens(stdin)
		documentSymbol(stdin)
		// workspaceSymbol(stdin)
		waitMessage(bufReader) // initialize response
		for {
			fmt.Println("waiting for messages...")
			buf := make([]byte, 256)
			bufReader.Read(buf)
			fmt.Print(string(buf))
		}
		// documentSymbol(stdin)
	}(reader)

	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()

}
