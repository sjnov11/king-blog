<template>

  <div class="post">   
    <div class="post-html" v-html="post">      
    </div>
    <div class="loading" v-if="loading">
      Loading...
    </div>

    <div class="error" v-if="error">
      There is some problem. {{ error }}
    </div>
  </div>
</template>

<script>
import PageTitle from "./Title.vue";
export default {
  name: "Post",
  props: ["post-meta"],
  components: PageTitle,
  data () {
    return {
      loading: false,
      post: "",
      error: null
    }
  },
  created () {
    this.fetchPost()
  },
  watch: {
    '$route' () {
      this.fetchPost();            
    }
  },
  methods: {
    fetchPost : function () {
      this.error = this.post = null;
      this.loading = true;
    
      const path = "/blog/" + this.$route.params.post +".html";
      fetch(path)
        .then(response => response.text())
        .then(data => {
          this.post = data;
          this.loading = false;
          // Do mathjax rendering when all DOM is loaded
          this.$nextTick().then(()=> {
            window.MathJax.Hub.Queue(["Typeset", window.MathJax.Hub]);
          });
        })
        .catch( (err) => { this.error = err } )
    }
  }
};
</script>

<style scoped>
.loading {
  text-align: center;
}

.post {
  width: 80%;
  margin: auto;
  text-align: center;
}
.post >>> .post-header {
  text-align: center;
}
.post >>> .body {
  text-align: left;
}


.post >>> .body img {
  width: 75%;
  height: 75%;
  margin: auto;
}

img {
  
}
</style>
