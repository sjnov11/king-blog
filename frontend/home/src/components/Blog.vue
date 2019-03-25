<template>
  <div>    
    <div class="post-list" v-if="!this.$route.params.post">
      <ul v-for="post in postMetaData.index" :key="post.title">
        <router-link class="post-title header" :to="post.uri">{{post.title}}</router-link>
        <p>{{post.date}}</p>
        <div class="post-tags">
          <router-link to="/" class="post-tag" v-for="tag in post.tags" :key="tag">{{tag}}</router-link>
        </div>
      </ul>
    </div>
    <div>{{postMetaData}}</div>
    <router-view :post-meta="postMetaData"/>
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
.title-container h1 {
  font-size: 50px;
  margin-bottom: 10px;
  font-weight: 800;
}
@media screen and (min-width: 768px) {
  .title-container h1 {
    font-size: 70px;
  }
}
.title {
  font-size: 30px;
  margin-bottom: 50px;
  font-weight: 600;

}
.post-list {
  text-align: left; 
  padding: 20px;   
}
.post-list ul {
  padding: 0 0 20px 0;
  margin-bottom: 30px;
  border-bottom:1px solid #eaeaea;
}
.post-list ul .post-title {
  font-weight: 600;
  font-size: 25px;
}
</style>
<style>
.post-tags .post-tag {
  margin-right: 5px;
  padding: 2px 8px;
  background-color: #fafafa;
}
</style>