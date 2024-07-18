<script lang="ts" setup>
import {
  faInfoCircle
} from '@fortawesome/pro-regular-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
const props = defineProps<{
  label?: string,
  infoText?: string,
}>()

const id = props.label ? useId() : undefined

const input = defineModel<boolean>()

</script>

<template>
  <span>
    <label 
      v-if="label"
      class="label"
      :for="id"
    >
      {{ label }}
    </label>
  <BcTooltip 
    v-if="infoText"
    class="bc-input-checkbox__info"
    tooltip-width="217px"
    tooltip-text-align="left"
  >
    <FontAwesomeIcon :icon="faInfoCircle" />
    <template #tooltip>
      {{ infoText }}
    </template>
  </BcTooltip>
</span>
  <Checkbox 
    v-model="input"
    :input-id="id" 
    v-bind="$attrs"
    binary
    />
</template> 

<style lang="scss">
.label{
  cursor: pointer;
}
.bc-input-checkbox__info {
  margin-left: var(--padding);
}
</style>