<template>
  <v-app>
    <v-navigation-drawer app v-model="drawer" clipped>
      navigation drawer
    </v-navigation-drawer>

    <v-app-bar color="primary" dark app clipped-left>
      <v-app-bar-nav-icon @click="drawer=!drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>hoge</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn outlined>Buuton</v-btn>
    </v-app-bar>

    <v-main>
      <router-view/>
    </v-main>

    <v-footer color="primary" dark app>
      fuga
    </v-footer>
  </v-app>
</template>

<script>

export default {

  data: () => ({
    drawer: null
  }),
};
</script>
