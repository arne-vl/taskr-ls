package lsp

type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label            string                     `json:"label"`
	LabelDetails     CompletionItemLabelDetails `json:"labelDetails"`
	InsertText       string                     `json:"insertText"`
	InsertTextFormat InsertTextFormat           `json:"insertTextFormat"`
	Kind             CompletionItemKind         `json:"kind"`
	Detail           string                     `json:"detail"`
	Documentation    string                     `json:"documentation"`
}

type CompletionItemKind int

const (
	CompletionItemText          CompletionItemKind = 1
	CompletionItemMethod        CompletionItemKind = 2
	CompletionItemFunction      CompletionItemKind = 3
	CompletionItemConstructor   CompletionItemKind = 4
	CompletionItemField         CompletionItemKind = 5
	CompletionItemVariable      CompletionItemKind = 6
	CompletionItemClass         CompletionItemKind = 7
	CompletionItemInterface     CompletionItemKind = 8
	CompletionItemModule        CompletionItemKind = 9
	CompletionItemProperty      CompletionItemKind = 10
	CompletionItemUnit          CompletionItemKind = 11
	CompletionItemValue         CompletionItemKind = 12
	CompletionItemEnum          CompletionItemKind = 13
	CompletionItemKeyword       CompletionItemKind = 14
	CompletionItemSnippet       CompletionItemKind = 15
	CompletionItemColor         CompletionItemKind = 16
	CompletionItemFile          CompletionItemKind = 17
	CompletionItemReference     CompletionItemKind = 18
	CompletionItemFolder        CompletionItemKind = 19
	CompletionItemEnumMember    CompletionItemKind = 20
	CompletionItemConstant      CompletionItemKind = 21
	CompletionItemStruct        CompletionItemKind = 22
	CompletionItemEvent         CompletionItemKind = 23
	CompletionItemOperator      CompletionItemKind = 24
	CompletionItemTypeParameter CompletionItemKind = 25
)

type InsertTextFormat int

const (
	InsertTextPlainText InsertTextFormat = 1
	InsertTextSnippet   InsertTextFormat = 2
)

type CompletionItemLabelDetails struct {
	Detail      string `json:"detail"`
	Description string `json:"description"`
}
