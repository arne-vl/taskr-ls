package taskr

import "taskr_ls/lsp"

func BaseCompletionItems() []lsp.CompletionItem {
	return []lsp.CompletionItem{
		{
			Label:            "task",
			InsertText:       "task ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemKeyword,
			Detail:           "task",
			Documentation: `Defines a **task block** which can be executed.

A task can have following keys:
- run   = command to execute (required)
- desc  = description
- needs = dependencies
- alias = aliases

Example:
` + "```taskrfile\n" + `task build:
  run   = go build
  desc  = Compile the project
  alias = b
` + "```",
		},
		{
			Label:            "env",
			InsertText:       "env ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemKeyword,
			Detail:           "env",
			Documentation: `Defines an **environment block** which can be loaded.

An environment needs the following key:
- file = file name

Example:
` + "```taskrfile\n" + `env dev:
  file = .env
` + "```",
		},
		{
			Label:            "default",
			InsertText:       "default ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemKeyword,
			Detail:           "default",
			Documentation: `Specifies the **default environment**.

The default environment is used when no environment is passed in the` + " **taskr** " + `command.
Only one default environment is possible per` + " **taskrfile**." + `

Example:
` + "```taskrfile\n" + `default env dev:
  file = .env
` + "```",
		},
		{
			Label:            "file",
			InsertText:       "file = ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemProperty,
			Detail:           "env:file",
			Documentation:    "Defines the **file** associated with the environment.",
		},
		{
			Label:            "run",
			InsertText:       "run = ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemProperty,
			Detail:           "task:run",
			Documentation:    "Defines the **command** that will be executed when the task runs.",
		},
		{
			Label:            "desc",
			InsertText:       "desc = ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemProperty,
			Detail:           "task:desc",
			Documentation:    "Provides a **human-readable description** of the task.",
		},
		{
			Label:            "needs",
			InsertText:       "needs = ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemProperty,
			Detail:           "task:needs",
			Documentation: `List of tasks that must **run before this task**.

Comma separated list of task names or aliases. Will be executed in the defined order.`,
		},
		{
			Label:            "alias",
			InsertText:       "alias = ",
			InsertTextFormat: lsp.InsertTextPlainText,
			Kind:             lsp.CompletionItemProperty,
			Detail:           "task:alias",
			Documentation: `List of **alternative names** that can be used to run this task.

Comma separated list. Useful for quick shortcuts combined with having clear naming.`,
		},
	}
}
