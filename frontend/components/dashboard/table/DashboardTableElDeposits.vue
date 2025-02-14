<script setup lang="ts">
import type { DataTableSortEvent } from 'primevue/datatable'
import type { VDBExecutionDepositsTableRow } from '~/types/api/validator_dashboard'
import type { Cursor, TableQueryParams } from '~/types/datatable'
import { useValidatorDashboardOverviewStore } from '~/stores/dashboard/useValidatorDashboardOverviewStore'
import { getGroupLabel } from '~/utils/dashboard/group'
import { useValidatorDashboardElDepositsStore } from '~/stores/dashboard/useValidatorDashboardElDepositsStore'

const { dashboardKey } = useDashboardKey()

const cursor = ref<Cursor>()
const pageSize = ref<number>(5)
const { t: $t } = useI18n()

const { deposits, query: lastQuery, getDeposits, getTotalAmount, totalAmount, isLoadingDeposits, isLoadingTotal } = useValidatorDashboardElDepositsStore()
const { value: query, bounce: setQuery } = useDebounceValue<TableQueryParams | undefined>(undefined, 500)

const { overview, hasValidators } = useValidatorDashboardOverviewStore()
const { groups } = useValidatorDashboardGroups()

const { width } = useWindowSize()
const colsVisible = computed(() => {
  return {
    group: width.value > 1200,
    block: width.value >= 1100,
    withdrawalCredentials: width.value >= 1060,
    from: width.value >= 960,
    depositer: width.value >= 860,
    txHash: width.value >= 760,
    valid: width.value >= 660,
    publicKey: width.value >= 560
  }
})

const loadData = (query?: TableQueryParams) => {
  if (!query) {
    query = { limit: pageSize.value }
  }
  setQuery(query, true, true)
}

watch([dashboardKey, overview], () => {
  loadData()
  getTotalAmount(dashboardKey.value)
}, { immediate: true })

watch(query, async (q) => {
  if (q) {
    await getDeposits(dashboardKey.value, q)
  }
}, { immediate: true })

const tableData = computed(() => {
  if (!deposits.value?.data?.length) {
    return
  }
  return {
    paging: deposits.value.paging,
    data: [
      {
        amount: totalAmount.value
      },
      ...deposits.value.data
    ]
  }
})

const groupNameLabel = (groupId?: number) => {
  return getGroupLabel($t, groupId, groups.value)
}

const onSort = (sort: DataTableSortEvent) => {
  loadData(setQuerySort(sort, lastQuery.value))
}

const setCursor = (value: Cursor) => {
  cursor.value = value
  loadData(setQueryCursor(value, lastQuery.value))
}

const setPageSize = (value: number) => {
  pageSize.value = value
  loadData(setQueryPageSize(value, lastQuery.value))
}

const getRowClass = (row: VDBExecutionDepositsTableRow) => {
  if (row.index === undefined) {
    return 'total-row'
  }
}

const isRowExpandable = (row: VDBExecutionDepositsTableRow) => {
  return row.index !== undefined
}

</script>
<template>
  <div>
    <BcTableControl :title="$t('dashboard.validator.el_deposits.title')">
      <template #table>
        <ClientOnly fallback-tag="span">
          <BcTable
            :data="tableData"
            data-key="index"
            :expandable="!colsVisible.group"
            class="el_deposits_table"
            :cursor="cursor"
            :page-size="pageSize"
            :row-class="getRowClass"
            :is-row-expandable="isRowExpandable"
            :loading="isLoadingDeposits"
            @set-cursor="setCursor"
            @sort="onSort"
            @set-page-size="setPageSize"
          >
            <Column
              v-if="colsVisible.publicKey"
              field="public_key"
              :header="$t('dashboard.validator.col.public_key')"
            >
              <template #body="slotProps">
                <BcFormatHash
                  v-if="slotProps.data.index !== undefined"
                  :hash="slotProps.data.public_key"
                  :no-wrap="true"
                  type="public_key"
                />
                <span v-else>Σ</span>
              </template>
            </Column>
            <Column field="index" :header="$t('common.index')">
              <template #body="slotProps">
                <BcLink
                  v-if="slotProps.data.index !== undefined"
                  :to="`/validator/${slotProps.data.index}`"
                  target="_blank"
                  class="link"
                >
                  {{ slotProps.data.index }}
                </BcLink>
                <span v-else-if="!colsVisible.publicKey">Σ</span>
              </template>
            </Column>
            <Column
              v-if="colsVisible.group"
              field="group_id"
              body-class="group-id"
              header-class="group-id"
              :header="$t('dashboard.validator.col.group')"
            >
              <template #body="slotProps">
                <span v-if="slotProps.data.index !== undefined">
                  {{ groupNameLabel(slotProps.data.group_id) }}
                </span>
              </template>
            </Column>
            <Column
              v-if="colsVisible.block"
              field="block"
              :header="$t('common.block')"
            >
              <template #body="slotProps">
                <BcLink
                  v-if="slotProps.data.index !== undefined"
                  :to="`/block/${slotProps.data.block}`"
                  target="_blank"
                  class="link"
                >
                  <BcFormatNumber :value="slotProps.data.block" />
                </BcLink>
              </template>
            </Column>
            <Column field="age" body-class="age-field">
              <template #header>
                <BcTableAgeHeader />
              </template>
              <template #body="slotProps">
                <BcFormatTimePassed
                  v-if="slotProps.data.index !== undefined"
                  :value="slotProps.data.timestamp"
                  type="go-timestamp"
                />
              </template>
            </Column>
            <Column v-if="colsVisible.from" :header="$t('table.from')">
              <template #body="slotProps">
                <BcFormatHash
                  v-if="slotProps.data.index !== undefined"
                  :hash="slotProps.data.from.hash"
                  :ens="slotProps.data.from.ens"
                  :no-wrap="true"
                  type="address"
                />
              </template>
            </Column>
            <Column
              v-if="colsVisible.depositer"
              field="depositor"
              :header="$t('dashboard.validator.col.depositor')"
            >
              <template #body="slotProps">
                <BcFormatHash
                  v-if="slotProps.data.index !== undefined"
                  :hash="slotProps.data.depositor.hash"
                  :ens="slotProps.data.depositor.ens"
                  :no-wrap="true"
                  type="address"
                />
              </template>
            </Column>
            <Column v-if="colsVisible.txHash" :header="$t('block.col.tx_hash')">
              <template #body="slotProps">
                <BcFormatHash v-if="slotProps.data.index !== undefined" :hash="slotProps.data.tx_hash" :no-wrap="true" type="tx" />
              </template>
            </Column>
            <Column
              v-if="colsVisible.withdrawalCredentials"
              header-class="withdrawal-credentials"
              :header="$t('dashboard.validator.col.withdrawal_credential')"
            >
              <template #body="slotProps">
                <BcFormatHash
                  v-if="slotProps.data.index !== undefined"
                  :hash="slotProps.data.withdrawal_credential"
                  :no-wrap="true"
                  type="withdrawal_credentials"
                />
              </template>
            </Column>
            <Column field="amount" :header="$t('table.amount')">
              <template #body="slotProps">
                <div v-if="slotProps.data.index === undefined && isLoadingTotal">
                  <BcLoadingSpinner :loading="true" size="small" />
                </div>
                <BcFormatValue v-else :value="slotProps.data.amount" :options="{ fixedDecimalCount: 0 }" />
              </template>
            </Column>
            <Column
              v-if="colsVisible.valid"
              field="valid"
              :header="$t('table.valid')"
            >
              <template #body="slotProps">
                <BcTableValidTag v-if="slotProps.data.index !== undefined" :valid="slotProps.data.valid" />
              </template>
            </Column>
            <template #expansion="slotProps">
              <div class="expansion">
                <div class="row">
                  <div class="label">
                    {{ $t('dashboard.validator.col.public_key') }}
                  </div>
                  <BcFormatHash :hash="slotProps.data.public_key" type="public_key" :no-wrap="true" />
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('dashboard.validator.col.group') }}
                  </div>
                  <div class="value">
                    {{ groupNameLabel(slotProps.data.group_id) }}
                  </div>
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('common.block') }}
                  </div>
                  <BcLink :to="`/block/${slotProps.data.block}`" target="_blank" class="link">
                    <BcFormatNumber :value="slotProps.data.block" />
                  </BcLink>
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('table.from') }}
                  </div>
                  <BcFormatHash
                    v-if="slotProps.data.index !== undefined"
                    :hash="slotProps.data.from.hash"
                    :ens="slotProps.data.from.ens"
                    :no-wrap="true"
                    type="address"
                  />
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('dashboard.validator.col.depositor') }}
                  </div>
                  <BcFormatHash
                    v-if="slotProps.data.index !== undefined"
                    :hash="slotProps.data.depositor.hash"
                    :ens="slotProps.data.depositor.ens"
                    :no-wrap="true"
                    type="address"
                  />
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('block.col.tx_hash') }}
                  </div>
                  <BcFormatHash v-if="slotProps.data.index !== undefined" :hash="slotProps.data.tx_hash" :no-wrap="true" type="tx" />
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('dashboard.validator.col.withdrawal_credential') }}
                  </div>
                  <BcFormatHash :hash="slotProps.data.withdrawal_credential" type="withdrawal_credentials" :no-wrap="true" />
                </div>
                <div class="row">
                  <div class="label">
                    {{ $t('table.valid') }}
                  </div>
                  <div>
                    <BcTableValidTag :valid="slotProps.data.valid" />
                  </div>
                </div>
              </div>
            </template>
            <template #empty>
              <DashboardTableAddValidator v-if="!hasValidators" />
            </template>
          </BcTable>
        </ClientOnly>
      </template>
    </BcTableControl>
  </div>
</template>

<style lang="scss" scoped>
@use "~/assets/css/utils.scss";

:deep(.el_deposits_table) {
  >.p-datatable-wrapper {
    min-height: 335px;
  }

  .withdrawal-credentials {
    @include utils.truncate-text;
  }

  .group-id {
    @include utils.set-all-width(120px);
    @include utils.truncate-text;
  }

  .total-row {
    td {
      font-weight: var(--standard_text_medium_font_weight);
      border-bottom-color: var(--primary-color);
      white-space: nowrap;
      overflow: visible;
    }
  }

  .age-field {
    white-space: nowrap;
  }
  tr>td.age-field {
    padding: 0 7px;
    @include utils.set-all-width(110px);
  }
}

.expansion {
  color: var(--container-color);
  background-color: var(--container-background);
  display: flex;
  flex-direction: column;
  gap: var(--padding);
  padding: var(--padding);
  font-size: var(--small_text_font_size);

  .row {
    display: flex;
    gap: var(--padding);

    .label {
      width: 164px;
      font-weight: var(--standard_text_bold_font_weight);
    }

    .value {
      @include utils.truncate-text;
      max-width: 140px;
    }
  }
}
</style>
