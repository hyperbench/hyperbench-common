package common

// Blockchain define the service need provided in blockchain.
type Blockchain interface {

	// DeployContract should deploy contract with config file
	DeployContract() error

	// Invoke just invoke the contract
	Invoke(Invoke, ...Option) *Result

	// Transfer a amount of money from a account to the other one
	Transfer(Transfer, ...Option) *Result

	// Confirm check the result of `Invoke` or `Transfer`
	Confirm(*Result, ...Option) *Result

	// Query do some query
	Query(Query, ...Option) interface{}

	// Option pass the options to affect the action of client
	Option(Option) error

	// GetContext Generate TxContext based on New/Init/DeployContract
	// GetContext will only be run in master
	// return the information how to invoke the contract, maybe include
	// contract address, abi or so.
	// the return value will be send to worker to tell them how to invoke the contract
	GetContext() (string, error)

	// SetContext set test context into go client
	// SetContext will be run once per worker
	SetContext(ctx string) error

	// ResetContext reset test group context in go client
	ResetContext() error

	// Statistic query the statistics information in the time interval defined by
	// nanosecond-level timestamps `from` and `to`
	Statistic(statistic Statistic) (*RemoteStatistic, error)

	// LogStatus records blockheight and time
	LogStatus() (*ChainInfo, error)
}
