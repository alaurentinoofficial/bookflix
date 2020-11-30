<template>
  <div id="content-wrapper" class="d-flex flex-column">
    <div class="body-wrapper" v-if="bookToReview != null && typeof(bookToReview) != 'undefined' ">
      <div class="row mb-4" style="height: 5rem"></div>
      <BookDetails :bookTitle="bookToReview.title" :bookImageUrl="bookToReview.image_url" :bookRate="bookToReview.rating" :bookDetails="bookToReview.resume" />
      <ReviewForm :bookId="id" />
      <ReviewList :reviews="reviewList" />
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

import BookDetails from "@/components/Books/BookDetails.vue";
import ReviewForm from "@/components/Review/ReviewForm.vue";
import ReviewList from "@/components/Review/ReviewList.vue";

export default {
  props: ["id"],
  name: "Review",
  components: {
    BookDetails,
    ReviewForm,
    ReviewList,
  },
  computed: mapState({
    bookToReview: state => state.bookToReview,
    reviewList: state => state.reviews
  }),
  created() {
    this.$store.dispatch("getReviewsByBookId", { self: this, bookId: this.id });
    this.$store.dispatch("getBookById", { self: this, id: this.id });
  }
};
</script>

<style scoped>
#content-wrapper {
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