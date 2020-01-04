<template>
  <div class="container SignupForm">
    <b-form @submit="onSubmit" @reset="onReset">
      <b-form-group
        id="input-group-1"
        label="Email address:"
        label-for="input-1"
        description="We'll never share your email with anyone else."
      >
        <b-form-input
          id="input-1"
          v-model="user.email"
          type="email"
          required
          placeholder="Enter email"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-2" label="First name" label-for="input-2">
        <b-form-input
          id="input-2"
          v-model="user.firstName"
          required
          placeholder="Enter your first name"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-3" label="Last name" label-for="input-3">
        <b-form-input
          id="input-3"
          v-model="user.lastName"
          required
          placeholder="Enter your last name"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-4" label="Password" label-for="input-4">
        <b-form-input
          id="input-4"
          type="password"
          v-model="user.password"
          required
          placeholder="Enter your password"
        ></b-form-input>
      </b-form-group>

      <b-form-group id="input-group-5" label="Birthday" label-for="input-5">
        <datepicker v-model="user.date_of_birth"></datepicker>
      </b-form-group>

      <b-button class="b-button" type="submit" variant="primary">Submit</b-button>
      <b-button class="b-button" type="reset" variant="danger">Reset</b-button>
    </b-form>
  </div>
</template>

<script>
  import datepicker from 'vuejs-datepicker';
  export default {
    components: {
      datepicker
    },
    data() {
      return {
        user: {
          email: '',
          firstName: '',
          lastName:'',
          password:'',
          date_of_birth: new Date(),
          access_level: 2
        }
      }
    },
    methods: {
      onSubmit(evt) {
        evt.preventDefault()
        this.$store.dispatch('createUser', this.user)
        .then(() => {
          this.$router.push('/login');
        });
      },
      onReset(evt) {
        evt.preventDefault()
        // Reset our form values
        this.user.email = ''
        this.user.firstName = ''
        this.user.lastName = ''
        this.user.password = ''
        this.user.date = ''
      }
    }
  }
</script>

<style>
    .b-button{
    margin: 5%;
    }
    .SignupForm{
    width : 35%;
    background-color : white;
    margin-top : 50px;
    border-radius : 30px;
    padding : 30px;
    box-shadow: 5px 5px 5px black;
    }
</style>