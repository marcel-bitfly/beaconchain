package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gobitfly/beaconchain/pkg/api/enums"
	types "github.com/gobitfly/beaconchain/pkg/api/types"

	"github.com/gorilla/mux"
)

// --------------------------------------
// Premium Plans

func (h *HandlerService) InternalGetProductSummary(w http.ResponseWriter, r *http.Request) {
	data, err := h.dai.GetProductSummary(r.Context())
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetProductSummaryResponse{
		Data: *data,
	}
	returnOk(w, response)
}

// --------------------------------------
// Latest State

func (h *HandlerService) InternalGetLatestState(w http.ResponseWriter, r *http.Request) {
	latestSlot, err := h.dai.GetLatestSlot()
	if err != nil {
		handleErr(w, err)
		return
	}

	exchangeRates, err := h.dai.GetLatestExchangeRates()
	if err != nil {
		handleErr(w, err)
		return
	}
	data := types.LatestStateData{
		LatestSlot:    latestSlot,
		ExchangeRates: exchangeRates,
	}

	response := types.InternalGetLatestStateResponse{
		Data: data,
	}
	returnOk(w, response)
}

// All handler function names must include the HTTP method and the path they handle
// Internal handlers may only be authenticated by an OAuth token

// --------------------------------------
// Ad Configurations

func (h *HandlerService) InternalPostAdConfigurations(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h *HandlerService) InternalGetAdConfigurations(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalPutAdConfiguration(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalDeleteAdConfiguration(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

// --------------------------------------
// User

func (h *HandlerService) InternalGetUserInfo(w http.ResponseWriter, r *http.Request) {
	// TODO patrick
	user, err := h.getUserBySession(r)
	if err != nil {
		handleErr(w, err)
		return
	}
	userInfo, err := h.dai.GetUserInfo(r.Context(), user.Id)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetUserInfoResponse{
		Data: *userInfo,
	}
	returnOk(w, response)
}

// --------------------------------------
// Dashboards

func (h *HandlerService) InternalGetUserDashboards(w http.ResponseWriter, r *http.Request) {
	userId, err := h.GetUserIdBySession(r)
	if err != nil {
		handleErr(w, err)
		return
	}
	data, err := h.dai.GetUserDashboards(r.Context(), userId)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.ApiDataResponse[types.UserDashboardsData]{
		Data: *data,
	}
	returnOk(w, response)
}

// --------------------------------------
// Account Dashboards

func (h *HandlerService) InternalPostAccountDashboards(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h *HandlerService) InternalGetAccountDashboard(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalDeleteAccountDashboard(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h *HandlerService) InternalPostAccountDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h *HandlerService) InternalDeleteAccountDashboardGroups(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h *HandlerService) InternalPostAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h *HandlerService) InternalGetAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalDeleteAccountDashboardAccounts(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h *HandlerService) InternalPutAccountDashboardAccount(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalPostAccountDashboardPublicIds(w http.ResponseWriter, r *http.Request) {
	returnCreated(w, nil)
}

func (h *HandlerService) InternalPutAccountDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalDeleteAccountDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	returnNoContent(w)
}

func (h *HandlerService) InternalGetAccountDashboardTransactions(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

func (h *HandlerService) InternalPutAccountDashboardTransactionsSettings(w http.ResponseWriter, r *http.Request) {
	returnOk(w, nil)
}

// --------------------------------------
// Validator Dashboards

func (h *HandlerService) InternalPostValidatorDashboards(w http.ResponseWriter, r *http.Request) {
	var v validationError
	userId, ok := r.Context().Value(ctxUserIdKey).(uint64)
	if !ok {
		handleErr(w, errors.New("error getting user id from context"))
		return
	}
	req := struct {
		Name    string      `json:"name"`
		Network intOrString `json:"network"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	name := v.checkNameNotEmpty(req.Name)
	chainId := v.checkNetwork(req.Network)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	userInfo, err := h.dai.GetUserInfo(r.Context(), userId)
	if err != nil {
		handleErr(w, err)
		return
	}
	dashboardCount, err := h.dai.GetUserValidatorDashboardCount(r.Context(), userId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if dashboardCount >= userInfo.PremiumPerks.ValidatorDasboards {
		returnConflict(w, errors.New("maximum number of validator dashboards reached"))
		return
	}

	data, err := h.dai.CreateValidatorDashboard(r.Context(), userId, name, chainId)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.ApiDataResponse[types.VDBPostReturnData]{
		Data: *data,
	}
	returnCreated(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboard(w http.ResponseWriter, r *http.Request) {
	dashboardIdParam := mux.Vars(r)["dashboard_id"]
	dashboardId, err := h.handleDashboardId(r.Context(), dashboardIdParam)
	if err != nil {
		handleErr(w, err)
		return
	}
	// set name depending on dashboard id
	var name string
	if reInteger.MatchString(dashboardIdParam) {
		name, err = h.dai.GetValidatorDashboardName(r.Context(), dashboardId.Id)
	} else if reValidatorDashboardPublicId.MatchString(dashboardIdParam) {
		var publicIdInfo *types.VDBPublicId
		publicIdInfo, err = h.dai.GetValidatorDashboardPublicId(r.Context(), types.VDBIdPublic(dashboardIdParam))
		name = publicIdInfo.Name
	}
	if err != nil {
		handleErr(w, err)
		return
	}

	// add premium chart perk info for shared dashboards
	premiumPerks, err := h.getDashboardPremiumPerks(r.Context(), *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardOverview(r.Context(), *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	data.ChartHistorySeconds = premiumPerks.ChartHistorySeconds
	data.Name = name

	response := types.InternalGetValidatorDashboardResponse{
		Data: *data,
	}

	returnOk(w, response)
}

func (h *HandlerService) InternalDeleteValidatorDashboard(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	err := h.dai.RemoveValidatorDashboard(r.Context(), dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	returnNoContent(w)
}

func (h *HandlerService) InternalPutValidatorDashboardName(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	req := struct {
		Name string `json:"name"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	name := v.checkNameNotEmpty(req.Name)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	data, err := h.dai.UpdateValidatorDashboardName(r.Context(), dashboardId, name)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.ApiDataResponse[types.VDBPostReturnData]{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalPostValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	req := struct {
		Name string `json:"name"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	name := v.checkNameNotEmpty(req.Name)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	ctx := r.Context()
	// check if user has reached the maximum number of groups
	userId, ok := ctx.Value(ctxUserIdKey).(uint64)
	if !ok {
		handleErr(w, errors.New("error getting user id from context"))
		return
	}
	userInfo, err := h.dai.GetUserInfo(ctx, userId)
	if err != nil {
		handleErr(w, err)
		return
	}
	groupCount, err := h.dai.GetValidatorDashboardGroupCount(ctx, dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if groupCount >= userInfo.PremiumPerks.ValidatorGroupsPerDashboard {
		returnConflict(w, errors.New("maximum number of validator dashboard groups reached"))
		return
	}

	data, err := h.dai.CreateValidatorDashboardGroup(ctx, dashboardId, name)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.ApiResponse{
		Data: data,
	}

	returnCreated(w, response)
}

func (h *HandlerService) InternalPutValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId := v.checkPrimaryDashboardId(vars["dashboard_id"])
	groupId := v.checkExistingGroupId(vars["group_id"])
	req := struct {
		Name string `json:"name"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	name := v.checkNameNotEmpty(req.Name)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	groupExists, err := h.dai.GetValidatorDashboardGroupExists(r.Context(), dashboardId, groupId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if !groupExists {
		returnNotFound(w, errors.New("group not found"))
		return
	}
	data, err := h.dai.UpdateValidatorDashboardGroup(r.Context(), dashboardId, groupId, name)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.ApiResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h *HandlerService) InternalDeleteValidatorDashboardGroups(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	groupId := v.checkExistingGroupId(vars["group_id"])
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	if groupId == types.DefaultGroupId {
		returnBadRequest(w, errors.New("cannot delete default group"))
		return
	}
	groupExists, err := h.dai.GetValidatorDashboardGroupExists(r.Context(), dashboardId, groupId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if !groupExists {
		returnNotFound(w, errors.New("group not found"))
		return
	}
	err = h.dai.RemoveValidatorDashboardGroup(r.Context(), dashboardId, groupId)
	if err != nil {
		handleErr(w, err)
		return
	}

	returnNoContent(w)
}

func (h *HandlerService) InternalPostValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	req := struct {
		GroupId           uint64        `json:"group_id,omitempty"`
		Validators        []intOrString `json:"validators,omitempty"`
		DepositAddress    string        `json:"deposit_address,omitempty"`
		WithdrawalAddress string        `json:"withdrawal_address,omitempty"`
		Graffiti          string        `json:"graffiti,omitempty"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	// check if exactly one of validators, deposit_address, withdrawal_address, graffiti is set
	fields := []interface{}{req.Validators, req.DepositAddress, req.WithdrawalAddress, req.Graffiti}
	var count int
	for _, set := range fields {
		if !reflect.ValueOf(set).IsZero() {
			count++
		}
	}
	if count != 1 {
		v.add("request body", "exactly one of `validators`, `deposit_address`, `withdrawal_address`, `graffiti` must be set. please check the API documentation for more information")
	}
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	groupId := req.GroupId
	ctx := r.Context()
	groupExists, err := h.dai.GetValidatorDashboardGroupExists(ctx, dashboardId, groupId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if !groupExists {
		returnNotFound(w, errors.New("group not found"))
		return
	}
	userId, ok := ctx.Value(ctxUserIdKey).(uint64)
	if !ok {
		handleErr(w, errors.New("error getting user id from context"))
		return
	}
	userInfo, err := h.dai.GetUserInfo(ctx, userId)
	if err != nil {
		handleErr(w, err)
		return
	}
	limit := userInfo.PremiumPerks.ValidatorsPerDashboard
	if req.Validators == nil && !userInfo.PremiumPerks.BulkAdding {
		returnConflict(w, errors.New("bulk adding not allowed with current subscription plan"))
		return
	}
	var data []types.VDBPostValidatorsData
	var dataErr error
	switch {
	case req.Validators != nil:
		indices, pubkeys := v.checkValidators(req.Validators, forbidEmpty)
		if v.hasErrors() {
			handleErr(w, v)
			return
		}
		validators, err := h.dai.GetValidatorsFromSlices(indices, pubkeys)
		if err != nil {
			handleErr(w, err)
			return
		}
		// check if adding more validators than allowed
		existingValidatorCount, err := h.dai.GetValidatorDashboardExistingValidatorCount(ctx, dashboardId, validators)
		if err != nil {
			handleErr(w, err)
			return
		}
		if uint64(len(validators)) > existingValidatorCount+limit {
			returnConflict(w, fmt.Errorf("adding more validators than allowed, limit is %v new validators", limit))
			return
		}
		data, dataErr = h.dai.AddValidatorDashboardValidators(ctx, dashboardId, groupId, validators)

	case req.DepositAddress != "":
		depositAddress := v.checkRegex(reEthereumAddress, req.DepositAddress, "deposit_address")
		if v.hasErrors() {
			handleErr(w, v)
			return
		}
		data, dataErr = h.dai.AddValidatorDashboardValidatorsByDepositAddress(ctx, dashboardId, groupId, depositAddress, limit)

	case req.WithdrawalAddress != "":
		withdrawalAddress := v.checkRegex(reWithdrawalCredential, req.WithdrawalAddress, "withdrawal_address")
		if v.hasErrors() {
			handleErr(w, v)
			return
		}
		data, dataErr = h.dai.AddValidatorDashboardValidatorsByWithdrawalAddress(ctx, dashboardId, groupId, withdrawalAddress, limit)

	case req.Graffiti != "":
		graffiti := v.checkRegex(reNonEmpty, req.Graffiti, "graffiti")
		if v.hasErrors() {
			handleErr(w, v)
			return
		}
		data, dataErr = h.dai.AddValidatorDashboardValidatorsByGraffiti(ctx, dashboardId, groupId, graffiti, limit)
	}

	if dataErr != nil {
		handleErr(w, dataErr)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	returnCreated(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	q := r.URL.Query()
	groupId := v.checkGroupId(q.Get("group_id"), allowEmpty)
	pagingParams := v.checkPagingParams(q)
	sort := checkSort[enums.VDBManageValidatorsColumn](&v, q.Get("sort"))
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	data, paging, err := h.dai.GetValidatorDashboardValidators(r.Context(), *dashboardId, groupId, pagingParams.cursor, *sort, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardValidatorsResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalDeleteValidatorDashboardValidators(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	var indices []uint64
	var publicKeys []string
	if validatorsParam := r.URL.Query().Get("validators"); validatorsParam != "" {
		indices, publicKeys = v.checkValidatorList(validatorsParam, allowEmpty)
		if v.hasErrors() {
			handleErr(w, v)
			return
		}
	}
	validators, err := h.dai.GetValidatorsFromSlices(indices, publicKeys)
	if err != nil {
		handleErr(w, err)
		return
	}
	err = h.dai.RemoveValidatorDashboardValidators(r.Context(), dashboardId, validators)
	if err != nil {
		handleErr(w, err)
		return
	}

	returnNoContent(w)
}

func (h *HandlerService) InternalPostValidatorDashboardPublicIds(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	req := struct {
		Name          string `json:"name,omitempty"`
		ShareSettings struct {
			ShareGroups bool `json:"share_groups"`
		} `json:"share_settings"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	name := v.checkName(req.Name, 0)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	publicIdCount, err := h.dai.GetValidatorDashboardPublicIdCount(r.Context(), dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if publicIdCount >= 1 {
		returnConflict(w, errors.New("cannot create more than one public id"))
		return
	}

	data, err := h.dai.CreateValidatorDashboardPublicId(r.Context(), dashboardId, name, req.ShareSettings.ShareGroups)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	returnCreated(w, response)
}

func (h *HandlerService) InternalPutValidatorDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	req := struct {
		Name          string `json:"name"`
		ShareSettings struct {
			ShareGroups bool `json:"share_groups"`
		} `json:"share_settings"`
	}{}
	if err := v.checkBody(&req, r); err != nil {
		handleErr(w, err)
		return
	}
	name := v.checkNameNotEmpty(req.Name)
	publicDashboardId := v.checkValidatorDashboardPublicId(vars["public_id"])
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	dashboardInfo, err := h.dai.GetValidatorDashboardInfoByPublicId(r.Context(), publicDashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if dashboardInfo.Id != dashboardId {
		handleErr(w, newNotFoundErr("public id %v not found", publicDashboardId))
	}

	data, err := h.dai.UpdateValidatorDashboardPublicId(r.Context(), publicDashboardId, name, req.ShareSettings.ShareGroups)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.ApiResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h *HandlerService) InternalDeleteValidatorDashboardPublicId(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId := v.checkPrimaryDashboardId(mux.Vars(r)["dashboard_id"])
	publicDashboardId := v.checkValidatorDashboardPublicId(vars["public_id"])
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	dashboardInfo, err := h.dai.GetValidatorDashboardInfoByPublicId(r.Context(), publicDashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	if dashboardInfo.Id != dashboardId {
		handleErr(w, newNotFoundErr("public id %v not found", publicDashboardId))
	}

	err = h.dai.RemoveValidatorDashboardPublicId(r.Context(), publicDashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}

	returnNoContent(w)
}

func (h *HandlerService) InternalGetValidatorDashboardSlotViz(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}

	groupIds := v.checkExistingGroupIdList(r.URL.Query().Get("group_ids"))
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	data, err := h.dai.GetValidatorDashboardSlotViz(r.Context(), *dashboardId, groupIds)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardSlotVizResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardSummary(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	q := r.URL.Query()
	pagingParams := v.checkPagingParams(q)
	sort := checkSort[enums.VDBSummaryColumn](&v, q.Get("sort"))

	period := checkEnum[enums.TimePeriod](&v, q.Get("period"), "period")
	// allowed periods are: all_time, last_30d, last_7d, last_24h, last_1h
	allowedPeriods := []enums.Enum{enums.TimePeriods.AllTime, enums.TimePeriods.Last30d, enums.TimePeriods.Last7d, enums.TimePeriods.Last24h, enums.TimePeriods.Last1h}
	v.checkEnumIsAllowed(period, allowedPeriods, "period")
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardSummary(r.Context(), *dashboardId, period, pagingParams.cursor, *sort, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardSummaryResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardGroupSummary(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId, err := h.handleDashboardId(r.Context(), vars["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	groupId := v.checkGroupId(vars["group_id"], forbidEmpty)
	period := checkEnum[enums.TimePeriod](&v, r.URL.Query().Get("period"), "period")
	// allowed periods are: all_time, last_30d, last_7d, last_24h, last_1h
	allowedPeriods := []enums.Enum{enums.TimePeriods.AllTime, enums.TimePeriods.Last30d, enums.TimePeriods.Last7d, enums.TimePeriods.Last24h, enums.TimePeriods.Last1h}
	v.checkEnumIsAllowed(period, allowedPeriods, "period")
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, err := h.dai.GetValidatorDashboardGroupSummary(r.Context(), *dashboardId, groupId, period)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupSummaryResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardSummaryChart(w http.ResponseWriter, r *http.Request) {
	var v validationError
	ctx := r.Context()
	dashboardId, err := h.handleDashboardId(ctx, mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	premiumPerks, err := h.getDashboardPremiumPerks(ctx, *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	q := r.URL.Query()
	groupIds := v.checkGroupIdList(q.Get("group_ids"))
	efficiencyType := checkEnum[enums.VDBSummaryChartEfficiencyType](&v, q.Get("efficiency_type"), "efficiency_type")
	aggregation := checkEnum[enums.ChartAggregation](&v, q.Get("aggregation"), "aggregation")
	maxAge := getMaxChartAge(aggregation, premiumPerks.ChartHistorySeconds)
	minAllowedTs := uint64(time.Now().Unix()) - maxAge
	afterTs, beforeTs := v.checkTimestamps(q.Get("after_ts"), q.Get("before_ts"), minAllowedTs)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	if maxAge == 0 {
		returnConflict(w, fmt.Errorf("requested aggregation is not available for dashboard owner's premium subscription"))
		return
	}
	// afterTs is inclusive, beforeTs is exclusive
	if afterTs < minAllowedTs || beforeTs <= minAllowedTs {
		returnConflict(w, fmt.Errorf("requested time range is too old, maximum age for dashboard owner's premium subscription is %v seconds", maxAge))
		return
	}

	data, err := h.dai.GetValidatorDashboardSummaryChart(ctx, *dashboardId, groupIds, efficiencyType, aggregation, afterTs, beforeTs)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardSummaryChartResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardSummaryValidators(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	groupId := v.checkGroupId(r.URL.Query().Get("group_id"), allowEmpty)
	q := r.URL.Query()
	duty := checkEnum[enums.ValidatorDuty](&v, q.Get("duty"), "duty")
	period := checkEnum[enums.TimePeriod](&v, q.Get("period"), "period")
	// allowed periods are: all_time, last_30d, last_7d, last_24h, last_1h
	allowedPeriods := []enums.Enum{enums.TimePeriods.AllTime, enums.TimePeriods.Last30d, enums.TimePeriods.Last7d, enums.TimePeriods.Last24h, enums.TimePeriods.Last1h}
	v.checkEnumIsAllowed(period, allowedPeriods, "period")
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	// get indices based on duty
	var indices interface{}
	duties := enums.ValidatorDuties
	switch duty {
	case duties.None:
		indices, err = h.dai.GetValidatorDashboardSummaryValidators(r.Context(), *dashboardId, groupId)
	case duties.Sync:
		indices, err = h.dai.GetValidatorDashboardSyncSummaryValidators(r.Context(), *dashboardId, groupId, period)
	case duties.Slashed:
		indices, err = h.dai.GetValidatorDashboardSlashingsSummaryValidators(r.Context(), *dashboardId, groupId, period)
	case duties.Proposal:
		indices, err = h.dai.GetValidatorDashboardProposalSummaryValidators(r.Context(), *dashboardId, groupId, period)
	}
	if err != nil {
		handleErr(w, err)
		return
	}
	// map indices to response format
	data, err := mapVDBIndices(indices)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.InternalGetValidatorDashboardSummaryValidatorsResponse{
		Data: data,
	}

	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardRewards(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	q := r.URL.Query()
	pagingParams := v.checkPagingParams(q)
	sort := checkSort[enums.VDBRewardsColumn](&v, q.Get("sort"))
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardRewards(r.Context(), *dashboardId, pagingParams.cursor, *sort, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardRewardsResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardGroupRewards(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId, err := h.handleDashboardId(r.Context(), vars["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	groupId := v.checkGroupId(vars["group_id"], forbidEmpty)
	epoch := v.checkUint(vars["epoch"], "epoch")
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, err := h.dai.GetValidatorDashboardGroupRewards(r.Context(), *dashboardId, groupId, epoch)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupRewardsResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardRewardsChart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dashboardId, err := h.handleDashboardId(r.Context(), vars["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}

	data, err := h.dai.GetValidatorDashboardRewardsChart(r.Context(), *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardRewardsChartResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardDuties(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId, err := h.handleDashboardId(r.Context(), vars["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	q := r.URL.Query()
	groupId := v.checkGroupId(q.Get("group_id"), allowEmpty)
	epoch := v.checkUint(vars["epoch"], "epoch")
	pagingParams := v.checkPagingParams(q)
	sort := checkSort[enums.VDBDutiesColumn](&v, q.Get("sort"))
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardDuties(r.Context(), *dashboardId, epoch, groupId, pagingParams.cursor, *sort, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardDutiesResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardBlocks(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	q := r.URL.Query()
	pagingParams := v.checkPagingParams(q)
	sort := checkSort[enums.VDBBlocksColumn](&v, q.Get("sort"))
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardBlocks(r.Context(), *dashboardId, pagingParams.cursor, *sort, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardBlocksResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardEpochHeatmap(w http.ResponseWriter, r *http.Request) {
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}

	// implicit time period is last hour
	data, err := h.dai.GetValidatorDashboardEpochHeatmap(r.Context(), *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardHeatmapResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardDailyHeatmap(w http.ResponseWriter, r *http.Request) {
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}

	var v validationError
	period := checkEnum[enums.TimePeriod](&v, r.URL.Query().Get("period"), "period")
	// allowed periods are: last_7d, last_30d, last_365d
	allowedPeriods := []enums.Enum{enums.TimePeriods.Last7d, enums.TimePeriods.Last30d, enums.TimePeriods.Last365d}
	v.checkEnumIsAllowed(period, allowedPeriods, "period")
	if v.hasErrors() {
		handleErr(w, v)
		return
	}
	data, err := h.dai.GetValidatorDashboardDailyHeatmap(r.Context(), *dashboardId, period)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardHeatmapResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardGroupEpochHeatmap(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	groupId := v.checkExistingGroupId(vars["group_id"])
	epoch := v.checkUint(vars["epoch"], "epoch")
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, err := h.dai.GetValidatorDashboardGroupEpochHeatmap(r.Context(), *dashboardId, groupId, epoch)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupHeatmapResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardGroupDailyHeatmap(w http.ResponseWriter, r *http.Request) {
	var v validationError
	vars := mux.Vars(r)
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	groupId := v.checkExistingGroupId(vars["group_id"])
	date := v.checkDate(vars["date"])
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, err := h.dai.GetValidatorDashboardGroupDailyHeatmap(r.Context(), *dashboardId, groupId, date)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardGroupHeatmapResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardExecutionLayerDeposits(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	pagingParams := v.checkPagingParams(r.URL.Query())
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardElDeposits(r.Context(), *dashboardId, pagingParams.cursor, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardExecutionLayerDepositsResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardConsensusLayerDeposits(w http.ResponseWriter, r *http.Request) {
	var v validationError
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	pagingParams := v.checkPagingParams(r.URL.Query())
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardClDeposits(r.Context(), *dashboardId, pagingParams.cursor, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.InternalGetValidatorDashboardConsensusLayerDepositsResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardTotalConsensusLayerDeposits(w http.ResponseWriter, r *http.Request) {
	var err error
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardTotalClDeposits(r.Context(), *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.InternalGetValidatorDashboardTotalConsensusDepositsResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardTotalExecutionLayerDeposits(w http.ResponseWriter, r *http.Request) {
	var err error
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	data, err := h.dai.GetValidatorDashboardTotalElDeposits(r.Context(), *dashboardId)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.InternalGetValidatorDashboardTotalExecutionDepositsResponse{
		Data: *data,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardWithdrawals(w http.ResponseWriter, r *http.Request) {
	var v validationError
	q := r.URL.Query()
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	pagingParams := v.checkPagingParams(q)
	sort := checkSort[enums.VDBWithdrawalsColumn](&v, q.Get("sort"))
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, paging, err := h.dai.GetValidatorDashboardWithdrawals(r.Context(), *dashboardId, pagingParams.cursor, *sort, pagingParams.search, pagingParams.limit)
	if err != nil {
		handleErr(w, err)
		return
	}
	response := types.InternalGetValidatorDashboardWithdrawalsResponse{
		Data:   data,
		Paging: *paging,
	}
	returnOk(w, response)
}

func (h *HandlerService) InternalGetValidatorDashboardTotalWithdrawals(w http.ResponseWriter, r *http.Request) {
	var v validationError
	q := r.URL.Query()
	dashboardId, err := h.handleDashboardId(r.Context(), mux.Vars(r)["dashboard_id"])
	if err != nil {
		handleErr(w, err)
		return
	}
	pagingParams := v.checkPagingParams(q)
	if v.hasErrors() {
		handleErr(w, v)
		return
	}

	data, err := h.dai.GetValidatorDashboardTotalWithdrawals(r.Context(), *dashboardId, pagingParams.search)
	if err != nil {
		handleErr(w, err)
		return
	}

	response := types.InternalGetValidatorDashboardTotalWithdrawalsResponse{
		Data: *data,
	}
	returnOk(w, response)
}
