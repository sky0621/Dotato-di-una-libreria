<template>
  <v-app>
    <nuxt />

    <!-- エラー通知用スナックバー -->
    <notification />
  </v-app>
</template>

<script>
import notification from '~/components/notification.vue'

export default {
  components: {
    notification
  }
}
</script>
