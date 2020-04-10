<template>
  <div id="user-edit">
    <div class="container">
      <div class="profile-header">プロフィールを編集</div>
      <div class="columns">
        <div class="column is-4">
          <img src="https://bulma.io/images/placeholders/96x96.png" alt="Avatar" class="avatar" />
          <b-button class="column__button" rounded @click="editAvater">画像を編集</b-button>
        </div>
        <div class="column is-8"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "UserEdit",
  data() {
    return {};
  }
};
</script>

<style scoped lang="scss">
.avatar {
  vertical-align: middle;
  width: 40%;
  border-radius: 50%;
}
.column {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 1.2em;
  .column__button {
    margin-top: 1.2em;
  }
}
</style>

