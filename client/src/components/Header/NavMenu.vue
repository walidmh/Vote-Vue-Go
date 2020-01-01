<template>
  <b-navbar toggleable="md" type="dark" variant="dark">
    <b-container>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item v-if="!isLoggedIn">
            <b-link to="/login">Login</b-link>
          </b-nav-item>

          <b-nav-item v-if="!isLoggedIn">
            <b-link to="/signup">Inscription</b-link>
          </b-nav-item>

          <b-nav-item v-if="isLoggedIn">
            <b-link to="/">Home</b-link>
          </b-nav-item>

          <b-nav-item v-if="isLoggedIn">
            <b-link to="/createPoll">Create Poll</b-link>
          </b-nav-item>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav v-if="isLoggedIn" class="ml-auto">
          <b-nav-item-dropdown right>
            <!-- Using 'button-content' slot -->
            <template v-if="loggedUser" v-slot:button-content>
                <em>{{ loggedUser.firstname }} {{ loggedUser.lastname }}</em>
            </template>

            <b-dropdown-item>
              <b-link to="/profil">Profil</b-link>
            </b-dropdown-item>

            <b-dropdown-item v-on:click="logout()">
              <b-link to="/login">Logout</b-link>
            </b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-container>
  </b-navbar>
</template>

<script>

export default {
  name: 'NavMenu',
  computed : {
    isLoggedIn : function() { return this.$store.getters.isLoggedIn },
    loggedUser : function() { return this.$store.getters.loggedUser }
  },
  methods: {
    logout: function () {
      this.$store.dispatch('logout')
      .then(() => this.$router.push('/login'))
    }
  },
}
</script>

<style>

</style>
