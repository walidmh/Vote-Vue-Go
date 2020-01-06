<template>
  <div class="container PollCreationForm">
    <h2>Create a Poll</h2>
    <form class="poll" @submit.prevent="createPoll">
      <div class="form-group">
        <label for="Title">Title of Poll</label>
        <input required v-model="title" type="title" class="form-control" id="titlPoll" aria-describedby="emailHelp" placeholder="Entrer le nom du Poll">
      </div>
      <div class="form-group">
        <label for="description">Description</label>
        <input required v-model="description" type="Description" class="form-control" id="Description" placeholder="Description">
      </div>
      <button type="submit" class="btn btn-primary">Créer</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'PollCreationForm',

  data() {
    return {
      title : '',
      description : ''
    }
  },

  methods: {
    createPoll: function () {
      let title = this.title
      let description = this.description

      this.$store.dispatch('createPoll', {
        'Title': title,
        'Desc': description,
        'AuthorID': JSON.parse(localStorage.getItem('user')).id
      })
      .then(() => this.$router.push('/'))
      .catch(err => console.log(err))
    }
  }
}
</script>

<style>
.PollCreationForm {
  width : 35%;
  background-color : white;
  margin-top : 50px;
  border-radius : 30px;
  padding : 30px;
  box-shadow: 5px 5px 5px black;
}
</style>
