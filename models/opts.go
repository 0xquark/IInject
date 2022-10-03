package models

type CmdLineOpts struct {
	Input    string
	Output   string
	Meta     bool
	Suppress bool
	Offset   string
	Inject   bool
	Payload  string
	Type     string
	Encode   bool
	Decode   bool
	Key      string
}
