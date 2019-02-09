<template>
  <!-- 通知用スナックバー -->
  <v-snackbar
    v-model="snackbar"
    :color="color"
    :timeout="timeout"
  >
    {{ message }}
    <v-btn
      dark
      flat
      @click="closeSnackbar"
    >
      閉じる
    </v-btn>
  </v-snackbar>
</template>

<script>
import * as types from '~/store/types'

export default {
  data() {
    return {
      timeout: 5000,

      snackbar: false,
      message: '',
      color: ''
    }
  },

  computed: {
    notification() {
      return this.$store.getters[types.NOTIFICATION]
    }
  },

  watch: {
    notification(val) {
      this.apply(val)
    }
  },

  // after Create vm.$el and replace "el" with it
  // https://jp.vuejs.org/v2/guide/instance.html#%E3%83%A9%E3%82%A4%E3%83%95%E3%82%B5%E3%82%A4%E3%82%AF%E3%83%AB%E3%83%80%E3%82%A4%E3%82%A2%E3%82%B0%E3%83%A9%E3%83%A0
  mounted() {
    const val = this.$store.getters[types.NOTIFICATION]
    this.apply(val)
  },

  methods: {
    closeSnackbar() {
      this.$store.dispatch(types.DEACTIVATE_NOTIFICATION)
    },

    apply(val) {
      this.message = val.message
      this.color = val.color
      this.snackbar = val.activate
      // TODO: スナックバーのタイムアウト発動タイミングをフックする方法が不明のため、苦肉の策で「同じ時間経過後にスナックバーを非活性化」する対応を仕込む
      if (this.snackbar) {
        console.log('setTimeout')
        setTimeout(this.closeSnackbar, this.timeout)
      }
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
