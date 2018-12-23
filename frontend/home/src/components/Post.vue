<template>
  <div>
    <div>{{this.$route.params.post}}</div>
    <div>Post!!!!</div>
    <div v-html="post">
      here:
    </div>
    </div>
</template>

<script>
//import MathJax from 'mathjax'
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
    // 뷰가 생성되고 데이터가 이미 감시 되고 있을 때 데이터를 가져온다.
    this.fetchPost()
  },
  watch: {
    // 라우트가 변경되면 메소드를 다시 호출됩니다.
    '$route' () {
      this.fetchPost();
      
      this.window.console.log("mm..");
    }
  },
  methods: {
    fetchPost : function () {
      const path = "/blog/" + this.$route.params.post +".html";
      fetch(path)
        .then(response => response.text())
        .then(data => (this.post = data))        
    }
  }
};
</script>

<style scoped>

</style>
