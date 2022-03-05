<template>
  <v-container>
    <v-row justify="center">
      <v-col cols="4"><label>筋力</label></v-col>
      <v-col cols="8"
        ><v-text-field v-model="innerGoukei['筋力']" readonly></v-text-field
      ></v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="4"><label>敏捷</label></v-col>
      <v-col cols="8"
        ><v-text-field v-model="innerGoukei['敏捷']" readonly></v-text-field
      ></v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="4"><label>技術</label></v-col>
      <v-col cols="8"
        ><v-text-field v-model="innerGoukei['技術']" readonly></v-text-field
      ></v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="4"><label>変化球</label></v-col>
      <v-col cols="8"
        ><v-text-field v-model="innerGoukei['変化球']" readonly></v-text-field
      ></v-col>
    </v-row>
    <v-row justify="center">
      <v-col cols="4"><label>精神</label></v-col>
      <v-col cols="8"
        ><v-text-field v-model="innerGoukei['精神']" readonly></v-text-field
      ></v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
    name: 'DisplayTotal',
    props: {
      goukei: {type:Number,default:0}
    },
    computed: {
      innerGoukei: {
        get() {
          return this.$props.goukei
        },
        set(value) {
        this.$emit('change',value )
        }
      }
    } 
}
</script>