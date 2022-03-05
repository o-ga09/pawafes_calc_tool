<template>
  <v-container>
    <v-btn color="red accent-2" elevation="2" outlined @click="reset"
      >リセット</v-btn
    >
    <div class="balloon6">
      <div class="faceicon">
        <img src="../assets/yabe.png" width="150px" height="150px" />
      </div>
      <div class="chatting">
        <div class="says">
          <p>{{ yabes_message }}</p>
        </div>
      </div>
    </div>
  </v-container>
</template>

<script>
export default {
  name: "ResetCalc",

  data: () => ({
    yabes_message: "左側の＋／ーを押してほしいでやんす！！！",
  }),

  methods: {
    reset() {
      this.$parent.server_flg = 0;
      this.$parent.selected_num_of_tornament = "";
      this.$parent.selected_gokigen = "";
      this.$parent.selected_character = "";
      this.$parent.selected_support = "";
      this.$parent.count_hit = 0;
      this.$parent.count_twobase = 0;
      this.$parent.count_threebase = 0;
      this.$parent.count_homerun = 0;
      this.$parent.count_bunt = 0;
      this.$parent.count_fly = 0;
      this.$parent.count_steal = 0;
      this.$parent.count_inings = 0;
      this.$parent.count_sunsinwithfastball = 0;
      this.$parent.count_gorowithfastball = 0;
      this.$parent.count_flywithfastball = 0;
      this.$parent.count_sunsinwithbreakingball = 0;
      this.$parent.count_gorowithbreakingball = 0;
      this.$parent.count_flywithbreakingball = 0;
      this.$parent.count_gettsu = 0;
      this.$parent.goukei = { 筋力: 0, 敏捷: 0, 技術: 0, 変化球: 0, 精神: 0 };
      this.$parent.storage.removeItem("point_data");
    },
  },
};
</script>

<style lang="scss" scoped>
::v-deep .v-select__selections {
  input {
    width: 0;
  }
}

.balloon6 {
  position: relative;
  top: 50px;
  left: 150px;
  width: 100%;
  margin: 10px 0;
  overflow: hidden;
}

.balloon6 .faceicon {
  float: left;
  margin-right: -50px;
  width: 40px;
}

.balloon6 .faceicon img {
  width: 100%;
  height: auto;
  border-radius: 50%;
}
.balloon6 .chatting {
  width: 100%;
  text-align: left;
}
.says {
  display: inline-block;
  position: relative;
  margin: 0 0 0 50px;
  padding: 10px;
  max-width: 250px;
  border-radius: 12px;
  background: #edf1ee;
}

.says:after {
  content: "";
  display: inline-block;
  position: absolute;
  top: 3px;
  left: -19px;
  border: 8px solid transparent;
  border-right: 18px solid #edf1ee;
  -webkit-transform: rotate(35deg);
  transform: rotate(35deg);
}
.says p {
  margin: 0;
  padding: 0;
}
.v-select {
  font-size: 1em;
}
</style>
