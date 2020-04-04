<template>
  <div id="post">
    <div class="columns is-desktop">
      <div class="column" v-if="imageData">
        <div class="buttuns">
          <b-button rounded>キャンセル</b-button>
          <b-button rounded @click="post">投稿</b-button>
        </div>
        <img class="image-preview" :src="imageData" />
      </div>
      <div class="column" v-else>
        <div class="image-form">
          <input type="file" accept="image/*" @change="onFileChange($event)" />
        </div>
      </div>
    </div>
    <div class="post-form-area">
      <i class="fas fa-camera fa-5x post-form"></i>
      <input type="file" accept="image/*" @change="onFileChange($event)" />
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      fileName: "",
      fileType: "",
      imageData: ""
    };
  },
  methods: {
    onFileChange(e) {
      const files = e.target.files;
      if (files.length > 0) {
        const file = files[0];
        const reader = new FileReader();
        this.fileName = file.name;
        this.fileType = file.type;
        reader.onload = e => {
          this.imageData = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },

    post() {
      let params = new URLSearchParams();
      params.append("fileName", this.fileName);
      params.append("fileType", this.fileType);
      params.append("image", this.imageData);
      this.$axios
        .post("http://localhost/api/uploads3", params)
        .then(response => {
          console.log(response);
          this.$router.push("list");
        });
    }
  }
};
</script>

<style scoped>
.image-form {
  max-width: 500px;
  height: 300px;
  background-color: red;
  border-radius: 1rem;
  margin-left: auto;
  margin-right: auto;
}

.image-preview {
  max-width: 500px;
}

.post-form-area {
  max-width: 550px;
  height: 300px;
  margin: 0 auto 1rem auto;
  width: 100%;
  background-color: #e9ecef;
}

.post-form {
  background-color: rgba(0, 0, 0, 0.1);
  height: 80%;
  width: 80%;
}
</style>

