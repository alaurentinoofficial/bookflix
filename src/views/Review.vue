<template>
  <div id="content-wrapper" class="d-flex flex-column">
    <div class="body-wrapper">
      <div class="row mb-4" style="height: 5rem"></div>
      <BookDetails />
      <ReviewForm />
      <ReviewList />
    </div>
  </div>
</template>

<script>
import BookDetails from "@/components/Books/BookDetails.vue";
import ReviewForm from "@/components/Review/ReviewForm.vue";
import ReviewList from "@/components/Review/ReviewList.vue";

export default {
  name: "Review",
  components: {
    BookDetails,
    ReviewForm,
    ReviewList,
  },
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