<template>
  <v-layout
    class="ma-2"
  >
    <v-expansion-panel focusable>
      <v-expansion-panel-content
        v-for="notice in notices"
        :key="notice.id"
      >
        <div slot="header">
          {{ notice.sentence }}
        </div>
        <v-card>
          <v-card-text class="lime lighten-4">
            お知らせの詳細な内容を、ごにょごにょ〜〜〜〜〜〜〜〜〜〜〜。
          </v-card-text>
        </v-card>
      </v-expansion-panel-content>
    </v-expansion-panel>
  </v-layout>
</template>

<script>
export default {
  data() {
    return {
      notices: []
    }
  },

  mounted() {
    this.$axios
      .get(process.env.apiBaseUrl + '/notices')
      .then((res) => {
        console.log(res.data)
        this.notices = res.data
      })
      .catch((err) => {
        console.log(err)
      })
  }
}
</script>
