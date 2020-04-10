<template>
  <div id="mypage">
    <div class="container">
      <div class="columns">
        <div class="column is-4">
          <img src="https://bulma.io/images/placeholders/96x96.png" alt="Avatar" class="avatar" />
        </div>
        <div class="column is-8">
          <div align="left" class="user-name">みねお</div>
          <div align="left" class="user-anotation">@mineo</div>
          <router-link to="useredit">
            <b-button class="column__button" rounded @click="clearUploadImg">プロフィールの変更</b-button>
          </router-link>
          <b-button class="column__button" rounded @click="post">ログアウト</b-button>
          <div align="left" class="profile"></div>
        </div>
      </div>
      <div class="body-list">
        <b-tabs v-model="activeTab" position="is-centered">
          <b-tab-item :label="postTab">Lorem ipsum dolor sit amet.</b-tab-item>

          <b-tab-item :label="followTab">
            Lorem
            <br />ipsum
            <br />dolor
            <br />sit
            <br />amet.
          </b-tab-item>

          <b-tab-item :label="followerTab">
            What light is light, if Silvia be not seen?
            <br />What joy is joy, if Silvia be not by—
            <br />Unless it be to think that she is by
            <br />And feed upon the shadow of perfection?
            <br />Except I be by Silvia in the night,
            <br />There is no music in the nightingale.
          </b-tab-item>

          <b-tab-item :label="likeTab">
            Nunc nec velit nec libero vestibulum eleifend.
            Curabitur pulvinar congue luctus.
            Nullam hendrerit iaculis augue vitae ornare.
            Maecenas vehicula pulvinar tellus, id sodales felis lobortis eget.
          </b-tab-item>
        </b-tabs>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "MyPage",
  computed: {
    postTab() {
      return "投稿 " + this.postCount;
    },
    followTab() {
      return "フォロー中 " + this.followCount;
    },
    followerTab() {
      return "フォロワー " + this.followerCount;
    },
    likeTab() {
      return "いいね " + this.likeCount;
    }
  },
  data() {
    return {
      postCount: 22,
      followCount: 10,
      followerCount: 12,
      likeCount: 100
    };
  }
};
</script>

<style scoped lang="scss">
.avatar {
  vertical-align: middle;
  width: 70%;
  border-radius: 50%;
}

.column {
  .user-name {
    margin-top: 0.5em;
    font-size: 3em;
  }
  .user-anotation {
    font-size: 1.5em;
  }
  &__button {
    margin: 1.2em 5px 5px 0px;
    float: left;
  }
}
</style>

