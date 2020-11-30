<template>
  <div>
    <div class="book-details-content mt-3 ml-3">
      <div class="book-details-content mt-3 ml-3" style="text-align: justify">
        <h4 class="book-title">{{ reviewUser }}</h4>
        <div class="mt-3">
          <em class="fa fa-star book-rate-star-comment"></em>
          <span>{{ reviewRate.toFixed(2) }}</span><span class="review-star-span">/5</span>
        </div>
        <p class="mt-2 review-details">{{ reviewComment }}</p>
      </div>
  </div>
  <br />
</div>
  
</template>

<script>
export default {
  props: ["reviewUser", "reviewRate", "reviewComment", "imageUrl" ],
  name: "Review"
};
</script>

<style scoped>
.book-rate {
  font-weight: bold;
  color: white;
  top: -125%;
  padding-left: 10%;
  padding-right: 10%;
  width: 100%;
  position: relative;
}

.book-rate-star-comment {
  color: #f5c30b;
  font-size: 1rem;
}

.img-profile-comment {
  height: 4rem;
  width: 4rem;
}

.img-profile {
  height: 2rem;
  width: 2rem;
}

.review-star-span {
  color: #676565;
}

.review-details {
  font-weight: 100;
  font-size: medium;
}

.review-photo {
  float: left;
}
</style>