<template>
  <div class="login">
    <el-header>
      <div>
        <span>お試し機能をご利用したい場合は、</span>
        <el-link href="https://element.eleme.io" target="_blank">こちらからゲストユーザーでログインしてください。</el-link>
      </div>
    </el-header>
    <el-main>
      <div v-if="!signedIn">
        <amplify-authenticator></amplify-authenticator>
      </div>
      <div v-if="signedIn">
        <amplify-sign-out></amplify-sign-out>
      </div>
    </el-main>
  </div>
</template>

<script>
import { AmplifyEventBus } from "aws-amplify-vue";
import { Auth } from "aws-amplify";
export default {
  name: "login",
  data() {
    return {
      signedIn: false
    };
  },
  async beforeCreate() {
    // 認証状態の設定
    try {
      const user = await Auth.currentAuthenticatedUser();
      const userSession = await Auth.currentSession();
      this.signedIn = true;
      this.$store.state.userName = user.username;
      this.$store.state.email = user.attributes.email;
      console.log("user=", user);
      console.log("userSession=", userSession);
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
