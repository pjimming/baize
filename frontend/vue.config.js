const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    allowedHosts: "all",
    port: 8080,
    proxy: {
      "/api": {
        target: "http://0.0.0.0:8888",
        changOrigin: true,
        pathRewrite: { "^/api": "" },
      },
    },
  },
});
