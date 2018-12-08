<template>
  <div class="navbar-container">
    <b-navbar class="navbar-custom" toggleable="lg" type="light" variant="light">
      <b-navbar-brand to="/">SEUNGJUN OH</b-navbar-brand>
      <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>

      <b-collapse is-nav id="nav_collapse">
        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-item to="/blog">Blog</b-nav-item>
          <b-nav-item to="/about">About</b-nav-item>
          <b-nav-form class="navbar-search">
            <b-form-input size="sm" class="mr-sm-2" type="text" placeholder="Search"/>
            <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
          </b-nav-form>
          <br>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <!-- avatar img container -->
    <div class="avatar-container">
      <div class="avatar-img-border">
        <a title="Seungjun Oh" to="/">
          <img class="avatar-img" src="../assets/avatar.jpg" alt="Seungjun Oh">
        </a>
      </div>
    </div>
  </div>
</template>
  
<script>
export default {
  name: "NavBar",
  props: {}
};
</script>

<style scoped>
.navbar-container {
  margin-bottom: 80px;
}
.navbar-custom {
  padding: 25px 15px;
  margin-bottom: 20px;
  margin-right: auto;
  margin-left: auto;
  background: #f5f5f5 !important;
  border-bottom: 1px solid #eaeaea;
  font-weight: 800;
}

.navbar-brand:hover,
.navbar-brand:focus {
  color: rgba(0, 0, 0, 0.5) !important;
}
.navbar-search {
  padding-left: 20px;
  margin-left: auto;
  margin-right: auto;
}

.avatar-container {
  width: 100px;
  height: 100px;
  margin-top: -60px;
  opacity: 1;
  position: absolute;
  left: 50%;
}

.avatar-container .avatar-img-border {
  width: 100%;
  height: 100%;
  box-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
  -webkit-box-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
  -moz-box-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
  border-radius: 50%;
  margin-left: -100%;
  display: inline-block;
}

.avatar-container .avatar-img {
  width: 100%;
  border-radius: 50%;
  display: block;
}
img {
  display: block;
  margin: auto;
  max-width: 100%;
  max-height: 100%;
}
</style>
