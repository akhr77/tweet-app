<template>
  <div id="logout">
    <el-container>
      <el-main>
        <amplify-sign-out v-bind:signOutConfig="signOutConfig"></amplify-sign-out>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "logout",
  data() {
    return {};
  }
};
</script>
