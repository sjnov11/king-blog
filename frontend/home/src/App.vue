<template>
  <div id="app">
    <NavBar/>
    <PageTitle/>

    <div class="main container">
      <router-view/>
    </div>
  </div>
</template>

<script>
import NavBar from "./components/NavBar.vue";
import PageTitle from "./components/Title.vue";
export default {
  name: "app",
  components: {
    NavBar,
    PageTitle
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;

  margin-top: 0px;
}
</style>
