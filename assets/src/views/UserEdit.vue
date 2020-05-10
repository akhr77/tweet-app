<template>
  <div id="user-edit">
    <div class="container">
      <div class="profile-header">プロフィールを編集</div>
      <div class="columns">
        <div class="column is-4">
          <img v-show="!resizedImg" :src="avaterData" alt="Avater" class="avater" />
          <img v-show="resizedImg" :src="resizedImg" alt="Avater" class="avater" />
          <canvas v-show="false" ref="canvas" class="resize-img__preview__canvas" />
          <label for="file" class="resize-img__post__label">
            <div class="edit-avater-label">
              <span>画像を編集</span>
            </div>
            <input id="file" ref="fileInput" type="file" accept=".jpeg, .png" @change="uploadImg" />
          </label>
        </div>
        <div class="column is-4">
          <div class="control">
            <div class="input-area">
              <input class="input is-medium" type="text" placeholder="名前" v-model="username" />
            </div>
            <div class="input-area">
              <textarea class="textarea is-medium" placeholder="自己紹介" v-model="userProfile"></textarea>
            </div>
            <b-button class="column__button" rounded @click="clear">キャンセル</b-button>
            <b-button class="column__button" rounded @click="post">変更を保存</b-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Post from "./Post";
export default {
  name: "UserEdit",
  created() {
    this.getUserInfo();
  },
  data() {
    return {
      username: "",
      avater: "",
      userProfile: "",
      avaterData: "",
      resizedImg: null
    };
  },
  mixins: [Post],
  methods: {
    getUserInfo() {
      const queries = {
        id: this.$store.getters.userId,
        avater: this.$store.getters.avater
      };
      this.$axios
        .get("http://localhost/api/user", { params: queries })
        .then(response => {
          console.log(response.data);
          this.username = response.data.Username;
          this.avater = response.data.Avater;
          this.userProfile = response.data.UserProfile;
          this.avaterData = response.data.Image;
        });
    },
    post() {
      let params = new URLSearchParams();
      params.append("id", this.$store.getters.userId);
      params.append("fileName", this.fileName);
      params.append("fileType", this.fileType);
      params.append("avater", this.imageData);
      params.append("username", this.username);
      params.append("userProfile", this.userProfile);
      this.$axios.post("http://localhost/api/user", params).then(response => {
        console.log(response);
        this.$router.push("mypage");
      });
    },
    clear() {
      this.$router.push("mypage");
    }
  }
};
</script>

<style scoped lang="scss">
.avater {
  vertical-align: middle;
  width: 40%;
  border-radius: 50%;
}
.column {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 1.2em;
  &__button {
    margin: 0 5px 5px 0px;
  }
  &__avater_edit {
    margin-top: 1.2em;
  }
}
.input-area {
  margin: 1.3em 0em;
  width: 100%;
}
.resize-img {
  &__post {
    &__label {
      .edit-avater-label {
        border: 1px solid #dbdbdb;
        border-radius: 290486px;
        box-shadow: inset 0 1px 2px rgba(10, 10, 10, 0.1);
        margin: 1.3em 0;
        padding: 0.5em;
      }
      & > input {
        display: none;
      }
    }
  }
}
</style>

