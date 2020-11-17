<template>
  <div id="app">
    <Topbar/>
    <router-view id="content-wrapper" class="d-flex flex-column" />
    <Footer/>
  </div>
</template>

<script>
  import Topbar from '@/components/Template/Topbar.vue'
  import Footer from '@/components/Template/Footer.vue'

  export default {
    name: 'App',
    components: {
      Topbar,
      Footer
    }
}
</script>

<style>
#content-wrapper {
    background-color: #000000;
    height: 100vh;
    width: 100%;
    overflow-x: hidden;
}
</style>
