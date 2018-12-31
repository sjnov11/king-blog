<template>
  <div class="about">
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

      
        <div class="footer">
          <span class="icon">
            <a
              href="https://docs.google.com/document/d/1sNvgFoduVu6Hwt1E3Iq1Hk-MJSHUKuYysahcdnc4wdI/edit#"
              title="resume"
              target="_blank"
            >
              <font-awesome-layers class="fa-3x">
                <font-awesome-icon icon="circle"></font-awesome-icon>
                <font-awesome-icon icon="file-alt" class="fa-inverse" transform="shrink-7"></font-awesome-icon>
              </font-awesome-layers>
            </a>
          </span>
          <span class="icon">
            <a href="https://github.com/sjnov11" title="github" target="_blank">
              <font-awesome-layers class="fa-3x">
                <font-awesome-icon icon="circle"></font-awesome-icon>
                <font-awesome-icon :icon="['fab','github']" class="fa-inverse" transform="shrink-7"></font-awesome-icon>
              </font-awesome-layers>
            </a>
          </span>
          <span class="icon">
            <a href="mailto:sjnov11@gmail.com" title="mail">
              <font-awesome-layers class="fa-3x">
                <font-awesome-icon icon="circle"></font-awesome-icon>
                <font-awesome-icon icon="envelope" class="fa-inverse" transform="shrink-7"></font-awesome-icon>
              </font-awesome-layers>
            </a>
          </span>
        </div>
      
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
.footer {
  margin-top: 50px;
}
.footer a {
  color: #404040;
}
.footer a:hover {
  color: #008aff;
}
.footer .icon {
  margin-right: 22px;
}
</style>
