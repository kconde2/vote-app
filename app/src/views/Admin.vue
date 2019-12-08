<template>
  <fragment>
    <Header></Header>
    <router-view></router-view>
    <Footer></Footer>
  </fragment>
</template>

<script>
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";

export default {
  components: {
    Header,
    Footer
  }
};
</script>

<style>
</style>
