<template>
  <div id="body-genre">
    <GenreTitle :genreTitle="genre.name" :genreId="genre.id" :isToSeeMore="isToSeeMore" />
    <div class="row justify-content-between">
      <template v-for="book in genre.books">
        <Book :bookRate="book.rating" :bookTitle="book.title" :bookIdUrl="book.id" :bookImageUrl="book.image_url" />
      </template>
    </div>
  </div>
</template>

<script>
import GenreTitle from "@/components/Books/GenreTitle.vue";
import Book from "@/components/Books/Book.vue";

export default {
  props: ["genre", "isToSeeMore"],
  name: "Genre",
  components: {
    GenreTitle,
    Book,
  }
};
</script>

<style scoped>
</style>