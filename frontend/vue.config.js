const path = require("path");
module.exports = {
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    overlay: false,
    port: 8083
  },
  lintOnSave: false,
  configureWebpack: {
    devtool: 'source-map'
  },
  outputDir: path.resolve(__dirname, "../static"),
  indexPath: path.resolve(__dirname, "../static/index.html"),
}
