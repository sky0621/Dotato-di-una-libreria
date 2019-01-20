<template>
  <v-layout
    class="py-3"
  >
    <v-list two-line>
      <v-list-tile v-for="notice in notices" :key="notice.id">
        <v-list-tile-content class="mb-2">
          <v-list-tile-title>{{ notice.sentence }}</v-list-tile-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
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
      .then(res => {
        console.log(res.data)
        this.notices = res.data
      })
      .catch(err => {
        console.log(err)
      })
  }
}
</script>
