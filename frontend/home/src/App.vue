<template>
  <div id="app">
    <NavBar/>
    <div>
      <p>
        <router-link to="/">About</router-link>&nbsp;
        <router-link to="/blog">Blog</router-link>&nbsp;
      </p>
      <p>
        <router-view></router-view>
      </p>
    </div>

    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import NavBar from "./components/NavBar.vue";
export default {
  name: "app",
  components: {
    HelloWorld,
    NavBar
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}
</style>
