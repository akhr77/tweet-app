<template>
  <div id="post">
    <!-- <div class="post-form-area">
      <i class="fas fa-camera fa-5x post-form"></i>
      <input type="file" accept="image/*" @change="onFileChange($event)" />
      <img class="preview" :src="imageData" />
    </div>-->

    <div class="container">
      <div class="resize-img">
        <!-- 画像選択 -->
        <div v-show="!resizedImg" class="resize-img__post">
          <label for="file" class="resize-img__post__label">
            <i class="fas fa-camera fa-5x"></i>
            <input id="file" ref="fileInput" type="file" accept=".jpeg, .png" @change="uploadImg" />
          </label>
        </div>
        <!-- プレビュー -->
        <div v-show="resizedImg" class="resize-img__preview">
          <canvas ref="canvas" class="resize-img__preview__canvas" />
          <div class="buttuns">
            <b-button rounded @click="clearUploadImg">キャンセル</b-button>
            <b-button rounded @click="post">投稿</b-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      fileName: "",
      fileType: "",
      imageData: "",
      resizedImg: null
    };
  },
  destroyed() {
    this.clearUploadImg();
  },
  methods: {
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
    },

    uploadImg(e) {
      const file = e.target.files[0];
      const reader = new FileReader();
      this.fileName = file.name;
      this.fileType = file.type;
      reader.onload = e => {
        this.generateImgUrl(e.target.result);
      };
      reader.readAsDataURL(file);
    },

    generateImgUrl(file) {
      const image = new Image();
      image.crossOrigin = "Anonymous";

      image.onload = () => {
        const resizedBase64 = this.makeResizeImg(image);
        this.imageData = resizedBase64;
        // リサイズ済みのBase64をblobに変換
        const resizedBlob = this.base64ToBlob(resizedBase64);
        // urlを生成してプレビュー表示できるようにする
        const resizedImg = window.URL.createObjectURL(resizedBlob);
        this.resizedImg = resizedImg;
      };
      image.src = file;
    },

    makeResizeImg(image) {
      const canvas = this.$refs.canvas;
      const ctx = canvas.getContext("2d");
      // 縦横で長い方の最大値を1000
      const MAX_SIZE = 580;

      if (image.width < MAX_SIZE && image.height < MAX_SIZE) {
        [canvas.width, canvas.height] = [image.width, image.height];
        ctx.drawImage(image, 0, 0);
        return canvas.toDataURL("image/jpeg");
      }

      let dstWidth;
      let dstHeight;
      // 縦横比の計算
      if (image.width > image.height) {
        dstWidth = MAX_SIZE;
        dstHeight = (image.height * MAX_SIZE) / image.width;
      } else {
        dstHeight = MAX_SIZE;
        dstWidth = (image.width * MAX_SIZE) / image.height;
      }
      canvas.width = dstWidth;
      canvas.height = dstHeight;
      // リサイズ
      ctx.drawImage(
        image,
        0,
        0,
        image.width,
        image.height,
        0,
        0,
        dstWidth,
        dstHeight
      );

      return canvas.toDataURL("image/jpeg");
    },

    clearUploadImg() {
      this.resizedImg = null;
      if (this.$refs.fileInput && this.$refs.fileInput.value !== undefined) {
        this.$refs.fileInput.value = "";
      }
    },

    base64ToBlob(base64) {
      const bin = atob(base64.replace(/^.*,/, ""));
      const buffer = new Uint8Array(bin.length);
      for (let i = 0; i < bin.length; i++) {
        buffer[i] = bin.charCodeAt(i);
      }
      return new Blob([buffer.buffer], {
        type: "image/png"
      });
    }
  }
};
</script>

<style scoped lang="scss">
.post-form-area {
  max-width: 550px;
  margin: 0 auto 1rem auto;
  width: 100%;
}

.resize-img {
  width: 300px;
  height: 300px;
  margin: 0 auto;
  margin-top: 20px;

  &__post {
    border: 1px solid rgba(#000, 0.16);
    line-height: 30rem;

    &__label {
      display: inline-block;
      width: 100%;
      color: rgba(0, 0, 0, 0.4);
      text-align: center;

      & > input {
        display: none;
      }
    }
  }

  &__preview {
    width: 300px;
    height: 300px;

    &__canvas {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }
}
</style>

