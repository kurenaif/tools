S = '{"jsonrpc": "2.0", "method": "initialize", "params":{"capabilities":{"textDocument":{"codeAction":{"codeActionLiteralSupport":{"codeActionKind":{"valueSet":["quickfix","refactor","refactor.extract","refactor.inline","refactor.rewrite","source","source.organizeImports"]}}},"codeLens":{"dynamicRegistration":true},"colorProvider":{"dynamicRegistration":false},"completion":{"completionItem":{"insertReplaceSupport":false,"snippetSupport":false}},"declaration":{"linkSupport":true},"definition":{"linkSupport":true},"hover":{},"implementation":{"linkSupport":true},"publishDiagnostics":{"relatedInformation":true},"semanticHighlightingCapabilities":{"semanticHighlighting":true},"signatureHelp":{"signatureInformation":{"parameterInformation":{"labelOffsetSupport":true}}},"typeDefinition":{"linkSupport":true}},"workspace":{"applyEdit":true,"didChangeWatchedFiles":{"dynamicRegistration":true}}},"clientInfo":{"name":"LanguageClient-neovim","version":"0.1.158 "},"processId":null,"rootPath":"/home/kurenaif/Desktop/tools/gopls","rootUri":"file:///home/kurenaif/Desktop/tools/gopls","trace":"off"}'
S1 = "Content-Length: {}\r\n\r\n{}".format(len(S), S)
print(S1)