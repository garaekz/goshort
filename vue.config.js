/* module.exports = {
  assetsDir: 'assets',
};
 */
const path = require("path");

module.exports = {
  assetsDir: 'assets',
  chainWebpack: config => {
    config
      .entry("app")
      .clear()
      .add("./view/src/main.js")
      .end();
    config.resolve.alias
      .set("@", path.join(__dirname, "./view/src"))
  }
};