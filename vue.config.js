const path = require("path");

module.exports = {
  outputDir: path.resolve(__dirname, "./view/dist"),
  assetsDir: 'assets',
  chainWebpack: config => {
    config
      .entry("app")
      .clear()
      .add("./view/src/main.js")
      .end();
    config.resolve.alias
      .set("@", path.resolve(__dirname, "./view/src"))
  }
};