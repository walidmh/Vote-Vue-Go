<template>
  <div class="RegisterForm">
    <h2>Inscription</h2>
    <form class="register" @submit.prevent="register">
      <div class="form-group">
        <label for="exampleInputEmail1">Email address</label>
        <input required v-model="email" type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email">
        <small id="emailHelp" class="form-text text-muted">We'll never share your email with anyone else.</small>
      </div>
      <div class="form-group">
        <label for="exampleInput1">FirstName</label>
        <input required v-model="email" type="text" class="form-control" id="exampleInput1">
      </div>
      <div class="form-group">
        <label for="exampleInput2">LastName</label>
        <input required v-model="email" type="text" class="form-control" id="exampleInput2">
      </div>
      <div class="form-group">
        <label for="exampleInputPassword1">Password</label>
        <input required v-model="password" type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
      </div>
      <button type="submit" class="btn btn-primary">Submit</button>
    </form>
  </div>
</template>

<script>


export default {
  name: 'RegisterForm'
}
</script>

<style>
.LoginForm {
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
