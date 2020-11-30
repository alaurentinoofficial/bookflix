<template>
  <div id="home">
    <div class="body-wrapper">
      <div class="row mb-4" style="height: 10rem"></div>
      <div class="row mt-3 justify-content-start">
        <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
          <div class="book-header-rect"></div>
          <h1 class="book-genre-title ml-2" v-if="genresFilter.length > 0">
            Resultados de "{{ searchText }}"
          </h1>
          <h1 class="book-genre-title ml-2" v-else>
            Nenhum resultado encontrado para "{{ searchText }}"
          </h1>
        </div>
      </div>

      <template v-for="genre in genresFilter">
        <Genre :genre="genre" :isToSeeMore="true" :isToSearch="true" />
      </template>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Genre from "@/components/Books/Genre.vue";

export default {
  props: ["searchText"],
  components: {
    Genre,
  },
  computed: {
    ...mapState(["genres"]),
    genresFilter() {
      let genresFilter = [];
      let books = [];

      this.genres.forEach((genre) => {
        books.push(
          ...genre.books.filter(
            (book) =>
              book.title
                .trim()
                .toUpperCase()
                .includes(this.searchText.trim().toUpperCase()) ||
              book.resume
                .trim()
                .toUpperCase()
                .includes(this.searchText.trim().toUpperCase()) ||
              book.author
                .trim()
                .toUpperCase()
                .includes(this.searchText.trim().toUpperCase())
          )
        );
      });

      for (let index = 0; index < books.length; index++) {
        const element = books[index];
        if (books.filter((b) => b.id == element.id).length > 1) {
          books.splice(index, 1);
          index--;
        }
      }

      if (books.length > 0) {
        genresFilter.push({
          name: this.searchText,
          books: books,
        });
      }

      return genresFilter;
    },
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

.book-genre-title {
  color: white;
}
</style>