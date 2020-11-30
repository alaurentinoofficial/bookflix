<template>
  <div id="content-wrapper" class="d-flex flex-column">
    <div class="body-wrapper" v-if="bookToReview != null && typeof(bookToReview) != 'undefined' ">
      <div class="row mb-4" style="height: 5rem"></div>
      <BookDetails :bookTitle="bookToReview.title" :bookImageUrl="book.bookImageUrl" :bookRate="bookToReview.rating" :bookDetails="bookToReview.resume" />
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
    bookToReview: state => state.bookToReview.body,
    reviewList: state => state.reviews
  }),
  created() {
    this.$store.dispatch("getReviewsByBookId", { self: this, bookId: this.id });
    this.$store.dispatch("getBookById", { self: this, id: this.id });
  },
  data() {
    return {
      book: {
        "Title": "Antifrágil",
        "bookImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
        "Rate": 4.7,
        "Details": `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam eu
          aliquet enim. Pellentesque egestas consequat nibh, a aliquam enim
          ultricies id. Nam tincidunt felis urna, vitae porttitor lacus accumsan
          eu. Phasellus iaculis mauris et orci interdum bibendum. Aliquam ut leo
          augue. Maecenas efficitur nunc in erat viverra, sit amet malesuada
          odio rhoncus. Proin dignissim efficitur libero, a faucibus lacus
          ornare nec. neque iaculis ligula, at viverra quam ante et diam.
          Vestibulum eu euismod justo, sed molestie metus. Nulla tristique leo
          nec ipsum viverra eleifend.`
      },
      reviews: [
        {
          "User": "Anderson Laurentino",
          "Rate": 4.7,
          "Comment": `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam eu
            aliquet enim. Pellentesque egestas consequat nibh, a aliquam enim
            ultricies id. Nam tincidunt felis urna, vitae porttitor lacus accumsan
            eu. Phasellus iaculis mauris et orci interdum bibendum. Aliquam ut leo
            augue. Maecenas efficitur nunc in erat viverra, sit amet malesuada odio
            rhoncus. Proin dignissim efficitur libero, a faucibus lacus ornare nec.
            neque iaculis ligula, at viverra quam ante et diam. Vestibulum eu
            euismod justo, sed molestie metus. Nulla tristique leo nec ipsum viverra
            eleifend.`
        },
        {
          "User": "Helton Alves",
          "Rate": 4.5,
          "Comment": `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam eu
            aliquet enim. Pellentesque egestas consequat nibh, a aliquam enim
            ultricies id. Nam tincidunt felis urna, vitae porttitor lacus accumsan
            eu. Phasellus iaculis mauris et orci interdum bibendum. Aliquam ut leo
            augue. Maecenas efficitur nunc in erat viverra, sit amet malesuada odio
            rhoncus. Proin dignissim efficitur libero, a faucibus lacus ornare nec.
            neque iaculis ligula, at viverra quam ante et diam. Vestibulum eu
            euismod justo, sed molestie metus. Nulla tristique leo nec ipsum viverra
            eleifend.`
        }
      ]
    }
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