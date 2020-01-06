<template>
  <div class="ProfilForm">
    <h2>Mon Profil</h2>
    <b-alert variant="success" :show="show">Informations updated successfully</b-alert>

    <form class="profil" @submit.prevent="update">
      <div class="form-group">
        <label for="exampleInputEmail1">Email address</label>
        <input required v-model="user.email" type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" >
      </div>

        <div class="form-group">
        <label for="lastname">Nom</label>
        <input required v-model="user.lastname" type="text" class="form-control" id="lastname">
      </div>

        <div class="form-group">
        <label for="firstname">Prénom</label>
        <input required v-model="user.firstname" type="text" class="form-control" id="firstname">
      </div>

        <div class="form-group">
        <label for="newPassword">Nouveau Password</label>
        <input required v-model="user.password" type="password" class="form-control" id="newPassword" placeholder="Password">
      </div>

      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'ProfilForm',

	data : function() {
    return {
      user: {
        id: '',
        email: '',
        lastname: '',
        firstname: '',
        password: ''
      },
      success: false,
      show: false
    }
  },

  created() {
    const { id, email, firstname, lastname } = JSON.parse(localStorage.getItem('user'));
    this.user = {
      id,
      email,
      firstname,
      lastname,
      password: ''
    };
  },

	methods: {
    update: function () {
      this.$store.dispatch('updateUser', this.user)
      .then(() => {
          this.show = true
      })
      .catch(err => console.log(err))
    }
	}
}
</script>

<style>
.ProfilForm{
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
