<script lang="ts" setup>
const id = useId()

const input = defineModel<string>()

defineProps<{
  label: string
}>()

</script>

<template>
  <label :for="id">{{ label }}</label>
  <InputText :id="id" v-model="input" v-bind="$attrs"/>
</template>

<style lang="scss">

</style>