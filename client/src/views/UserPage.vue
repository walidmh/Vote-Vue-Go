<template>
  <div class="ProfilPage">
    <ProfilForm v-bind:user="loggedUser"/>
  </div>
</template>

<script>
import ProfilForm from "../components/ProfilForm"

export default {
  name: 'ProfilPage',
  computed : {
      loggedUser : function() { return this.$store.getters.loggedUser }
  },
  components:{
    ProfilForm
  }
}
</script>
