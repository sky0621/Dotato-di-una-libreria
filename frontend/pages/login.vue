<template>
  <v-content>
    <login />
  </v-content>
</template>

<script>
import login from '~/components/login.vue'

export default {
  layout: 'login',

  components: {
    login
  }
}
</script>
