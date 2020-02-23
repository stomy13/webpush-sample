package json

type Endpoint struct {
	Endpoint string  `json:"endpoint"`
	Keys     CliKeys `json:"keys"`
}

type CliKeys struct {
	Key   string `json:"p256dh"`
	Token string `json:"auth"`
}
