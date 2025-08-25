package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	METHOD string `json:"method"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	METHOD string `json:"method"`
}
