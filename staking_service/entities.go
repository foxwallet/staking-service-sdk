package staking_service

type Response[T any] struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type PageResult[T any] struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	List     []T   `json:"list"`
}

type User struct {
	ID       uint64 `json:"id"`
	BrokerID uint64 `json:"brokerID"`
	UID      string `json:"uid"`
}

type UserDetails struct {
	User
	DepositAddresses []string            `json:"depositAddresses"`
	Validators       ValidatorDetailList `json:"validators"`
}

type Validator struct {
	UID            string `json:"uid"`
	ValidatorIndex uint64 `json:"validatorIndex"`
	Pubkey         string `json:"pubkey"`
}

type ValidatorDetail struct {
	Validator
	Balance        string `json:"balance"`
	TotalRewards   string `json:"totalRewards"`
	Status         string `json:"status"`
	ActivationTime string `json:"ActivationTime"`
	ExitTime       string `json:"exitTime"`
	Principal      string `json:"principal"`
}

type ValidatorDetailList []ValidatorDetail
type ValidatorList []Validator

type CreateUserParams struct {
	UID              string   `json:"uid"`
	DepositAddresses []string `json:"depositAddresses"`
}

type AddDepositAddressRequest struct {
	DepositAddresses []string `json:"deposit_addresses"`
}

type Deposit struct {
	ID              uint64 `json:"id"`
	BrokerID        uint64 `json:"brokerId"`
	FromAddress     string `json:"fromAddress"`
	ContractAddress string `json:"contractAddress"`
	TxHash          string `json:"txHash"`
	ChainName       string `json:"chainName"`
	BlockHeight     uint64 `json:"blockHeight"`
	Slot            uint64 `json:"slot"`
	BlockTime       uint64 `json:"blockTime"`
}

type DepositDetails struct {
	Deposit
	Validators ValidatorDetailList `json:"validators"`
}

type UserValidators struct {
	UID        string        `json:"uid"`
	Validators ValidatorList `json:"validators"`
}

type ValidatorAssignmentParams struct {
	Assignments []Assignment `json:"assignments"`
}

type BatchValidatorAssignmentParams struct {
	Assignments []Assignment `json:"assignments"`
	TxHashs     []string     `json:"tx_hashs"`
}

type DepositDataRequestParams struct {
	Uid               string `json:"uid"`
	Quantity          int64  `json:"quantity"`
	WithdrawalAddress string `json:"withdrawalAddress"`
}

type SendTransactionRequestParams struct {
	SignedTransaction string `json:"signedTransaction"`
}

type SendTransactionResponse struct {
	TxHash string `json:"txHash"`
}

type Assignment struct {
	UID            string `json:"uid"`
	ValidatorCount int    `json:"validatorCount"`
}

type Reward struct {
	Date              string `json:"date"`
	TotalCount        int64  `json:"totalCount"`
	ActiveCount       int64  `json:"activeCount"`
	ExitCount         int64  `json:"exitCount"`
	Balance           string `json:"balance"`
	Principal         string `json:"principal"`
	Apr               string `json:"apr"`
	CumulativeRewards string `json:"cumulativeRewards"`
	CumulativeStaking string `json:"cumulativeStaking"`
	CumulativeGasFee  string `json:"cumulativeGasFee"`
	Rewards           string `json:"rewards"`
	Staking           string `json:"staking"`
	GasFee            string `json:"gasFee"`
	Withdrawal        string `json:"withdrawal"`
	Rate              string `json:"rate"`
}

type RewardList []Reward

type DepositDataResponse struct {
	Network  string            `json:"network"`
	Protocol string            `json:"protocol"`
	Period   string            `json:"period"`
	Ethereum *EthereumResponse `json:"ethereum"`
}

type EthereumResponse struct {
	ContractAddress     string             `json:"contractAddress"`
	EstimatedGas        uint64             `json:"estimatedGas"`
	UnsignedTransaction string             `json:"unsignedTransaction"`
	DepositData         []*DepositDataInfo `json:"depositData"`
}

type DepositDataInfo struct {
	DepositDataRoot      string `json:"depositDataRoot"`
	Pubkey               string `json:"pubkey"`
	Signature            string `json:"signature"`
	WithdrawalCredential string `json:"withdrawalCredential"`
}

type PoolRewardInfo struct {
	AvgYearApr       string `json:"avg_year_apr"`
	BeginSlot        int    `json:"begin_slot"`
	CumulativeReward string `json:"cumulative_reward"`
	DailyReward      string `json:"daily_reward"`
	Date             string `json:"date"`
	EndSlot          int    `json:"end_slot"`
}

type PoolRewardInfoList []PoolRewardInfo

type PooledStakingInfo struct {
	Principal        string `json:"principal"`
	RealizedReward   string `json:"realized_reward"`
	UnrealizedReward string `json:"unrealized_reward"`
	Point            string `json:"point"`
	UserAddress      string `json:"user_address"`
}

type PooledStakingInfoList []PooledStakingInfo

type UserPooledStakingInfo struct {
	DepositedEth string `json:"deposited_eth"`
	TotalEth     string `json:"total_eth"`
}

type PooledWithdrawRequestInfo struct {
	RequestID      int    `json:"request_id"`
	User           string `json:"user"`
	Point          string `json:"point"`
	Amount         string `json:"amount"`
	Principal      string `json:"principal"`
	WithdrawTxhash string `json:"withdraw_txhash"`
	Description    string `json:"description"`
}

type WithdrawPossibleBlock struct {
	CanWithdrawBlock string `json:"can_withdraw_block"`
}

type TransactionList []*TransactionInfo

type TransactionInfo struct {
	Amount         float32 `json:"amount"`
	Type           uint64  `json:"type"`
	TxTime         uint64  `json:"txTime"`
	Status         string  `json:"status"`
	TxHash         string  `json:"txHash"`
	WithdrawalHash string  `json:"withdrawalHash"`
}
