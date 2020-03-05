<template>
  <div class="auth">
    <el-header>
      <div>
        <span>お試し機能をご利用したい場合は、</span>
        <el-link href="https://element.eleme.io" target="_blank">こちらからゲストユーザーでログインしてください。</el-link>
      </div>
    </el-header>
    <el-main>
      <div v-if="!isAuthenticated">
        <amplify-authenticator></amplify-authenticator>
      </div>
      <div v-if="isAuthenticated">
        <amplify-sign-out></amplify-sign-out>
      </div>
    </el-main>
  </div>
</template>

<script>
export default {
  name: "auth",
  data() {
    return {};
  },
  beforeCreate() {
    this.$store.dispatch("login");
    this.$store.dispatch("loginCheck");
  },
  computed: {
    isAuthenticated() {
      return this.$store.getters.signIn;
    }
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
