<template>
  <div id="home" class="d-flex flex-column">
            <div class="body-wrapper">
                <div class="row mb-4" style="height: 5rem;"></div>
                <Carousel/>
                <Genre/>
                <Genre/>
                <Genre/>
            </div>
        </div>
</template>

<script>
  import Carousel from "@/components/widgets/Carousel.vue"
  import Genre from '@/components/Books/Genre.vue'

export default {
  name: 'Home',
  components: {
      Carousel,
      Genre
  }
}
</script>

<style scoped>
#home {
    background-color: #000000;
    height: 100vh;
    width: 100%;
    overflow-x: hidden;
}

.body-wrapper {
    width:50%; 
    margin-left: auto;
    margin-right: auto;
}
</style>
