<template>
  <div id="home">
    <div class="body-wrapper">
      <div class="row mb-4" style="height: 10rem"></div>
      <template v-for="genre in genresFilter">
        <Genre :genre="genre" :isToSeeMore="true" />
      </template>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Genre from "@/components/Books/Genre.vue";

export default {
  props: ["genreId"],
  computed: {
    ...mapState(["genres"])
  },
  components: {
    Genre
  },
  created() {
      this.genresFilter = this.genres.filter((element) => element.id == this.genreId);
  },
  data() {
      return {
          genresFilter: []
      }
  }
};
</script>

<style scoped>
#home {
  background-color: #000000;
  height: 100vh;
  width: 100%;
  overflow-x: hidden;
}

.body-wrapper {
  width: 50%;
  margin-left: auto;
  margin-right: auto;
}
</style>