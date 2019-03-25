<template>
  <div id="app">
    <div class="header">
      <NavBar/>
      <PageTitle title="king"/>
    </div>

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
  
}
#app .main a {
  color: #404040;
} 
#app .main a:hover {
  color: rgba(0, 0, 0, 0.5) !important;
  text-decoration-line: none;
}
#app .header {
  font-family: "Arial", sans-serif;
   -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #404040;
  margin-bottom: 70px;
}

</style>
