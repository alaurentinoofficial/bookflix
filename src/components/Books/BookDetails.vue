<template>
  <div class="row justify-content-between">
    <div class="col-lg-3 col-md-4 col-sm-6 col-xs-12 mt-3">
      <img
        class="bookimage"
        :src="bookImageUrl"
        alt="Imagem do livro Antifragil"
      />
    </div>
    <div class="col-lg-9 col-md-8 col-sm-6 col-xs-12 mt-3 book-details">
      <div class="book-details-content mt-3 ml-3">
        <h1 class="book-title">{{ bookTitle }}</h1>
        <div class="mt-3">
          <em class="fa fa-star book-rate-star"></em>
          <span>{{ bookRate }}</span><span class="review-star-span">/5</span>
        </div>
        <p class="review-details mt-3">
          {{ bookDetails }}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: [ 'bookDetails', 'bookTitle', 'bookImageUrl', 'bookRate'],
  name: "BookDetails",
};
</script>

<style scoped>
.bookimage {
  border-radius: 3%;
  width: 100%;
  height: 100%;
}

.review-star-span {
  color: #676565;
}

.book-rate-star {
  color: #f5c30b;
  font-size: 1.5rem;
}

.review-details {
  font-weight: 100;
  font-size: medium;
}

.book-details {
  background-color: #232323;
  width: auto;
  border-radius: 1%;
}

.book-details {
  display: flex;
  flex-direction: column;
  color: white;
}
</style>