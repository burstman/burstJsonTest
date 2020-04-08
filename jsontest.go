package jsontest

const url = "https://wallet1.burst-team.us:2083/burst?requestType=getState"

// GetState import burst api getsate
type GetState struct {
	Application                    string `json:"application"`
	Version                        string `json:"version"`
	Time                           int    `json:"time"`
	LastBlock                      string `json:"lastBlock"`
	CumulativeDifficulty           string `json:"cumulativeDifficulty"`
	TotalEffectiveBalanceNXT       int64  `json:"totalEffectiveBalanceNXT"`
	NumbesrOfBlocks                int    `json:"numberOfBlocks"`
	NumberOfTransactions           int    `json:"numberOfTransactions"`
	NumberOfAccounts               int    `json:"numberOfAccounts"`
	NumberOfAssets                 int    `json:"numberOfAssets"`
	NumberOfOrders                 int    `json:"numberOfOrders"`
	NumberOfAskOrders              int    `json:"numberOfAskOrders"`
	NumberOfBidOrders              int    `json:"numberOfBidOrders"`
	NumberOfTrades                 int    `json:"numberOfTrades"`
	NumberOfTransfers              int    `json:"numberOfTransfers"`
	NumberOfAliases                int    `json:"numberOfAliases"`
	NumberOfPeers                  int    `json:"numberOfPeers"`
	NumberOfUnlockedAccounts       int    `json:"numberOfUnlockedAccounts"`
	LastBlockchainFeeder           string `json:"lastBlockchainFeeder"`
	LastBlockchainFeederHeight     int    `json:"lastBlockchainFeederHeight"`
	IsScanning                     bool   `json:"isScanning"`
	AvailableProcessors            int    `json:"availableProcessors"`
	MaxMemory                      int    `json:"maxMemory"`
	TotalMemory                    int    `json:"totalMemory"`
	FreeMemory                     int    `json:"freeMemory"`
	IndirectIncomingServiceEnabled bool   `json:"indirectIncomingServiceEnabled"`
	GrpcAPIEnabled                 bool   `json:"grpcApiEnabled"`
	GrpcAPIPort                    int    `json:"grpcApiPort"`
	RequestProcessingTime          int    `json:"requestProcessingTime"`
}
