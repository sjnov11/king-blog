import Vue from "vue";
import App from "./App.vue";
import router from "./router.js";

import BootStrapVue from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";

import { library } from "@fortawesome/fontawesome-svg-core";
import {
  faCoffee,
  faLaptop,
  faCircle,
  faUserGraduate,
  faHome,
  faFileAlt,
  faEnvelope
} from "@fortawesome/free-solid-svg-icons";

import { faGithub, faTwitter, faLinkedin } from "@fortawesome/free-brands-svg-icons";
import {
  FontAwesomeIcon,
  FontAwesomeLayers
} from "@fortawesome/vue-fontawesome";

Vue.use(BootStrapVue);
Vue.config.productionTip = false;

library.add(
  faCoffee,
  faLaptop,
  faCircle,
  faUserGraduate,
  faHome,
  faFileAlt,
  faEnvelope,
);
library.add(faGithub, faTwitter, faLinkedin);

Vue.component("font-awesome-icon", FontAwesomeIcon);
Vue.component("font-awesome-layers", FontAwesomeLayers);
Vue.config.productionTip = false;

new Vue({
  router,
  render: h => h(App)  
}).$mount("#app");
