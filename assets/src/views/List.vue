<template>
  <div id="list">
    <el-container>
      <el-header>
        <router-link to="login">
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
      <el-main>
        <div class="nav-menu"></div>
        <!-- 写真イメージをCardで表示 -->
        <el-row>
          <el-col
            class="card-col"
            :span="6"
            v-for="(o, index) in 3"
            :key="o"
            :offset="index > 0 ? 3 : 0"
          >
            <el-card :body-style="{ padding: '0px' }" shadow="hover">
              <img
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                class="image"
              />
              <div style="padding: 14px;">
                <span>{{ info[0].Email }}</span>
              </div>
              <el-row class="avatar">
                <el-col :span="12">
                  <div class="sub-title">{{ info[0].UserName }}</div>
                  <div class="avater-circle">
                    <div class="block">
                      <el-avatar :size="70" :src="info[0].Avater"></el-avatar>
                    </div>
                  </div>
                </el-col>
              </el-row>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      info: [
        {
          ID: "",
          UserName: "",
          Email: "",
          EncrypedPassword: "",
          Avater: "",
          Image: ""
        }
      ],
      image: {}
    };
  },
  mounted() {
    this.$axios.get("http://localhost/api/").then(response => {
      this.info = response.data;
    });

    this.$axios.get("http://localhost/api/downloadS3").then(response => {
      this.image = response.data;
    });
  }
};
</script>

<style scoped>
.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both;
}

.home {
  font-size: 3em;
  text-decoration: none;
  float: left;
}

.nav-login {
  font-size: 1.2em;
  float: right;
  margin-top: 12px;
}

.card-col {
  width: 27%;
  margin-left: 6%;
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
