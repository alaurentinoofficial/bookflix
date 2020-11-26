<template>
  <div id="body-genre">
    <GenreTitle :bookTitle="genre.Title" />
    <div class="row justify-content-between">
      <template v-for="book in genre.books">
        <Book :bookRate="book.Rate" :bookTitle="book.Title" :bookIdUrl="book.id" :bookImageUrl="book.ImageUrl" />
      </template>
    </div>
  </div>
</template>

<script>
import GenreTitle from "@/components/Books/GenreTitle.vue";
import Book from "@/components/Books/Book.vue";

export default {
  props: ["genre"],
  name: "Genre",
  components: {
    GenreTitle,
    Book,
  }
};
</script>

<style scoped>
</style>