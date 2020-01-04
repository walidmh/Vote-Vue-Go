<template>
  <div class="LoginForm">
    <h2>Login</h2>
    <div class="alert alert-danger" role="alert" v-if="errors!=0">
      {{ errors.data.error }}
    </div>
    <FormComponent
      :initialValues="[
        {
          group: {
            label: 'Email',
            for: 'email',
            description: 'We\'ll never share your email with anyone else.'
          },
          id: 'email',
          type: 'email',
          required: 'required',
          placeholder: 'Email...'
        },
        {
          group: {
            label: 'Password',
            for: 'password'
          },
          id: 'password',
          type: 'password',
          required: 'required',
          placeholder: ''
        }
       ]"
       :onSubmit="login"
    />
  </div>
</template>

<script>
import FormComponent from './FormComponent';
export default {
  name: 'LoginForm',
  components: {
    FormComponent
  },
  data() {
    return {
      errors : 0
    }
  },
  methods: {
    login: function({ values }) {
      let email = values.email
      let password = values.password
      this.$store.dispatch('login', { email, password })
      .then(() => {
        this.$store.dispatch('getLoggedUser')
          .then(() => this.$router.push('/'))
          .catch(err => console.log(err))
      })
      .catch(err => this.errors = err.response)
    }
  }
}
</script>

<style>
.LoginForm{
  width : 35%;
  background-color : white;
  margin : auto;
  margin-top : 50px;
  padding-left : 20%;
  padding-right : 20%;
  border-radius : 30px;
  padding : 30px;
  box-shadow: 5px 5px 5px black;
}
</style>