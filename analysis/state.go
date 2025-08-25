package analysis

import (
	"taskr_ls/lsp"
	"taskr_ls/taskr"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	items := taskr.BaseCompletionItems()
	items = append(items, taskr.Snippets()...)

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}
