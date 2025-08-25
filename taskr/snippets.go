package taskr

import "taskr_ls/lsp"

func Snippets() []lsp.CompletionItem {
	return []lsp.CompletionItem{
		// basic task
		{
			Label: "btsk",
			LabelDetails: lsp.CompletionItemLabelDetails{
				Detail: "basic task",
			},
			InsertText: `task ${1:name}:
  run   = ${2:command}
  desc  = ${3:description}
$0`,
			InsertTextFormat: lsp.InsertTextSnippet,
			Kind:             lsp.CompletionItemSnippet,
			Detail:           "snippet: btsk",
			Documentation: `Create a basic task like:
` + "```taskrfile" + `
task name:
  run  = command
  desc = description
`,
		},

		// basic env
		{
			Label: "benv",
			LabelDetails: lsp.CompletionItemLabelDetails{
				Detail: "env",
			},
			InsertText: `env ${1:name}:
  file = ${2:file}
$0`,
			InsertTextFormat: lsp.InsertTextSnippet,
			Kind:             lsp.CompletionItemSnippet,
			Detail:           "snippet: benv",
			Documentation: `Create an environment like:
` + "```taskrfile" + `
env name:
  file = file
`,
		},

		// default env
		{
			Label: "denv",
			LabelDetails: lsp.CompletionItemLabelDetails{
				Detail: "default env",
			},
			InsertText: `default env ${1:name}:
  file = ${2:file}
$0`,
			InsertTextFormat: lsp.InsertTextSnippet,
			Kind:             lsp.CompletionItemSnippet,
			Detail:           "snippet: denv",
			Documentation: `Create a default environment like:
` + "```taskrfile" + `
default env name:
  file = file
`,
		},
	}
}
