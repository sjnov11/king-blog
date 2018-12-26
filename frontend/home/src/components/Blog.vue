<template>
  <div>
    <div v-if="!this.$route.params.post">
      <ul v-for="post in postMetaData.index" :key="post.title">
        <router-link :to="post.uri">{{post.title}}</router-link>
        <p>{{post.date}}</p>
      </ul>
    </div>
    <router-view></router-view>
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
