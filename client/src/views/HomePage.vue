<template>
  <div class="container HomePage">
    <br>

    <b-row>
      <b-col ols="6" md="4" v-for="poll in polls" :key="poll.id">
        <b-card
          :title="poll.title"
          tag="article"
          style="max-width: 20rem;"
          class="mb-2"
        >
          <b-card-text>
            {{ poll.desc }} <br>
            <small>{{ poll.users.length }} Votes</small> <br>

            <b-button pill
              v-if="poll.author_id === loggedUser.id"
              @click="initUpdate(poll)"
              variant="info"
            >
              Update
            </b-button>

            <b-button pill
              v-if="poll.author_id === loggedUser.id"
              @click="deletePoll(poll)"
              variant="danger"
            >
              Delete
            </b-button>
          </b-card-text>

          <b-button
            v-if="!userVoted(poll, loggedUser.id)"
            @click="addVote(poll.id)"
            href="#"
            variant="primary"
          >
            Vote
          </b-button>
          <b-button
            v-if="userVoted(poll, loggedUser.id)"
            @click="deleteVote(poll.id)"
            href="#"
            variant="danger"
          >
            delete Vote
          </b-button>

        </b-card>
      </b-col>
    </b-row>

    <b-modal
      id="modal"
      ref="modal"
      title="Update Poll"
      @ok="handleOk"
    >
      <form ref="form" @submit.stop.prevent="handleSubmit">
        <b-form-group
          label="Title"
          label-for="title-input"
          invalid-feedback="Title is required"
        >
          <b-form-input
            id="title-input"
            v-model="selectedPoll.title"
            required
          >
          </b-form-input>
        </b-form-group>

        <b-form-group
          label="Description"
          label-for="desc-input"
          invalid-feedback="Description is required"
        >
          <b-form-input
            id="desc-input"
            v-model="selectedPoll.desc"
            required
          >
          </b-form-input>
        </b-form-group>
      </form>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: 'HomePage',

  computed : {
    loggedUser : function() { return this.$store.getters.loggedUser }
  },

  data() {
    return {
      polls: [],
      selectedPoll: {
        title: '',
        desc: ''
      }
    }
  },

  created() {
    this.getPolls()
  },

  methods: {
    checkFormValidity() {
      const valid = this.$refs.form.checkValidity()
      this.nameState = valid ? 'valid' : 'invalid'
      return valid
    },

    handleOk(bvModalEvt) {
      // Prevent modal from closing
      bvModalEvt.preventDefault()
      // Trigger submit handler
      this.handleSubmit()
    },

    handleSubmit() {
      // Exit when the form isn't valid
      if (!this.checkFormValidity()) {
        return
      }

      // Update the poll
      this.$store.dispatch('updatePoll', this.selectedPoll)
      .then(() => {})
      .catch((err) => console.log(err))

      // Hide the modal manually
      this.$nextTick(() => {
        this.$refs.modal.hide()
      })
    },

    initUpdate(poll) {
      this.selectedPoll = poll;
      this.$refs.modal.show()
    },

    deletePoll(poll) {
      this.$store.dispatch('deletePoll', poll.id)
      .then(() => this.getPolls())
      .catch((err) => console.log(err))
    },

    getPolls() {
      this.$store.dispatch('getPolls')
      .then((resp) => {
        this.polls = resp.data.map(poll => ({ ...poll, users: poll.users || [] }))
      })
      .catch((err) => console.log(err))
    },

    addVote(id) {
      this.$store.dispatch('addVote', id)
      .then(() => this.getPolls())
      .catch((err) => console.log(err))
    },
    deleteVote(id) {
      this.$store.dispatch('deleteVote', id)
      .then(() => this.getPolls())
      .catch((err) => console.log(err))
    },


    userVoted(poll, userId) {
      let result = false

      poll.users.forEach(user => {
        if (user.id === userId)
          result = true
      })

      return result
    }
  }
}
</script>
