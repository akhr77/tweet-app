<template>
  <div id="list">
    <div class="container">
      <div class="tab">
        <b-tabs type="is-toggle" expanded>
          <b-tab-item label="新着"></b-tab-item>
          <b-tab-item label="フォロー中"></b-tab-item>
          <b-tab-item label="人気"></b-tab-item>
        </b-tabs>
      </div>
      <div class="columns is-mobile">
        <div v-for="info in postInfos" :key="info.ID">
          <div class="column is-3bydesktop">
            <div class="card">
              <div class="card-image">
                <figure class="image is-4by3">
                  <img :src="info.base64image" alt="Placeholder image" />
                </figure>
              </div>
              <div class="card-content">
                <div class="media">
                  <div class="media-left">
                    <figure class="image is-48x48">
                      <img
                        src="https://bulma.io/images/placeholders/96x96.png"
                        alt="Placeholder image"
                      />
                    </figure>
                  </div>
                  <div class="media-content">
                    <p class="title is-4">{{ "@" + info.username }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Vue from "vue";
export default {
  data() {
    return {
      postInfos: [],
      postInfo: {
        id: "",
        userName: "",
        email: "",
        avater: "",
        image: "",
        base64image: ""
      }
    };
  },
  created() {
    this.getUserPostInfo();
  },
  methods: {
    async getUserPostInfo() {
      const response = await this.$axios.get("http://localhost/api/post");
      this.downloadS3(response.data);
    },

    async downloadS3(userPosts) {
      for (const userPost of userPosts) {
        const response = await this.$axios.get("http://localhost/api/image", {
          params: { image: userPost.Image }
        });
        this.setBase64image(userPost, response.data.image);
      }
    },

    setUserId(userId) {
      this.$store.dispatch("userId", userId);
    },

    setAvater(avater) {
      this.$store.dispatch("avater", avater);
    },

    setBase64image(userPost, base64image) {
      this.postInfo.id = userPost.ID;
      this.postInfo.username = userPost.Username;
      this.postInfo.email = userPost.Email;
      this.postInfo.avater = userPost.Avater;
      this.postInfo.image = userPost.Image;
      this.postInfo.base64image = base64image;
      console.log(this.$store.getters.email);
      console.log(userPost.Email);
      if (this.$store.getters.email === userPost.Email) {
        this.setUserId(this.postInfo.id);
        this.setAvater(this.postInfo.avater);
      }
      this.postInfos.push(Vue.util.extend({}, this.postInfo));
    }
  }
};
</script>

<style scoped>
.column {
  padding: 2rem;
}

.is-3bydesktop {
  min-width: 430px;
  max-width: 700px;
}

.is-mobile {
  flex-wrap: wrap;
  margin: 0px 1em;
}

.tab {
  margin-left: auto;
  margin-right: auto;
  width: 80%;
}
</style>
