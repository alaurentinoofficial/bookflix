<template>
  <div class="row justify-content-between">
    <div class="col-12 ml-3 mt-5 book-details">
      <div class="book-details-content mt-3 ml-3">
        <div class="row ml-auto mr-auto review-stars">
          <Star :rating="ratingReview" />
        </div>
        <div class="form-group mt-4 ml-auto mr-auto" style="width: 80%">
          <input class="form-control review-textarea" type="text" placeholder="Nome do usuário" v-on:focusout="userName = $event.target.value">
          <textarea
            class="form-control review-textarea mt-3"
            id="exampleFormControlTextarea1"
            rows="4"
            v-on:focusout="reviewComment = $event.target.value"
            placeholder="Escreva sua crítica"
          ></textarea>
          <div class="col-md-12 text-center">
            <button type="submit" class="btn btn-primary mt-3 review-button" v-on:click="SendReview">
              ENVIAR
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Star from "@/components/widgets/Stars"

export default {
  props: ["bookId"],
  name: "ReviewForm",
  components: {
    Star
  },
  data() {
    return {
      reviewComment: "",
      ratingReview: 0,
      userName: ""
    }
  },
  computed: {
    ...mapState(["starToReview"])
  },
  methods: {                                                                          
    inputFile(event) {
      this.imageUserUrl = window.URL.createObjectURL(event.target.files[0]);
    },
    SendReview() {
      this.$store.dispatch("addReview", {
        self: this, 
        review: {
          resume: this.reviewComment,
          rating: this.starToReview,
          book_id: this.bookId,
          user_name: this.userName,
        }
      });
    }
  }
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

.book-rate-star-review {
  color: #f5c30b;
  font-size: 2.5rem;
}

.review-stars {
  width: 45%;
  justify-content: space-between;
}

.review-textarea {
  color: white;
  background-color: #171717;
  border-color: #474747;
}

.review-button {
  background-color: #b7323b;
  border-color: #b7323b;
}
</style>