<template>
  <v-footer>
    <v-container>
      <table border="1">
        <tr>
          <th v-for="thead in theads" :key="thead">{{ thead }}</th>
        </tr>
        <tr>
          <td v-for="tdata in innerGoukei" :key=tdata>{{ tdata }}</td>
        </tr>
      </table>
    </v-container>
  </v-footer>
</template>

<script>
export default {
  name: "FlexFooter",
  props: {
    goukei: { type: Number, default: 0 },
  },
  computde: {
    innerGoukei: {
      get() {
        return this.$props.goukei;
      },

      set(value) {
        this.$emit("change", value);
      },
    },
  },
  data: () => ({
    theads: ["筋力", "敏捷", "技術", "変化球", "精神"],
  }),
};
</script>
