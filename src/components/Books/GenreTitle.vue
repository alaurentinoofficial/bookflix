<template>
  <div class="row mt-3 justify-content-start">
    <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
      <div class="book-header-rect"></div>
      <h1 class="book-genre-title ml-2">{{ genreTitle }}</h1>
      <router-link :to="'/booksgenre/' + genreId" class="float-right mt-1" v-if="isToSeeMore == false">ver mais</router-link>
    </div>
  </div>
</template>

<script>
export default {
  props: ["genreTitle", "genreId", "isToSeeMore"],
  name: "GenreTitle",
};
</script>

<style scoped>
.book-header-rect {
  display: inline-block;
  vertical-align: sub;
  background-color: #b7323b;
  width: 0.3rem;
  height: 80%;
  overflow-y: 1;
}

.book-genre-title {
  color: white;
  font-size: 1.2rem;
  font-weight: bolder;
  display: inline-block;
}
</style>