<template>
  <div class="row justify-content-between">
    <div class="col-12 ml-3 mt-5 book-details">
      <template v-for="review in reviews">
        <Review :reviewUser="review.user_name" :imageUrl="review.image_url" :reviewRate="review.rating" :reviewComment="review.resume" />
      </template>
    </div>
  </div>
</template>

<script>
import Review from "@/components/Review/Review.vue";

export default {
  props:['reviews'],
  name: "ReviewList",
  components: {
    Review,
  },
};
</script>

<style scoped>
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