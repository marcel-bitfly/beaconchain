package dataaccess

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
	"github.com/gobitfly/beaconchain/pkg/api/enums"
	t "github.com/gobitfly/beaconchain/pkg/api/types"
	"github.com/shopspring/decimal"
)

type DummyService struct {
}

// ensure DummyService pointer implements DataAccessor
var _ DataAccessor = (*DummyService)(nil)

func NewDummyService() *DummyService {
	// define custom tags for faker
	_ = faker.AddProvider("eth", func(v reflect.Value) (interface{}, error) {
		return randomEthDecimal(), nil
	})
	_ = faker.AddProvider("cl_el_eth", func(v reflect.Value) (interface{}, error) {
		return t.ClElValue[decimal.Decimal]{
			Cl: randomEthDecimal(),
			El: randomEthDecimal(),
		}, nil
	})
	return &DummyService{}
}

// generate random decimal.Decimal, should result in somewhere around 0.001 ETH (+/- a few decimal places) in Wei
func randomEthDecimal() decimal.Decimal {
	//nolint:gosec
	decimal, _ := decimal.NewFromString(fmt.Sprintf("%d00000000000", rand.Int63n(10000000)))
	return decimal
}

// must pass a pointer to the data
func commonFakeData(a interface{}) error {
	// TODO fake decimal.Decimal
	return faker.FakeData(a, options.WithRandomMapAndSliceMaxSize(5))
}

func (d *DummyService) Close() {
	// nothing to close
}

func (d *DummyService) GetLatestSlot() (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetLatestExchangeRates() ([]t.EthConversionRate, error) {
	r := []t.EthConversionRate{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetUserByEmail(ctx context.Context, email string) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) CreateUser(ctx context.Context, email, password string) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) RemoveUser(ctx context.Context, userId uint64) error {
	return nil
}

func (d *DummyService) UpdateUserEmail(ctx context.Context, userId uint64) error {
	return nil
}

func (d *DummyService) UpdateUserPassword(ctx context.Context, userId uint64, password string) error {
	return nil
}

func (d *DummyService) GetEmailConfirmationTime(ctx context.Context, userId uint64) (time.Time, error) {
	r := time.Time{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) UpdateEmailConfirmationTime(ctx context.Context, userId uint64) error {
	return nil
}

func (d *DummyService) GetEmailConfirmationHash(ctx context.Context, userId uint64) (string, error) {
	r := ""
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) UpdateEmailConfirmationHash(ctx context.Context, userId uint64, email, confirmationHash string) error {
	return nil
}

func (d *DummyService) GetUserInfo(ctx context.Context, userId uint64) (*t.UserInfo, error) {
	r := t.UserInfo{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetUserCredentialInfo(ctx context.Context, userId uint64) (*t.UserCredentialInfo, error) {
	r := t.UserCredentialInfo{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetUserIdByApiKey(ctx context.Context, apiKey string) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetUserIdByConfirmationHash(hash string) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetProductSummary(ctx context.Context) (*t.ProductSummary, error) {
	r := t.ProductSummary{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetFreeTierPerks(ctx context.Context) (*t.PremiumPerks, error) {
	r := t.PremiumPerks{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardInfo(ctx context.Context, dashboardId t.VDBIdPrimary) (*t.DashboardInfo, error) {
	r := t.DashboardInfo{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardInfoByPublicId(ctx context.Context, publicDashboardId t.VDBIdPublic) (*t.DashboardInfo, error) {
	r := t.DashboardInfo{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardName(ctx context.Context, dashboardId t.VDBIdPrimary) (string, error) {
	r := ""
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetValidatorsFromSlices(indices []uint64, publicKeys []string) ([]t.VDBValidator, error) {
	r := []t.VDBValidator{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetUserDashboards(ctx context.Context, userId uint64) (*t.UserDashboardsData, error) {
	r := t.UserDashboardsData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) CreateValidatorDashboard(ctx context.Context, userId uint64, name string, network uint64) (*t.VDBPostReturnData, error) {
	r := t.VDBPostReturnData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardOverview(ctx context.Context, dashboardId t.VDBId) (*t.VDBOverviewData, error) {
	r := t.VDBOverviewData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) RemoveValidatorDashboard(ctx context.Context, dashboardId t.VDBIdPrimary) error {
	return nil
}

func (d *DummyService) UpdateValidatorDashboardName(ctx context.Context, dashboardId t.VDBIdPrimary, name string) (*t.VDBPostReturnData, error) {
	r := t.VDBPostReturnData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) CreateValidatorDashboardGroup(ctx context.Context, dashboardId t.VDBIdPrimary, name string) (*t.VDBPostCreateGroupData, error) {
	r := t.VDBPostCreateGroupData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) UpdateValidatorDashboardGroup(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, name string) (*t.VDBPostCreateGroupData, error) {
	r := t.VDBPostCreateGroupData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) RemoveValidatorDashboardGroup(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64) error {
	return nil
}

func (d *DummyService) GetValidatorDashboardGroupExists(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64) (bool, error) {
	return true, nil
}

func (d *DummyService) GetValidatorDashboardExistingValidatorCount(ctx context.Context, dashboardId t.VDBIdPrimary, validators []t.VDBValidator) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) AddValidatorDashboardValidators(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, validators []t.VDBValidator) ([]t.VDBPostValidatorsData, error) {
	r := []t.VDBPostValidatorsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) AddValidatorDashboardValidatorsByDepositAddress(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, address string, limit uint64) ([]t.VDBPostValidatorsData, error) {
	r := []t.VDBPostValidatorsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) AddValidatorDashboardValidatorsByWithdrawalAddress(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, address string, limit uint64) ([]t.VDBPostValidatorsData, error) {
	r := []t.VDBPostValidatorsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) AddValidatorDashboardValidatorsByGraffiti(ctx context.Context, dashboardId t.VDBIdPrimary, groupId uint64, graffiti string, limit uint64) ([]t.VDBPostValidatorsData, error) {
	r := []t.VDBPostValidatorsData{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetValidatorDashboardValidators(ctx context.Context, dashboardId t.VDBId, groupId int64, cursor string, colSort t.Sort[enums.VDBManageValidatorsColumn], search string, limit uint64) ([]t.VDBManageValidatorsTableRow, *t.Paging, error) {
	r := []t.VDBManageValidatorsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) RemoveValidatorDashboardValidators(ctx context.Context, dashboardId t.VDBIdPrimary, validators []t.VDBValidator) error {
	return nil
}

func (d *DummyService) CreateValidatorDashboardPublicId(ctx context.Context, dashboardId t.VDBIdPrimary, name string, shareGroups bool) (*t.VDBPublicId, error) {
	r := t.VDBPublicId{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardPublicId(ctx context.Context, publicDashboardId t.VDBIdPublic) (*t.VDBPublicId, error) {
	r := t.VDBPublicId{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) UpdateValidatorDashboardPublicId(ctx context.Context, publicDashboardId t.VDBIdPublic, name string, shareGroups bool) (*t.VDBPublicId, error) {
	r := t.VDBPublicId{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) RemoveValidatorDashboardPublicId(ctx context.Context, publicDashboardId t.VDBIdPublic) error {
	return nil
}

func (d *DummyService) GetValidatorDashboardSlotViz(ctx context.Context, dashboardId t.VDBId, groupIds []uint64) ([]t.SlotVizEpoch, error) {
	r := struct {
		Epochs []t.SlotVizEpoch `faker:"slice_len=4"`
	}{}
	err := commonFakeData(&r)
	return r.Epochs, err
}

func (d *DummyService) GetValidatorDashboardSummary(ctx context.Context, dashboardId t.VDBId, period enums.TimePeriod, cursor string, colSort t.Sort[enums.VDBSummaryColumn], search string, limit uint64) ([]t.VDBSummaryTableRow, *t.Paging, error) {
	r := []t.VDBSummaryTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}
func (d *DummyService) GetValidatorDashboardGroupSummary(ctx context.Context, dashboardId t.VDBId, groupId int64, period enums.TimePeriod) (*t.VDBGroupSummaryData, error) {
	r := t.VDBGroupSummaryData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardSummaryChart(ctx context.Context, dashboardId t.VDBId, groupIds []int64, efficiency enums.VDBSummaryChartEfficiencyType, aggregation enums.ChartAggregation, afterTs uint64, beforeTs uint64) (*t.ChartData[int, float64], error) {
	r := t.ChartData[int, float64]{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardSummaryValidators(ctx context.Context, dashboardId t.VDBId, groupId int64) (*t.VDBGeneralSummaryValidators, error) {
	r := t.VDBGeneralSummaryValidators{}
	err := commonFakeData(&r)
	return &r, err
}
func (d *DummyService) GetValidatorDashboardSyncSummaryValidators(ctx context.Context, dashboardId t.VDBId, groupId int64, period enums.TimePeriod) (*t.VDBSyncSummaryValidators, error) {
	r := t.VDBSyncSummaryValidators{}
	err := commonFakeData(&r)
	return &r, err
}
func (d *DummyService) GetValidatorDashboardSlashingsSummaryValidators(ctx context.Context, dashboardId t.VDBId, groupId int64, period enums.TimePeriod) (*t.VDBSlashingsSummaryValidators, error) {
	r := t.VDBSlashingsSummaryValidators{}
	err := commonFakeData(&r)
	return &r, err
}
func (d *DummyService) GetValidatorDashboardProposalSummaryValidators(ctx context.Context, dashboardId t.VDBId, groupId int64, period enums.TimePeriod) (*t.VDBProposalSummaryValidators, error) {
	r := t.VDBProposalSummaryValidators{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardRewards(ctx context.Context, dashboardId t.VDBId, cursor string, colSort t.Sort[enums.VDBRewardsColumn], search string, limit uint64) ([]t.VDBRewardsTableRow, *t.Paging, error) {
	r := []t.VDBRewardsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) GetValidatorDashboardGroupRewards(ctx context.Context, dashboardId t.VDBId, groupId int64, epoch uint64) (*t.VDBGroupRewardsData, error) {
	r := t.VDBGroupRewardsData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardRewardsChart(ctx context.Context, dashboardId t.VDBId) (*t.ChartData[int, decimal.Decimal], error) {
	r := t.ChartData[int, decimal.Decimal]{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardDuties(ctx context.Context, dashboardId t.VDBId, epoch uint64, groupId int64, cursor string, colSort t.Sort[enums.VDBDutiesColumn], search string, limit uint64) ([]t.VDBEpochDutiesTableRow, *t.Paging, error) {
	r := []t.VDBEpochDutiesTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) GetValidatorDashboardBlocks(ctx context.Context, dashboardId t.VDBId, cursor string, colSort t.Sort[enums.VDBBlocksColumn], search string, limit uint64) ([]t.VDBBlocksTableRow, *t.Paging, error) {
	r := []t.VDBBlocksTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) GetValidatorDashboardEpochHeatmap(ctx context.Context, dashboardId t.VDBId) (*t.VDBHeatmap, error) {
	r := t.VDBHeatmap{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardDailyHeatmap(ctx context.Context, dashboardId t.VDBId, period enums.TimePeriod) (*t.VDBHeatmap, error) {
	r := t.VDBHeatmap{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardGroupEpochHeatmap(ctx context.Context, dashboardId t.VDBId, groupId uint64, epoch uint64) (*t.VDBHeatmapTooltipData, error) {
	r := t.VDBHeatmapTooltipData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardGroupDailyHeatmap(ctx context.Context, dashboardId t.VDBId, groupId uint64, day time.Time) (*t.VDBHeatmapTooltipData, error) {
	r := t.VDBHeatmapTooltipData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardElDeposits(ctx context.Context, dashboardId t.VDBId, cursor string, search string, limit uint64) ([]t.VDBExecutionDepositsTableRow, *t.Paging, error) {
	r := []t.VDBExecutionDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) GetValidatorDashboardClDeposits(ctx context.Context, dashboardId t.VDBId, cursor string, search string, limit uint64) ([]t.VDBConsensusDepositsTableRow, *t.Paging, error) {
	r := []t.VDBConsensusDepositsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) GetValidatorDashboardTotalElDeposits(ctx context.Context, dashboardId t.VDBId) (*t.VDBTotalExecutionDepositsData, error) {
	r := t.VDBTotalExecutionDepositsData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardTotalClDeposits(ctx context.Context, dashboardId t.VDBId) (*t.VDBTotalConsensusDepositsData, error) {
	r := t.VDBTotalConsensusDepositsData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetValidatorDashboardWithdrawals(ctx context.Context, dashboardId t.VDBId, cursor string, colSort t.Sort[enums.VDBWithdrawalsColumn], search string, limit uint64) ([]t.VDBWithdrawalsTableRow, *t.Paging, error) {
	r := []t.VDBWithdrawalsTableRow{}
	p := t.Paging{}
	_ = commonFakeData(&r)
	err := commonFakeData(&p)
	return r, &p, err
}

func (d *DummyService) GetValidatorDashboardTotalWithdrawals(ctx context.Context, dashboardId t.VDBId, search string) (*t.VDBTotalWithdrawalsData, error) {
	r := t.VDBTotalWithdrawalsData{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetAllNetworks() ([]t.NetworkInfo, error) {
	r := []t.NetworkInfo{}
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetSearchValidatorByIndex(ctx context.Context, chainId, index uint64) (*t.SearchValidator, error) {
	r := t.SearchValidator{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetSearchValidatorByPublicKey(ctx context.Context, chainId uint64, publicKey []byte) (*t.SearchValidator, error) {
	r := t.SearchValidator{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetSearchValidatorsByDepositAddress(ctx context.Context, chainId uint64, address []byte) (*t.SearchValidatorsByDepositAddress, error) {
	r := t.SearchValidatorsByDepositAddress{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetSearchValidatorsByDepositEnsName(ctx context.Context, chainId uint64, ensName string) (*t.SearchValidatorsByDepositEnsName, error) {
	r := t.SearchValidatorsByDepositEnsName{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetSearchValidatorsByWithdrawalCredential(ctx context.Context, chainId uint64, credential []byte) (*t.SearchValidatorsByWithdrwalCredential, error) {
	r := t.SearchValidatorsByWithdrwalCredential{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetSearchValidatorsByWithdrawalEnsName(ctx context.Context, chainId uint64, ensName string) (*t.SearchValidatorsByWithrawalEnsName, error) {
	r := t.SearchValidatorsByWithrawalEnsName{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetSearchValidatorsByGraffiti(ctx context.Context, chainId uint64, graffiti string) (*t.SearchValidatorsByGraffiti, error) {
	r := t.SearchValidatorsByGraffiti{}
	err := commonFakeData(&r)
	return &r, err
}

func (d *DummyService) GetUserValidatorDashboardCount(ctx context.Context, userId uint64) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetValidatorDashboardGroupCount(ctx context.Context, dashboardId t.VDBIdPrimary) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetValidatorDashboardValidatorsCount(ctx context.Context, dashboardId t.VDBIdPrimary) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}

func (d *DummyService) GetValidatorDashboardPublicIdCount(ctx context.Context, dashboardId t.VDBIdPrimary) (uint64, error) {
	r := uint64(0)
	err := commonFakeData(&r)
	return r, err
}
