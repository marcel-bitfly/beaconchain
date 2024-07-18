<script lang="ts" setup>
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {
  faPaperPlane
} from '@fortawesome/pro-solid-svg-icons'
import BcInputCheckbox from '~/components/bc/input/BcInputCheckbox.vue';

const { registerAsChildOf } = useBcDialog()

registerAsChildOf('ManagementModal')
// onUnmounted(() => {
//   hasOpenDialogs.value = false
//   console.log('unmounted')
// })
const webHookUrl = ref('')
const hasWebhookDiscord = ref(false)
</script>
<template>
  
  <!-- {{ hasOpenDialogs }} -->
    <BcForm class="notifications-management-dialog-webhook">
      <BcFormRow>
        <BcInputText 
        v-model="webHookUrl"
        :label="$t('notifications.dashboards.dialog.label-webhook-url')"
        :placeholder="$t('notifications.dashboards.dialog.placeholder-webhook')"
        class="notifications-management-dialog-webhook-url"
        />
      </BcFormRow>
      <BcFormRow>
        <BcInputCheckbox
        v-model="hasWebhookDiscord"
        :label="$t('notifications.dashboards.dialog.label-send_via_discord')"
        :placeholder="$t('notifications.dashboards.dialog.placeholder-webhook')"
        :info-text="$t('notifications.dashboards.dialog.info-send_via_discord')"
        />
      </BcFormRow>
    </BcForm>
    <div class="notifications-management-dialog-webhook-footer">
      <BcButton
        font-awesome-icon="faPaperPlane"
        variant="secondary"
        :is-disabled="!webHookUrl"
      >
        {{ $t('notifications.dashboards.dialog.button-webhook-test') }}
        <template #icon>
          <FontAwesomeIcon :icon="faPaperPlane" />      
        </template>
      </BcButton>
      <BcButton >
        {{ $t('navigation.done') }}
      </BcButton>
    </div>
</template>

<style lang="scss" scoped>
.notifications-management-dialog-webhook {
  margin-top: var(--padding-large);
}
:deep(.notifications-management-dialog-webhook-url) {
  width: 200px;
}
.notifications-management-dialog-webhook-footer {
  display: flex;
  gap: var(--padding-small);
  justify-content: flex-end;
  margin-top: var(--padding-large);
}
</style>
