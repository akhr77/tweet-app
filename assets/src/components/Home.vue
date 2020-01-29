<template>
  <div id="home">
    <el-container>
      <el-header>
        <div>
          <span>お試しで機能をご利用したい場合は、</span>
          <el-link href="https://element.eleme.io" target="_blank">こちらからゲストユーザーでログインしてください。</el-link>
        </div>
      </el-header>
      <el-main>
        <div class="nav-menu">
          <a class="home" font-weight-bold href="/">
            <b>FavPic</b>
          </a>
          <el-row>
            <el-button round class="nav-login">
              <router-link to="/login">ログイン</router-link>
            </el-button>
          </el-row>
        </div>
        <!-- 写真イメージをCardで表示 -->
        <el-row>
          <el-col :span="6" v-for="(o, index) in 3" :key="o" :offset="index > 0 ? 3 : 0">
            <el-card :body-style="{ padding: '0px' }" shadow="hover">
              <img
                src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
                class="image"
              />
              <div style="padding: 14px;">
                <span>Yummy hamburger</span>
              </div>
              <el-row class="avatar">
                <el-col :span="12">
                  <div class="sub-title">@mineko</div>
                  <div class="avater-circle">
                    <div class="block">
                      <el-avatar :size="70" :src="circleUrl"></el-avatar>
                    </div>
                    <div class="block" v-for="size in sizeList" :key="size">
                      <el-avatar :size="size" :src="circleUrl"></el-avatar>
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
      circleUrl:
        "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
      //   squareUrl:
      //     "https://cube.elemecdn.com/9/c2/f0ee8a3c7c9638a54940382568c9dpng.png"
    };
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
  margin-top: 10px;
}
</style>
