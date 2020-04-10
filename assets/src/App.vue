<template>
  <div id="app">
    <div class="container">
      <el-header>
        <router-link to="list">
          <span class="header-brand">FavPic</span>
        </router-link>
        <router-link to="search">
          <i class="header-search fas fa-search fa-lg"></i>
        </router-link>
        <router-link to="users">
          <i class="header-users fas fa-users fa-lg"></i>
        </router-link>
        <router-link to="logout">
          <i class="header-logout fas fa-sign-out-alt fa-lg"></i>
        </router-link>
        <router-link to="post">
          <i class="header-post fas fa-plus fa-lg"></i>
        </router-link>
        <router-link to="mypage">
          <i class="header-mypage fas fa-user-cog fa-lg"></i>
        </router-link>
      </el-header>
    </div>

    <router-view />
  </div>
</template>

<script>
export default {
  name: "App"
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}

.header-brand {
  margin: 0px 12px;
  font-weight: 900;
  float: left;
  font-size: 24px;
  color: black;
}

.header-logout,
.el-icon-edit {
  margin: 6px 12px;
  float: right;
}

.header-search,
.header-users,
.header-post {
  margin: 6px 12px;
  float: left;
  color: black;
}

.header-mypage {
  margin: 6px 12px;
  float: right;
  color: black;
}
</style>
