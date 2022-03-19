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
var storage = sessionStorage

export default {
  name: "ResetCalc",
  props: {
    selectedPram: {type:Number,default:0},
    counters: {type:Number,default:0},
    goukei: {type:Number,dafault:0},
    serverflg: {type:Number,default:0}
  },

  data: () => ({
    yabes_message: "左側の＋／ーを押してほしいでやんす！！！",
  }),

  computed: {
    innerSelectedPram: {
      get() {
        return this.$props.selectedPram
      },
      set(value) {
        this.$emit("update_param",value)
      }
    },
    innerCounters: {
      get() {
        return this.$props.counters
      },
      set(value) {
        this.$emit("update_counter",value)
      }
    },
    innerGoukei: {
      get() {
        return this.$props.goukei
      },
      set(value) {
        this.$emit("update_goukei",value)
      }
    },
    innerServerflg: {
      get() {
        return this.$props.serverflg
      },
      set(value) {
        this.$emit("update_server_flg",value)
      }
    }
  },

  methods: {
    reset() {
      this.innerServerflg = 0;
      this.innerSelectedPram.selected_num_of_tornament = "";
      this.innerSelectedPram.selected_gokigen = "";
      this.innerSelectedPram.selected_character = "";
      this.innerSelectedPram.selected_support = "";
      this.innerCounters.count_hit = 0;
      this.innerCounters.count_twobase = 0;
      this.innerCounters.count_threebase = 0;
      this.innerCounters.count_homerun = 0;
      this.innerCounters.count_bunt = 0;
      this.innerCounters.count_fly = 0;
      this.innerCounters.count_steal = 0;
      this.innerCounters.count_inings = 0;
      this.innerCounters.count_sunsinwithfastball = 0;
      this.innerCounters.count_gorowithfastball = 0;
      this.innerCounters.count_flywithfastball = 0;
      this.innerCounters.count_sunsinwithbreakingball = 0;
      this.innerCounters.count_gorowithbreakingball = 0;
      this.innerCounters.count_flywithbreakingball = 0;
      this.innerCounters.count_gettsu = 0;
      this.innerGoukei.goukei = { 筋力: 0, 敏捷: 0, 技術: 0, 変化球: 0, 精神: 0 };
      storage.removeItem("point_data");
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
