<template>
    <div id="body-genre">
        <GenreTitle />
        <div class="row justify-content-between">
            <Book />
            <Book />
            <Book />
            <Book />
        </div>
    </div>
</template>

<script>
  import GenreTitle from '@/components/Books/GenreTitle.vue'
  import Book from '@/components/Books/Book.vue'


export default {
    name: "Genre",
    components: {
        GenreTitle,
        Book
    }
}
</script>

<style scoped>

</style>