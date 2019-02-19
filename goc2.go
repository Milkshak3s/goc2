package goc2

type Server interface {
	Listen()				error
	StageCommands(string)	error
	GetResults()			([]string, error)
}

type Client interface {
	GetCommands()			([]string, error)
	SendResults([]string)	error
}
