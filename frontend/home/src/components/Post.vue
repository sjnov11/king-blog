<template>
  <div>    
    <div class="post" v-html="post">      
    </div>
    <div class="post">
    <div class="loading" v-if="loading">
      Loading...
    </div>

    <div class="error" v-if="error">
      There is some problem. {{ error }}
    </div>

    <div class="content" v-if="post">      
      {{ post }}
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name: "Post",
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
          // Do mathjax rendering when all DOM is loaded
          this.$nextTick().then(()=> {
            window.MathJax.Hub.Queue(["Typeset", window.MathJax.Hub]);
          });
        })
        .catch( (err) => { this.error = err })
    }
  }
};
</script>

<style scoped>

</style>
