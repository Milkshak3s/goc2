package goc2

type Server interface {
	Listen()				int
	StageCommands(string)	int
	GetResults()			([]string, int)
}

type Client interface {
	GetCommands()			([]string, int)
	SendResults([]string)	int
}