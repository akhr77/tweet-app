<template>
  <div class="login">
    <el-header>
      <div>
        <span>お試し機能をご利用したい場合は、</span>
        <el-link href="https://element.eleme.io" target="_blank">こちらからゲストユーザーでログインしてください。</el-link>
      </div>
    </el-header>
    <div style="height: 200px">
      <el-image :src="src"></el-image>
    </div>
    <div v-if="!signedIn">
      <amplify-authenticator></amplify-authenticator>
    </div>
    <div v-if="signedIn">
      <amplify-sign-out></amplify-sign-out>
    </div>
  </div>
</template>

<script>
import { AmplifyEventBus } from "aws-amplify-vue";
import { Auth } from "aws-amplify";
export default {
  name: "login",
  data() {
    return {
      signedIn: false,
      src: "static/login_background.jpg"
    };
  },
  async beforeCreate() {
    // 認証状態の設定
    try {
      await Auth.currentAuthenticatedUser();
      this.signedIn = true;
      this.$router.push({ path: "list" });
    } catch (err) {
      this.signedIn = false;
    }
    AmplifyEventBus.$on("authState", info => {
      if (info === "signedIn") {
        this.signedIn = true;
      } else {
        this.signedIn = false;
      }
    });
  }
};
</script>
<style scoped>
.el-header {
  background-color: #fc8023;
  color: white;
  line-height: 60px;
}
</style>
