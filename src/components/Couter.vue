<template>
  <v-container>
    <v-tabs v-model="tab" grow class="mb-4">
      <v-tab>野手経験値</v-tab>
      <v-tab>投手経験値</v-tab>
    </v-tabs>
    <v-tabs-items v-model="tab">
      <v-tab-item>
        <v-row>
          <v-col cols="5"><label>ヒット</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_hit"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('hit', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('hit', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>ツーベース</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_twobase"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('twobase', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('twobase', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>スリーベース</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_threebase"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('threebase', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('threebase', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>ホームラン</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_homerun"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('homerun', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('homerun', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>バント</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_bunt"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('bunt', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('bunt', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>犠牲フライ</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_fly"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('fly', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('fly', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>盗塁</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_steal"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('steal', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('steal', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
      </v-tab-item>

      <v-tab-item>
        <v-row>
          <v-col cols="5"><label>登板イニング毎</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_inings"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('inings', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('inings', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>ストレート系で奪三振</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_sunsinwithfastball"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('sunsinwithfastball', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('sunsinwithfastball', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>ストレート系でゴロアウト</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_gorowithfastball"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('gorowithfastball', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('gorowithfastball', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>ストレート系でフライアウト</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_flywithfastball"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('flywithfastball', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('flywithfastball', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>変化球で奪三振</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_sunsinwithbreakingball"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('sunsinwithbreakingball', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('sunsinwithbreakingball', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>変化球でゴロアウト</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_gorowithbreakingball"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('gorowithbreakingball', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('gorowithbreakingball', '+')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>変化球でフライアウト</label></v-col>
          <v-col cols="2"
            ><v-text-field
              v-model="innerCounters.count_flywithbreakingball"
            ></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('flywithbreakingball', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('flywithbreakingball', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
        <v-row>
          <v-col cols="5"><label>併殺</label></v-col>
          <v-col cols="2"
            ><v-text-field v-model="innerCounters.count_gettsu"></v-text-field
          ></v-col>
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('gettsu', '+')"
              ><v-icon dark>mdi-plus</v-icon></v-btn
            ></v-col
          >
          <v-col cols="2"
            ><v-btn
              color="accent"
              elevation="2"
              outlined
              @click="counter('gettsu', '-')"
              ><v-icon dark>mdi-minus</v-icon></v-btn
            ></v-col
          >
        </v-row>
      </v-tab-item>
    </v-tabs-items>
  </v-container>
</template>

<script>
import axios from "axios";

const headers = {
  "Content-Type": "application/json",
};
var storage = sessionStorage;
var getSessionStorageData;
export default {
  name: "CalcCount",
  props: {
    counters: { type: Number, default: 0 },
    goukei: { type: Number, default: 0 },
    serverflg: { type: Number, default: 0 },
    selectedPram: { type: Number, default: 0 },
  },

  data: () => ({
    tab: null,
  }),

  computed: {
    innerCounters: {
      get() {
        return this.$props.counters;
      },

      set(value) {
        this.$emit("update_counter", value);
      },
    },
    innerGoukei: {
      get() {
        return this.$props.goukei;
      },

      set(value) {
        this.$emit("update_goukei", value);
      },
    },
    innerServerflg: {
      get() {
        console.log("get"+this.$props.serverflg)
        return this.$props.serverflg;
      },

      set(value) {
        console.log("set"+value)
        this.$emit("update_server_flg", value);
      },
    },
    innerSelectedPram: {
      get() {
        return this.$props.selectedPram;
      },
      set(value) {
        this.$emit("update_pram", value);
      }
    }
  },

  methods: {
    async counter(action, key) {
      if (
        this.innerSelectedPram.selected_character == "" ||
        this.innerSelectedPram.selected_support == "" ||
        this.innerSelectedPram.selected_gokigen == ""
      ) {
        alert("キャラクター/サポートアイテムを選択してください...");
        return;
      }

      console.log(this.innerServerflg + ":" + this.getSessionStorageData);
      if (this.innerServerflg == 0) {
        await axios
          .post(
            "/execute",
            {
              helper: this.selected_character,
              supportitem: this.selected_support,
              gokigen: this.selected_gokigen,
            },
            { headers: headers }
          )
          .then(function (res) {
            storage.setItem("point_data", JSON.stringify(res.data));
            getSessionStorageData = JSON.parse(storage.getItem("point_data"));
          });
        this.innerServerflg = 1;
      }
      this.innerGoukei = this.goukei;

      if (key == "+") {
        switch (action) {
          case "hit":
            this.innerCounters.count_hit++;
            this.innerGoukei["筋力"] += getSessionStorageData.hit.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.hit.agile;
            this.innerGoukei["技術"] += getSessionStorageData.hit.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.hit.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.hit.mental;
            break;
          case "twobase":
            this.innerCounters.count_twobase++;
            this.innerGoukei["筋力"] += getSessionStorageData.twobase.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.twobase.agile;
            this.innerGoukei["技術"] += getSessionStorageData.twobase.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.twobase.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.twobase.mental;
            break;
          case "threebase":
            this.innerCounters.count_threebase++;
            this.innerGoukei["筋力"] += getSessionStorageData.threebase.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.threebase.agile;
            this.innerGoukei["技術"] += getSessionStorageData.threebase.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.threebase.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.threebase.mental;
            break;
          case "homerun":
            this.innerCounters.count_homerun++;
            this.innerGoukei["筋力"] += getSessionStorageData.homerun.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.homerun.agile;
            this.innerGoukei["技術"] += getSessionStorageData.homerun.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.homerun.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.homerun.mental;
            break;
          case "bunt":
            this.innerCounters.count_bunt++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.SacrificeBunt.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.SacrificeBunt.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.SacrificeBunt.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.SacrificeBunt.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.SacrificeBunt.mental;
            break;
          case "fly":
            this.innerCounters.count_fly++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Sacrificefly.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Sacrificefly.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.Sacrificefly.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Sacrificefly.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Sacrificefly.mental;
            break;
          case "steal":
            this.innerCounters.count_steal++;
            this.innerGoukei["筋力"] += getSessionStorageData.Steal.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.Steal.agile;
            this.innerGoukei["技術"] += getSessionStorageData.Steal.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Steal.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.Steal.mental;
            break;
          case "inings":
            this.innerCounters.count_inings++;
            this.innerGoukei["筋力"] += getSessionStorageData.Innings.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.Innings.agile;
            this.innerGoukei["技術"] += getSessionStorageData.Innings.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Innings.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.Innings.mental;
            break;
          case "sunsinwithfastball":
            this.count_sunsinwithfastball++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Outwithfastball.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Outwithfastball.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.Outwithfastball.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Outwithfastball.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Outwithfastball.mental;
            break;
          case "gorowithfastball":
            this.innerCounters.count_gorowithfastball++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Grounderwithfastball.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Grounderwithfastball.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.Grounderwithfastball.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Grounderwithfastball.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Grounderwithfastball.mental;
            break;
          case "flywithfastball":
            this.innerCounters.count_flywithfastball++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Popflywithfastball.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Popflywithfastball.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.hiPopflywithfastballt.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.hPopflywithfastballit.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Popflywithfastball.mental;
            break;
          case "sunsinwithbreakingball":
            this.innerCounters.count_sunsinwithbreakingball++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Outwithbreakingball.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Outwithbreakingball.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.Outwithbreakingball.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Outwithbreakingball.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Outwithbreakingball.mental;
            break;
          case "gorowithbreakingball":
            this.innerCounters.count_gorowithbreakingball++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Grounderwithbreakingball.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Grounderwithbreakingball.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.Grounderwithbreakingball.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Grounderwithbreakingball.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Grounderwithbreakingball.mental;
            break;
          case "flywithbreakingball":
            this.innerCounters.count_flywithbreakingball++;
            this.innerGoukei["筋力"] +=
              getSessionStorageData.Popflywithbreaking.power;
            this.innerGoukei["敏捷"] +=
              getSessionStorageData.Popflywithbreaking.agile;
            this.innerGoukei["技術"] +=
              getSessionStorageData.Popflywithbreaking.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Popflywithbreaking.breaking_point;
            this.innerGoukei["精神"] +=
              getSessionStorageData.Popflywithbreaking.mental;
            break;
          case "gettsu":
            this.innerCounters.count_gettsu++;
            this.innerGoukei["筋力"] += getSessionStorageData.Doubleplay.power;
            this.innerGoukei["敏捷"] += getSessionStorageData.Doubleplay.agile;
            this.innerGoukei["技術"] += getSessionStorageData.Doubleplay.techni;
            this.innerGoukei["変化球"] +=
              getSessionStorageData.Doubleplay.breaking_point;
            this.innerGoukei["精神"] += getSessionStorageData.Doubleplay.mental;
            break;
        }
      } else if (key == "-") {
        switch (action) {
          case "hit":
            this.innerCounters.count_hit--;
            this.innerGoukei["筋力"] -= getSessionStorageData.hit.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.hit.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.hit.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.hit.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.hit.mental;
            break;
          case "twobase":
            this.count_twobase--;
            this.innerGoukei["筋力"] -= getSessionStorageData.twobase.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.twobase.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.twobase.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.twobase.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.twobase.mental;
            break;
          case "threebase":
            this.innerCounters.count_threebase--;
            this.innerGoukei["筋力"] -= getSessionStorageData.threebase.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.threebase.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.threebase.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.threebase.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.threebase.mental;
            break;
          case "homerun":
            this.innerCounters.count_homerun--;
            this.innerGoukei["筋力"] -= getSessionStorageData.homerun.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.homerun.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.homerun.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.homerun.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.homerun.mental;
            break;
          case "bunt":
            this.innerCounters.count_bunt--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.SacrificeBunt.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.SacrificeBunt.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.SacrificeBunt.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.SacrificeBunt.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.SacrificeBunt.mental;
            break;
          case "fly":
            this.innerCounters.count_fly--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Sacrificefly.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Sacrificefly.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.Sacrificefly.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Sacrificefly.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Sacrificefly.mental;
            break;
          case "steal":
            this.innerCounters.count_steal--;
            this.innerGoukei["筋力"] -= getSessionStorageData.Steal.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.Steal.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.Steal.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Steal.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.Steal.mental;
            break;
          case "inings":
            this.innerCounters.count_inings--;
            this.innerGoukei["筋力"] -= getSessionStorageData.Innings.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.Innings.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.Innings.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Innings.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.Innings.mental;
            break;
          case "sunsinwithfastball":
            this.innerCounters.count_sunsinwithfastball--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Outwithfastball.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Outwithfastball.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.Outwithfastball.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Outwithfastball.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Outwithfastball.mental;
            break;
          case "gorowithfastball":
            this.innerCounters.count_gorowithfastball--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Grounderwithfastball.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Grounderwithfastball.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.Grounderwithfastball.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Grounderwithfastball.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Grounderwithfastball.mental;
            break;
          case "flywithfastball":
            this.innerCounters.count_flywithfastball--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Popflywithfastball.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Popflywithfastball.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.hiPopflywithfastballt.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.hPopflywithfastballit.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Popflywithfastball.mental;
            break;
          case "sunsinwithbreakingball":
            this.innerCounters.count_sunsinwithbreakingball--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Outwithbreakingball.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Outwithbreakingball.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.Outwithbreakingball.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Outwithbreakingball.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Outwithbreakingball.mental;
            break;
          case "gorowithbreakingball":
            this.innerCounters.count_gorowithbreakingball--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Grounderwithbreakingball.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Grounderwithbreakingball.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.Grounderwithbreakingball.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Grounderwithbreakingball.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Grounderwithbreakingball.mental;
            break;
          case "flywithbreakingball":
            this.innerCounters.count_flywithbreakingball--;
            this.innerGoukei["筋力"] -=
              getSessionStorageData.Popflywithbreaking.power;
            this.innerGoukei["敏捷"] -=
              getSessionStorageData.Popflywithbreaking.agile;
            this.innerGoukei["技術"] -=
              getSessionStorageData.Popflywithbreaking.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Popflywithbreaking.breaking_point;
            this.innerGoukei["精神"] -=
              getSessionStorageData.Popflywithbreaking.mental;
            break;
          case "gettsu":
            this.innerCounters.count_gettsu--;
            this.innerGoukei["筋力"] -= getSessionStorageData.Doubleplay.power;
            this.innerGoukei["敏捷"] -= getSessionStorageData.Doubleplay.agile;
            this.innerGoukei["技術"] -= getSessionStorageData.Doubleplay.techni;
            this.innerGoukei["変化球"] -=
              getSessionStorageData.Doubleplay.breaking_point;
            this.innerGoukei["精神"] -= getSessionStorageData.Doubleplay.mental;
            break;
        }
        this.isnegative();
      }
    },

    isnegative() {
      if (this.goukei["筋力"] <= 0) {
        this.innerGoukei["筋力"] = 0;
      }
      if (this.goukei["敏捷"] <= 0) {
        this.innerGoukei["敏捷"] = 0;
      }
      if (this.goukei["技術"] <= 0) {
        this.innerGoukei["技術"] = 0;
      }
      if (this.goukei["変化球"] <= 0) {
        this.innerGoukei["変化球"] = 0;
      }
      if (this.goukei["精神"] <= 0) {
        this.innerGoukei["精神"] = 0;
      }
    },
  },
};
</script>
