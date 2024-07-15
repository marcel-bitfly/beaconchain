<script setup lang="ts">
import type { CryptoCurrency, Currency, FiatCurrency } from '~/types/currencies'

interface Props {
  currency?: Currency
}

defineProps<Props>()


</script>
<template>
  <IconFiat v-if="isFiat(currency)" :currency="currency as FiatCurrency" />
  <IconCrypto v-if="isCrypto(currency)" :currency="currency as CryptoCurrency" />
  <IconNetworkEthereum v-else class="monochromatic" />
</template>
