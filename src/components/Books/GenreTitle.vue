<template>
    <div class="row mt-3 justify-content-start">
        <div class="col-lg-12 col-md-12 col-sm-12 col-xs-12">
            <div class="book-header-rect"></div>
            <h1 class="book-genre-title ml-2">Best sellers</h1>
        </div>
    </div>
</template>

<script>
export default {
    name: "GenreTitle"
}
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