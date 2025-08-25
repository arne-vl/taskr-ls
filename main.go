package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"taskr_ls/analysis"
	"taskr_ls/lsp"
	"taskr_ls/rpc"
)

func main() {
	logger := getLogger(filepath.Join(os.Getenv("HOME"), ".local", "log", "taskr_ls", "log.txt"))
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()

		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
		}

		handleMessage(logger, writer, method, contents, state)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, method string, contents []byte, state analysis.State) {
	logger.Printf("Recieved message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not parse: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))

		logger.Println("Sent a reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Could not parse: %s", err)
		}
		logger.Printf("Opened: %s", request.Params.TextDocument.URI)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}

		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}

	case "textDocument/completion":
		var request lsp.CompletionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/completion: %s", err)
			return
		}

		response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)

		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))

}

func getLogger(filename string) *log.Logger {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic("Unable to create log directory: " + err.Error())
	}

	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Not a correct logfile")
	}

	return log.New(logfile, "[taskr_ls]", log.Ldate|log.Ltime|log.Lshortfile)
}
