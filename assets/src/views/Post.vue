<template>
  <div id="post">
    <h1>新しい投稿</h1>
    <i aria-hidden="true" class="fa fa-plus"></i>
    <input type="file" accept="image/*" @change="onFileChange($event)" />
    <div class="image-preview-area">
      <img class="image-preview" :src="imageData" v-if="imageData" />
    </div>
    <el-button plain>キャンセル</el-button>
    <el-button plain @click="post">投稿</el-button>
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
        });
    }
  }
};
</script>

<style scoped>
@media screen and (min-width: 768px) {
  .image-preview {
    max-width: 870px;
  }
}
</style>

