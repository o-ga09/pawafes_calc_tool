<template>
  <div>
    <div v-if="$mq === 'pc'">
      <SelectParameterPc :selectedpram="selected_pram" @update_pram="receivePram"/>
      <CalcCount :counters="counters" :goukei="goukei" :serverflg="server_flg" @update_goukei="receiveGoukei" @update_counter="receiveCounter" @update_server_flg="receiveServerflg" />
      <DisplayTotal :goukei="goukei" />
      <ResetCalc :selectedpram="selected_parm" :counters="counters" :goukei="goukei" :serverflg="server_flg" @update_pram="receivePram" @update_goukei="receiveGoukei" @update_counter="receiveCounter" @update_server_flg="receiveServerflg" />
    </div>

    <div v-if="$mq === 'sp'">
      <SelectParameterSp :selectedpram="selectedPram" @update_param="receivePram"/>
      <CalcCount :counters="counters" :goukei="goukei" :serverflg="server_flg" @update_goukei="receiveGoukei" @update_counter="receiveCounter" @update_server_flg="receiveServerflg" />
      <ResetCalc :selectedpram="selected_parm" :counters="counters" :goukei="goukei" :serverflg="server_flg" @update_pram="receivePram" @update_goukei="receiveGoukei" @update_counter="receiveCounter" @update_server_flg="receiveServerflg" />
      <FlexFooter :goukei="goukei" />
    </div>
  </div>
</template>

<script>
import SelectParameterPc from "../components/Parameter_pc";
import SelectParameterSp from "../components/Parameter_sp";
import CalcCount from "../components/Couter";
import DisplayTotal from "../components/Displaytotal";
import ResetCalc from "../components/Reset";
import FlexFooter from "../components/Flexfooter";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Index",
  components: {
    SelectParameterPc,
    SelectParameterSp,
    CalcCount,
    DisplayTotal,
    ResetCalc,
    FlexFooter,
  },
  data: () => ({
    server_flg: 0,
    selected_pram: {
      selected_character: "",
      selected_support: "",
      selected_num_of_tornament: "",
      selected_gokigen: "",
      selected_lost_point:"",
      selected_goal_difference:0
    },
    counters: {
      count_hit: 0,
      count_twobase: 0,
      count_threebase: 0,
      count_homerun: 0,
      count_bunt: 0,
      count_fly: 0,
      count_steal: 0,
      count_inings: 0,
      count_sunsinwithfastball: 0,
      count_gorowithfastball: 0,
      count_flywithfastball: 0,
      count_sunsinwithbreakingball: 0,
      count_gorowithbreakingball: 0,
      count_flywithbreakingball: 0,
      count_gettsu: 0,
    },
    goukei: { 筋力: 0, 敏捷: 0, 技術: 0, 変化球: 0, 精神: 0 },
  }),
  methods: {
    receiveGoukei(total){
      this.goukei = total
    },
    receiveCounter(counter) {
      this.counters = counter
    },
    receiveServerflg(server_flg) {
      this.server_flg = server_flg
      console.log("update"+this.server_flg)
    },
    receivePram(param) {
      this.selected_pram = param
    }
  }
};
</script>
