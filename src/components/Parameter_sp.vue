<template>
  <v-container>
    <v-col cols="12" sm="6" class="my-n4">
      <v-select
        v-model="innerSelectedPram.selected_character"
        :items="items_character"
        label="サポートキャラを選択"
        outlined
      ></v-select>
    </v-col>
    <v-col cols="12" sm="6" class="my-n4">
      <v-select
        v-model="innerSelectedPram.selected_support"
        :items="items_support"
        label="難易度を選択してください"
        outlined
      ></v-select>
    </v-col>
    <v-col cols="12" sm="6" class="my-n4">
      <v-select
        v-model="innerSelectedPram.selected_num_of_tornament"
        :items="num_of_tornament"
        label="何回戦かを選択してください"
        outlined
      ></v-select>
    </v-col>

    <v-col cols="12" sm="6" class="my-n4">
      <v-select
        v-model="innerSelectedPram.selected_gokigen"
        :items="items_gokigen"
        label="サポートキャラのご機嫌度"
        outlined
      ></v-select>
    </v-col>
    <v-col cols="12" sm="6" class="my-n4">
      <v-select
        v-model="innerSelectedPram.selected_lost_point"
        :items="lost_point"
        label="失点数を入力してください"
        outlined
      ></v-select>
    </v-col>
    <v-col cols="12" sm="6" class="my-n4">
      <v-select
        v-model="innerSelectedPram.selected_goal_difference"
        :items="goal_difference"
        label="得失点差を入力してください"
        outlined
      ></v-select>
    </v-col>
  </v-container>
</template>

<script>
export default {
  name: "SelectParameterSp",
  props: {
    selectedpram: {type:Number,default:0}
  },
  data: () => ({
    items_character: [
      "明星雪華",
      "木場静香",
      "七瀬はるか",
      "緒川美羽",
      "倉家凪",
      "須神絵久",
      "エミリ",
      "神良美砂",
      "嵐山美鈴",
      "鴨川しぐれ",
      "虹谷彩理",
      "我間摩夕",
      "姫野カレン",
      "紺野美崎",
      "黒沢愛",
      "四条澄香",
      "栗原舞",
      "咲須かのん",
      "NCM-753",
      "泡瀬満里南",
      "片桐恋",
      "片桐恋（ヤンデレ）",
      "久根美亜",
      "三ツ沢環",
      "氷上聡里",
      "日和ミヨ",
    ],
    items_support: ["ルーキーのお守り", "達人のお守り", "アイテムなし"],
    num_of_tornament: ["1", "2", "3", "4", "final", "boss"],
    lost_point: ["-1", "-2", "-3", "-4"],
    goal_difference: [
      "-5",
      "-4",
      "-3",
      "-2",
      "-1",
      "1",
      "2",
      "3",
      "4",
      "5",
      "6",
      "7",
      "8",
      "9",
      "10",
    ],
    items_gokigen: ["1", "2", "3"],
  }),
  computed: {
    innerSelectedPram: {
      get() {
        return this.$props.selectedpram
      },
      set(value) {
        this.$emit('update_param',value)
      }
    }
  }
};
</script>

<style lang="scss" scoped>
::v-deep .v-select__selections {
  input {
    width: 0;
  }
}

.v-select {
  font-size: 1em;
}
</style>
