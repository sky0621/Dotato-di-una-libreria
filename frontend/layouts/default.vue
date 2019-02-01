<template>
  <v-app dark>
    <v-btn
      class="mx-1 my-2 px-3 py-2 lime"
      @click="logout"
    >
      LOGOUT
    </v-btn>
    <nuxt />
  </v-app>
</template>

<script>
import firebase from '~/plugins/firebase'

export default {
  methods: {
    async logout() {
      await firebase
        .auth()
        .signOut()
        .then(res => {
          // ログアウト正常終了時はログイン画面に遷移する。
          this.$router.push('/login')
        })
        .catch(error => {
          console.log(
            'errorCode:' + error.code + ', errorMessage:' + error.message
          )
        })
    }
  }
}
</script>
