import { defineStore } from 'pinia'
import { useAllValidatorDashboardRewardsDetailsStore } from './useValidatorDashboardRewardsDetailsStore'
import type { VDBOverviewData, InternalGetValidatorDashboardResponse } from '~/types/api/validator_dashboard'
import type { DashboardKey } from '~/types/dashboard'
import { API_PATH } from '~/types/customFetch'

const validatorOverviewStore = defineStore('validator_overview_store', () => {
  const data = ref<VDBOverviewData | undefined | null>()
  return { data }
})

export function useValidatorDashboardOverviewStore () {
  const { fetch } = useCustomFetch()
  const { data } = storeToRefs(validatorOverviewStore())
  const { clearCache: clearRewardDetails } = useAllValidatorDashboardRewardsDetailsStore()

  const overview = computed(() => data.value)

  async function refreshOverview (key: DashboardKey) {
    if (!key) {
      data.value = undefined
      return
    }
    const res = await fetch<InternalGetValidatorDashboardResponse>(API_PATH.DASHBOARD_OVERVIEW, undefined, { dashboardKey: key })
    data.value = res.data

    clearOverviewDependentCaches()

    return overview.value
  }

  function clearOverviewDependentCaches () {
    clearRewardDetails()
  }

  return { overview, refreshOverview }
}
