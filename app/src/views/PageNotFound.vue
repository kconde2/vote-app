<template>
  <div class="page-404">
    <h1>404</h1>
    <p>Oops! Aucun vote ne se trouve par là.</p>
    <a href="/" class="btn btn-primary">
      <i class="fas fa-arrow-left"></i> Retournez à la page initiale, c’est mieux.
    </a>
  </div>
</template>

<style lang="scss" scoped>
.page-404 {
  text-align: center;
}

.button {
  font-weight: 300;
  color: #fff;
  font-size: 1.2em;
  text-decoration: none;
  border: 1px solid #efefef;
  padding: 0.5em;
  border-radius: 3px;
  float: left;
  margin: 6em 0 0 -155px;
  left: 50%;
  position: relative;
  transition: all 0.3s linear;

  &:hover {
    background-color: #0059ff;
    color: #fff;
  }
}

p {
  font-size: 2em;
  text-align: center;
  font-weight: 100;
}

h1 {
  text-align: center;
  font-size: 15em;
  font-weight: 100;
}
</style>
