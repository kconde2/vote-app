<template>
  <div class="container">
    <div class="row">
      <div class="col-12 d-flex mb-2">
        <div class="d-flex align-items-center">
          <h1 class="h3">Liste des utilisateurs</h1>
        </div>

        <div class="ml-auto d-flex align-items-center">
          <router-link
            :to="{ name: 'user-add'}"
            href="#"
            class="btn btn-primary"
          >Ajouter un utilisateur</router-link>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <div class="card border-0">
          <div class="card-body">
            <h5 class="card-title">Title</h5>
            <div class="card-text table-responsive">
              <table style="width:100%" class="table table-bordered">
                <tr>
                  <th>Firstname</th>
                  <th>Lastname</th>
                  <th>Age</th>
                </tr>
                <tr>
                  <td>Jill</td>
                  <td>Smith</td>
                  <td>50</td>
                </tr>
                <tr>
                  <td>Eve</td>
                  <td>Jackson</td>
                  <td>94</td>
                </tr>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style>
</style>
