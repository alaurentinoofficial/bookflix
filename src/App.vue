<template>
  <div id="app">
    <Topbar v-if="logginState" />
    <router-view id="content-wrapper" class="d-flex flex-column" />
  </div>
</template>

<script>
import Topbar from "@/components/Template/Topbar.vue";
// import Footer from "@/components/Template/Footer.vue";

export default {
  name: "App",
  components: {
    Topbar,
  },
  computed: {
    logginState() { 
      return this.$store.state.logginState; 
    }
  }
};
</script>
<style>
#app {
  margin: 0;
  padding: 0;
  overflow: hidden;
}
#content-wrapper {
  background-color: #000000;
  display: block;
  padding-bottom: 100px;
  /* height: 100vh; */
  min-height: 100vh;
  width: 100%;
  overflow-x: unset;
}
</style>
