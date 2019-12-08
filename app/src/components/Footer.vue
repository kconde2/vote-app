<template>
  <fragment>
    <div class="va-footer fixed-bottom">
      <div class="container d-flex va-footer__wrapper">
        <div class="va-footer__logo">{{ new Date().getFullYear() }}</div>

        <div class="va-footer__menu"></div>

        <div class="va-footer__account"></div>
      </div>
    </div>
  </fragment>
</template>

<script>
export default {};
</script>

<style lang="scss" scoped>
.va-footer {
  background-color: #fff;
  box-shadow: 0 0 30px rgba(28, 39, 60, 0.08);
  border-bottom: 1px solid rgba(28, 39, 60, 0.12);
  padding: 20px 5px;
}
</style>
