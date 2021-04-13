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
    config
      .plugin('html')
      .tap(args => {
        args[0].title = 'GoShort | Golang built URL shortener';
        args[0].template = './view/public/index.html';
        return args;
      });
    config.resolve.alias
      .set("@", path.resolve(__dirname, "./view/src"))
  }
};