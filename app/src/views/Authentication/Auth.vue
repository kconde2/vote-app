<template>
  <div class="auth-wrapper">
    <div class="auth-wrapper__content shadow-sm">
      <router-view></router-view>
    </div>
  </div>
</template>

<style lang="scss">
.auth-wrapper {
  display: flex;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  margin-top: 40px;

  &__content {
    background: white;
    max-width: 670px;
    width: 100%;
  }

  &__header {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
    align-items: center;
    align-content: center;
    justify-content: center;
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    height: 200px;
    color: white;
    position: relative;
    z-index: 1;
    font-size: 1.5em;
    font-weight: bold;
    text-transform: uppercase;

    &--login {
      background-image: url("../../assets/images/vote.jpg");
    }

    &--register {
      background-image: url("../../assets/images/register-to-vote.jpg");
    }

    &::before {
      content: "";
      display: block;
      position: absolute;
      z-index: -1;
      width: 100%;
      height: 100%;
      top: 0;
      left: 0;
      background-color: rgba(19, 109, 97, 0.6);
    }
  }

  &__form {
    padding: 30px;
    min-height: 300px;
  }
}
</style>
