<template>
  <div id="list">
    <div class="columns is-desktop">
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
                  <p class="title is-4">{{ info.userName }}</p>
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
      const response = await this.$axios.get("http://localhost/api/");
      this.downloadS3(response.data);
    },

    async downloadS3(userPosts) {
      for (const userPost of userPosts) {
        console.log("userPost=", userPost);
        const response = await this.$axios.get(
          "http://localhost/api/downloadS3",
          {
            params: { image: userPost.Image }
          }
        );
        this.setBase64image(userPost, response.data.image);
        console.log("response=", response);
      }
    },

    setBase64image(userPost, base64image) {
      this.postInfo.id = userPost.ID;
      this.postInfo.userName = userPost.UserName;
      this.postInfo.email = userPost.Email;
      this.postInfo.avater = userPost.Avater;
      this.postInfo.image = userPost.Image;
      this.postInfo.base64image = base64image;
      this.postInfos.push(Vue.util.extend({}, this.postInfo));
      console.log("set base64");
    }

    // format3items() {
    //   var baseArr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12];
    //   var b = baseArr.length,
    //     cnt = 3,
    //     newArr = [];

    //   for (var i = 0; i < Math.ceil(b / cnt); i++) {
    //     var j = i * cnt;
    //     var p = baseArr.slice(j, j + cnt);
    //     newArr.push(p);
    //   }
    // }
  }
};
</script>

<style scoped>
.column {
  padding: 2rem;
}

.is-3bydesktop {
  min-width: 380px;
}

.is-desktop {
  flex-wrap: wrap;
}
</style>
