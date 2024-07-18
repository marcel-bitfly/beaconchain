<script lang="ts" setup>

defineProps<{
  isDisabled?: boolean,
  isOutlined?: boolean,
  variant?: 'secondary' // | 'red'
}>()

</script>

<template> 
  <Button 
    type="button"
    :outlined="isOutlined"
    :disabled="isDisabled"
    :class="{
      'bc-button--secondary': variant === 'secondary',
      'bc-button--disabled': isDisabled,
      // 'bc-button--red': variant === 'red'
    }"
  >
    <slot/>
    <span
      v-if="$slots.icon"
      class="bc-button__icon"
    >
      <slot name="icon"/>
    </span>
  </Button>

 </template>

<style lang="scss" scoped>
  .bc-button--secondary {
    border-color: var(--button-secondary-border-color);
    background-color: var(--button-secondary-background-color);
    color: var(--color-white);
    &:hover {
      background-color: var(--button-secondary-background-color--hover);
      border-color: var(--button-secondary-border-color);
    }
    &:active {
      background-color: var(--button-secondary-background-color);
      border-color: var(--button-secondary-border-color);
    }
  }
  .bc-button--disabled {
    background-color: var(--button-color-disabled);
    border-color: var(--button-color-disabled);
    color: var(--button-text-color-disabled);
  }
  .bc-button__icon {
    margin-left: var(--padding-small);
  }
</style>