<template>
  <v-layout
    class="py-3"
  >
    <v-form>
      <v-text-field
        v-model="email"
        label="Email"
        class="py-2 ml-5"
      />
      <v-text-field
        v-model="password"
        label="Pass"
        :type="`password`"
        class="py-2 ml-5"
      />
      <v-btn
        class="mx-1 my-2 px-3 py-2 lime"
        @click="login"
      >
        LOGIN
      </v-btn>
    </v-form>
    <div style="color: red;">
      {{ errMsg }}
    </div>
  </v-layout>
</template>

<script>
import firebase from '~/plugins/firebase'

export default {
  data() {
    return {
      email: '',
      password: '',
      errMsg: ''
    }
  },

  async mounted() {
    await firebase.auth().onAuthStateChanged((user) => {
      if (user) {
      } else {
      }
    })
  },

  methods: {
    async login() {
      await firebase
        .auth()
        .signInWithEmailAndPassword(this.email, this.password)
        .then((res) => {
          // ログイン正常終了時はログイン後の初期画面に遷移する。
          this.$router.push('/')
        })
        .catch((error) => {
          this.errMsg = error.message
          console.log(
            'errorCode:' + error.code + ', errorMessage:' + error.message
          )
        })
    }
  }
}
</script>
