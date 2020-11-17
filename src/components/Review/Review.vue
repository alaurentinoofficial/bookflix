<template>
  <div class="book-details-content mt-3 ml-3">
    <div class="mt-3">
      <img
        src="@/assets/userImage.jpg"
        alt="Imagem do usuário"
        class="img-profile-comment rounded-circle mr-3 mb-1 ml-3 review-photo"
      />
    </div>

    <div class="book-details-content mt-3 ml-3" style="text-align: justify">
      <h1 class="book-title">Anderson Laurentino</h1>
      <div class="mt-3">
        <em class="fa fa-star book-rate-star-comment"></em>
        <span>4.7</span><span class="review-star-span">/5</span>
      </div>
      <p class="mt-2 review-details">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam eu
        aliquet enim. Pellentesque egestas consequat nibh, a aliquam enim
        ultricies id. Nam tincidunt felis urna, vitae porttitor lacus accumsan
        eu. Phasellus iaculis mauris et orci interdum bibendum. Aliquam ut leo
        augue. Maecenas efficitur nunc in erat viverra, sit amet malesuada odio
        rhoncus. Proin dignissim efficitur libero, a faucibus lacus ornare nec.
        neque iaculis ligula, at viverra quam ante et diam. Vestibulum eu
        euismod justo, sed molestie metus. Nulla tristique leo nec ipsum viverra
        eleifend.
      </p>
    </div>
  </div>
</template>

<script>
export default {};
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