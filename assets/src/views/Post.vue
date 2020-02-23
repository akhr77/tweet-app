<template>
  <div id="post">
    <el-container>
      <el-main>
        <h1>新しい投稿</h1>

        <div class="card" style="width: 20rem;margin:10px;">
          <div class="card-body">
            <input type="file" accept="image/*" @change="onFileChange($event)" />
          </div>
          <div class="image-preview">
            <img :src="imageData" v-if="imageData" />
          </div>
          <el-row>
            <el-button plain>キャンセル</el-button>
            <el-button plain @click="post">投稿</el-button>
          </el-row>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imageData: ""
    };
  },
  methods: {
    onFileChange(e) {
      const files = e.target.files;

      if (files.length > 0) {
        const file = files[0];
        const reader = new FileReader();
        reader.onload = e => {
          this.imageData = e.target.result;
        };
        reader.readAsDataURL(file);
      }
    },
    post() {
      let params = new URLSearchParams();
      params.append("message", "success");
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

