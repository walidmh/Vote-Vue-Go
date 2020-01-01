<template>
  <div id="app">
    <NavMenu />
    <router-view></router-view>
  </div>
</template>

<script>
  import NavMenu from "./components/Header/NavMenu"

  export default {
    name: 'App',
    data() {
      return {
      }
    },
    components:{
      NavMenu
    },
  }
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  background-color : #2C2F33;
  color: #2c3e50;
  margin : 0px;
  padding : 0px;
  height : 100%;
}
</style>
