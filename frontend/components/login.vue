<template>
  <v-layout
    class="py-3"
  >
    <v-form>
      <v-text-field
        v-model="email"
        label="Email"
        class="py-2"
      />
      <v-text-field
        v-model="password"
        label="Password"
        :type="`password`"
        class="py-2"
      />
      <v-btn
        class="mx-1 my-2 px-3 py-2 lime"
        @click="login"
      >
        LOGIN
      </v-btn>
    </v-form>
  </v-layout>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      password: ''
    }
  },

  methods: {
    login: function() {
      console.log('login')
      // FIXME: ここにFirebaseAuthを用いたログイン処理を書く想定

      // ログイン正常終了時はログイン後の初期画面に遷移する。
      this.$router.push('/')
    }
  }
}
</script>
