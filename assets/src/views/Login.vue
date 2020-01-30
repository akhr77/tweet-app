<template>
  <div id="login">
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
            <el-button round>
              <a href="https://element.eleme.io">ログイン</a>
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
                <div class="bottom clearfix">
                  <time class="time">{{ currentDate }}</time>
                  <el-button type="text" class="button">Operating</el-button>
                </div>
              </div>
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
      currentDate: new Date()
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
</style>
