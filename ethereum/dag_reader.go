package ethereum

// DAGReader provides a way for smartpool to retrieve DAG dataset. How the DAG
// is retrieve is upto structs implementing the interface.
type DAGReader interface {
	// NextWord return next data chunk of the DAG dataset. First 8 bytes must
	// be ignored.
	NextWord() ([]byte, error)
}
