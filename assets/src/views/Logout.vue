<template>
  <div id="logout">
    <amplify-sign-out v-bind:signOutConfig="signOutConfig"></amplify-sign-out>
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
