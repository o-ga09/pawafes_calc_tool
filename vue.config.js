const webpack = require('webpack'); 
module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  configureWebpack: { 
    plugins: [
      new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/),   
      '~/plugins/vue-mq'
    ]
  }
}
