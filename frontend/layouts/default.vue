<template>
  <v-app>
    <v-toolbar
      dark
      color="primary"
      class="font-italic"
    >
      <v-icon>library_books</v-icon>

      <v-toolbar-title>
        FirebaseAuth＋Nuxt.js＋Go(v1.11)＋GAE開発
      </v-toolbar-title>

      <v-spacer />

      <v-btn
        icon
        @click="home"
      >
        <v-icon>home</v-icon>
      </v-btn>

      <v-btn
        icon
        @click="signup"
      >
        <v-icon>person_add</v-icon>
      </v-btn>

      <v-btn
        icon
        @click="logout"
      >
        <v-icon>exit_to_app</v-icon>
      </v-btn>
    </v-toolbar>
    <v-content>
      <nuxt />
    </v-content>
    <v-footer
      dark
      color="primary"
      class="pa-3"
    >
      <v-spacer />
      <div>&copy; {{ new Date().getFullYear() }}</div>
    </v-footer>

    <!-- エラー通知用スナックバー -->
    <notification />
  </v-app>
</template>

<script>
import notification from '~/components/notification.vue'
import firebase from '~/plugins/firebase'

export default {
  components: {
    notification
  },

  methods: {
    home() {
      this.$router.push('/')
    },
    signup() {
      this.$router.push('/signup')
    },
    async logout() {
      await firebase
        .auth()
        .signOut()
        .then((res) => {
          // ログアウト正常終了時はログイン画面に遷移する。
          this.$router.push('/login')
        })
        .catch((error) => {
          console.log(
            'errorCode:' + error.code + ', errorMessage:' + error.message
          )
        })
    }
  }
}
</script>
