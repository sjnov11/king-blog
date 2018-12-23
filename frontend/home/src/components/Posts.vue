<template>
  <div>
    Posts
    <ul v-for="post in postMetaData.list" :key="post">
      <router-link :to="post.uri">{{post.title}}</router-link>
    </ul>
  </div>
</template>

<script>
export default {
  name: "Blog",
  created() {
    this.getPostMetaData();
  },
  data: function() {
    return {
      postMetaData: {}
    };
  },
  methods: {
    getPostMetaData: function() {
      const path = "/blog/list.json";
      fetch(path)
        .then(response => response.json())
        .then(data => (this.postMetaData = data));
    }
  }
};
</script>

<style scoped>
</style>
