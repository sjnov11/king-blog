<template>
  <div class="title-container">
    <h1>SEUNGJUN OH</h1>
    <hr class="hr-small">
  </div>
</template>

<script>
export default {
  name: "PageTitle",
  prop: {
    title: String
  }
};
</script>

<style scoped>
.title-container {  
  margin-bottom: 50px;
}
.title-container h1 {
  font-size: 40px;
  margin-bottom: 10px;
  font-weight: 800;
}
@media screen and (min-width: 768px) {
  .title-container h1 {
    font-size: 70px;
  }
}

.hr-small {
  max-width: 150px;
  margin: 25px auto;
  border-width: 4px;
  border-color: inherit;
  border-radius: 3px;
}
</style>
