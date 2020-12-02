<template>
  <div id="home">
    <div class="body-wrapper">
      <div class="row mb-4" style="height: 5rem"></div>
      <Carousel />
      <template v-for="genre in genresMapped()">
        <Genre :genre="genre" :isToSeeMore="false" :isToSearch="false" />
      </template>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Carousel from "@/components/widgets/Carousel.vue";
import Genre from "@/components/Books/Genre.vue";

export default {
  name: "Home",
  components: {
    Carousel,
    Genre,
  },
  computed: {
    ...mapState(["genres"])
  },
  created() {
    this.$store.dispatch("getBooks", { self: this });
  },
  methods: {
    genresMapped() {
      let genresMapped = JSON.parse(JSON.stringify(this.genres.map((element) => {
        element.books = element.books.map((book) => {
          book.categories = undefined;
          return book;
        });
        return element;
      })));

      return genresMapped.map((element) => {
        element.books = element.books.slice(0, 4);
        return element;
      });
    }
  },
  data() {
    return {
    }
  }
};
</script>

<style scoped>
#home {
  background-color: #000000;
  width: 100%;
}

.body-wrapper {
  width: 50%;
  margin-left: auto;
  margin-right: auto;
}
</style>
