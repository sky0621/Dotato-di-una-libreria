<template>
  <form>
    <v-text-field
      v-model="name"
      :error-messages="nameErrors"
      :counter="10"
      label="ニックネーム"
      required
      @input="$v.name.$touch()"
      @blur="$v.name.$touch()"
    />
    <v-text-field
      v-model="email"
      type="email"
      :error-messages="emailErrors"
      label="メールアドレス"
      required
      @input="$v.email.$touch()"
      @blur="$v.email.$touch()"
    />
    <v-text-field
      v-model="password"
      type="password"
      :error-messages="passwordErrors"
      label="パスワード"
      required
      @input="$v.password.$touch()"
      @blur="$v.password.$touch()"
    />

    <v-btn
      class="lime lighten-2"
      @click="submit"
    >
      登録
    </v-btn>
  </form>
</template>

<script>
import { required, maxLength, email, minLength } from 'vuelidate/lib/validators'
import * as types from "~/store/types";

export default {
  validations: {
    name: { required, maxLength: maxLength(10) },
    email: { required, email },
    password: { required, minLength: minLength(8) }
  },

  data() {
    return {
      name: '',
      email: '',
      password: ''
    }
  },

  computed: {
    nameErrors() {
      const errors = []
      if (!this.$v.name.$dirty) return errors
      !this.$v.name.maxLength && errors.push('ニックネームは10文字以内にしてください')
      !this.$v.name.required && errors.push('ニックネームは必須です')
      return errors
    },
    emailErrors() {
      const errors = []
      if (!this.$v.email.$dirty) return errors
      !this.$v.email.email && errors.push('適切なメールアドレスを入力してください')
      !this.$v.email.required && errors.push('メールアドレスは必須です')
      return errors
    },
    passwordErrors() {
      const errors = []
      if (!this.$v.password.$dirty) return errors
      !this.$v.password.minLength && errors.push('パスワードは8文字以上にしてください')
      !this.$v.password.required && errors.push('パスワードは必須です')
      return errors
    }
  },

  mounted() {
  },

  methods: {
    submit() {
      this.$v.$touch()
      if (this.$v.$invalid) return

      console.log('validate success')

      this.$axios
        .post(process.env.apiBaseUrl + '/users', {
          'name': this.name,
          'email': this.email,
          'password': this.password
        })
        .then((res) => {
          console.log(res)
          this.$store.dispatch(types.ACTIVATE_INFO_NOTIFICATION, {message: 'ユーザー「' + this.name + '」を登録しました。'});
          this.$router.push('/')
        })
        .catch((err) => {
          console.log(err)
          this.$store.dispatch(types.ACTIVATE_ERROR_NOTIFICATION, {message: err});
        })
    }
  }
}
</script>
