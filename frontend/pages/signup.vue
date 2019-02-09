<template>
  <v-container fluid>
    <v-layout
      align-start
      justify-start
      column
    >
      <v-flex>
        <v-label>ユーザー登録</v-label>
      </v-flex>
      <v-flex>
        <signup />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import signup from '~/components/signup.vue'

export default {
  components: {
    signup
  }
}
</script>
