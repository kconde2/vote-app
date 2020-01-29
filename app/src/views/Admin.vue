<template>
  <fragment>
    <Header></Header>
    <div class="container main">
      <router-view></router-view>
    </div>
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
.main {
  min-height: calc(100vh - 165px);
}
</style>
