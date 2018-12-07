<template>
  <b-navbar class="navbar-custom" toggleable="md" type="light" variant="light">
    <b-navbar-brand to="#">SEUNGJUN OH</b-navbar-brand>
    <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>

    <b-collapse is-nav id="nav_collapse">
      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item href="#">Blog</b-nav-item>
        <b-nav-item href="#">About</b-nav-item>
        <b-nav-form class="navbar-search">
          <b-form-input size="sm" class="mr-sm-2" type="text" placeholder="Search"/>
          <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
        </b-nav-form>
      </b-navbar-nav>
    </b-collapse>

    <!-- avatar img container -->
    <div class="avatar-container">
      <div class="avatar-img-border" v-show="viewAvatar">
        <a title="Seungjun Oh" to="/">
          <img class="avatar-img" src="https://sookocheff.com/img/avatar.png" alt="Seungjun Oh">
        </a>
      </div>
    </div>
  </b-navbar>
</template>
  
<script>
export default {
  name: "NavBar",
  props: {},
  data() {
    return {
      viewAvatar: true
    };
  },
  methods: {
    hideAvatar() {
      this.viewAvatar = false;
      this.console.log("Test");
    }
  }
};
</script>

<style scoped>
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
  margin-left: auto;
  margin-right: auto;
}

.avatar-container {
  width: 100px;
  margin-top: 50px;
  opacity: 1;
  position: absolute;
  left: 50%;
}

.avatar-container .avatar-img-border {
  width: 100%;
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
}
</style>
