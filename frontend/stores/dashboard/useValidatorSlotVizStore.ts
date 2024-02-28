import { defineStore } from 'pinia'
import { useCustomFetch } from '~/composables/useCustomFetch'
import type { SlotVizEpoch, InternalGetValidatorDashboardSlotVizResponse } from '~/types/api/slot_viz'

export const useValidatorSlotVizStore = defineStore('validator_slotViz', () => {
  const slotViz = ref<SlotVizEpoch[] | undefined | null>()
  async function getSlotViz () {
    const res = await useCustomFetch<InternalGetValidatorDashboardSlotVizResponse>(API_PATH.DASHBOARD_SLOTVIZ)
    slotViz.value = res.data
    return slotViz.value
  }

  return { slotViz, getSlotViz }
})
