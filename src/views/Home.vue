<template>
  <div id="home">
    <div class="body-wrapper">
      <div class="row mb-4" style="height: 5rem"></div>
      <Carousel />
      <template v-for="genre in genres">
        <Genre :genre="genre" />
      </template>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

import Carousel from "@/components/widgets/Carousel.vue";
import Genre from "@/components/Books/Genre.vue";

export default {
  name: "Home",
  components: {
    Carousel,
    Genre,
  },
  computed: {
    ...mapState(["books"])
  },
  created() {
    this.$store.dispatch("getBooks", { self: this });
  },
  data() {
    return {
      genres: [
        {
          id: 0,
          Title: "Best Sellers",
          books: [ 
            {
              "Rate": 4.8,
              "Title": "Antifrágil Part II",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": "5718382d-1868-4594-94c9-0fc550b1dbd0"
            },
            {
              "Rate": 4.7,
              "Title": "Antifrágil Part I",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 1
            },
            {
              "Rate": 4.5,
              "Title": "Antifrágil",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 2
            },
            {
              "Rate": 4.2,
              "Title": "Antifrágil",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 3
            }
          ]
        },
        {
          id: 1,
          Title: "Horror",
          books: [ 
            {
              "Rate": 4.8,
              "Title": "Antifrágil Part II",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 0
            },
            {
              "Rate": 4.7,
              "Title": "Antifrágil Part I",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 1
            },
            {
              "Rate": 4.5,
              "Title": "Antifrágil",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 2
            },
            {
              "Rate": 4.2,
              "Title": "Antifrágil",
              "ImageUrl": "https://trello-attachments.s3.amazonaws.com/5fab02b43fe0e3846cc6d2a8/5fb55d0810c9400ce41597c3/24eddd7f878b5e59ef417e80419546d6/bookImage1.jpg",
              "id": 3
            }
          ]
        }
      ]
    }
  }
};
</script>

<style scoped>
#home {
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
