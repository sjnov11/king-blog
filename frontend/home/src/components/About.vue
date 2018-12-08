<template>
  <div class="about-container">
    <img>
    <div class="profile">
      <div class="record">
        <span class="icon">
          <font-awesome-icon icon="laptop" class="fa-fw"></font-awesome-icon>
        </span>
        <span class="text">Software Engineer</span>
      </div>
      <div class="record">
        <span class="icon">
          <font-awesome-icon icon="user-graduate" class="fa-fw"></font-awesome-icon>
        </span>
        <span class="text">Hanyang Univ.</span>
      </div>
      <div class="record">
        <span class="icon">
          <font-awesome-icon icon="home" class="fa-fw"></font-awesome-icon>
        </span>
        <span class="text">Seoul, South Korea</span>
      </div>

      <footer>
        <div class="footer-container">
          <span class="icon">
            <font-awesome-layers class="fa-3x">
              <font-awesome-icon icon="circle"></font-awesome-icon>
              <font-awesome-icon icon="file-alt" class="fa-inverse" transform="shrink-7"></font-awesome-icon>
            </font-awesome-layers>
          </span>
          <span class="icon">
            <font-awesome-layers class="fa-3x">
              <font-awesome-icon icon="circle"></font-awesome-icon>
              <font-awesome-icon :icon="['fab','github']" class="fa-inverse" transform="shrink-7"></font-awesome-icon>
            </font-awesome-layers>
          </span>
          <span class="icon">
            <font-awesome-layers class="fa-3x">
              <font-awesome-icon icon="circle"></font-awesome-icon>
              <font-awesome-icon icon="envelope" class="fa-inverse" transform="shrink-7"></font-awesome-icon>
            </font-awesome-layers>
          </span>
        </div>
      </footer>
    </div>
  </div>
</template>

<script>
export default {
  name: "About"
};
</script>

<style scoped>
.profile {
  width: 220px;
  margin: auto;
  text-align: left;
  padding-left: 9px;
}
.profile .record {
  font-size: 18px;
}
.profile .record .icon {
  width: 30px;
  margin-right: 12px;
  text-align: center;
}
footer {
  margin-top: 50px;
}
footer .icon {
  margin-right: 22px;
}
</style>
