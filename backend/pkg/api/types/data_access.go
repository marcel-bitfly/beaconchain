package types

import (
	"github.com/gobitfly/beaconchain/pkg/api/enums"
	"github.com/gobitfly/beaconchain/pkg/consapi/types"
	"github.com/shopspring/decimal"
)

// everything that goes in this file is for the data access layer only
// it won't be converted to typescript or used in the frontend

const DefaultGroupId = 0
const AllGroups = -1
const NetworkAverage = -2
const DefaultGroupName = "default"

type Sort[T enums.Enum] struct {
	Column T
	Desc   bool
}

type VDBIdPrimary int
type VDBIdPublic string
type VDBIdValidatorSet []VDBValidator
type VDBId struct {
	Validators      VDBIdValidatorSet // if this is nil, then use the id
	Id              VDBIdPrimary
	AggregateGroups bool
}

// could replace if we want the import in all files
type VDBValidator = types.ValidatorIndex

type DashboardInfo struct {
	Id     VDBIdPrimary `db:"id"` // this must be the bigint id
	UserId uint64       `db:"user_id"`
}

type CursorLike interface {
	IsCursor() bool
	IsValid() bool
	IsReverse() bool
}

type GenericCursor struct {
	Reverse bool `json:"r"`
	Valid   bool `json:"-"`
}

func (c GenericCursor) IsCursor() bool {
	return true
}

func (c GenericCursor) IsValid() bool {
	return c.Valid
}

// note: dont have to check for valid when calling this
func (c GenericCursor) IsReverse() bool {
	return c.Reverse && c.Valid
}

type CLDepositsCursor struct {
	GenericCursor
	Slot      int64
	SlotIndex int64
}

type ELDepositsCursor struct {
	GenericCursor
	BlockNumber int64
	LogIndex    int64
}

type ValidatorsCursor struct {
	GenericCursor

	Index uint64 `json:"vi"`
}

type RewardsCursor struct {
	GenericCursor

	Epoch   uint64
	GroupId int64
}

type ValidatorDutiesCursor struct {
	GenericCursor

	Index  uint64
	Reward decimal.Decimal
}

type WithdrawalsCursor struct {
	GenericCursor

	Slot            uint64
	WithdrawalIndex uint64
	Index           uint64
	Recipient       []byte
	Amount          uint64
}

type UserCredentialInfo struct {
	Id             uint64 `db:"id"`
	Email          string `db:"email"`
	EmailConfirmed bool   `db:"email_confirmed"`
	Password       string `db:"password"`
	ProductId      string `db:"product_id"`
	UserGroup      string `db:"user_group"`
}

type BlocksCursor struct {
	GenericCursor
	Slot uint64 // basically the same as Block, Epoch, Age; mandatory, used to index

	// optional, max one of those (for now)
	Proposer uint64
	Group    uint64
	Status   uint64
	Reward   decimal.Decimal
}

type NetworkInfo struct {
	ChainId uint64
	Name    string
}

// -------------------------
// validator indices structs, only used between data access and api layer

type VDBGeneralSummaryValidators struct {
	// fill slices with indices of validators
	Deposited   []uint64
	Pending     []IndexTimestamp
	Online      []uint64
	Offline     []uint64
	Slashing    []uint64
	Exiting     []IndexTimestamp
	Slashed     []uint64
	Exited      []uint64
	Withdrawing []IndexTimestamp
	Withdrawn   []uint64
}

type IndexTimestamp struct {
	Index     uint64
	Timestamp uint64
}

type VDBSyncSummaryValidators struct {
	// fill slices with indices of validators
	Upcoming []uint64
	Current  []uint64
	Past     []struct {
		Index uint64
		Count uint64
	}
}

type VDBSlashingsSummaryValidators struct {
	// fill with the validator index that got slashed and the index of the validator that slashed it
	GotSlashed []struct {
		Index     uint64
		SlashedBy uint64
	}
	// fill with the validator index that slashed and the index of the validators that got slashed
	HasSlashed []struct {
		Index          uint64
		SlashedIndices []uint64
	}
}

type VDBProposalSummaryValidators struct {
	Proposed []struct {
		Index          uint64
		ProposedBlocks []uint64
	}
	Missed []struct {
		Index        uint64
		MissedBlocks []uint64
	}
}
