import Vue from "vue";
import Router from "vue-router";

import About from "./components/About.vue";
import Blog from "./components/Blog.vue";
import Post from "./components/Post.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "Home",
      component: About
    },
    {
      path: "/about",
      name: "About",
      component: About
    },
    {
      path: "/blog",
      name: "Blog",
      component: Blog,
      children: [        
        { path: ":post", component: Post }
      ]
    }
  ]
});
