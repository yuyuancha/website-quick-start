const path = require("path");

module.exports = {
    // 只作用于当前目录
    root: true,
    // 运行环境
    env: {
        node: true,
        es6: true,
    },
    // 解析器
    parser: "@babel/eslint-parser",
    // 解析器选项
    parserOptions: {
        sourceType: "module",
    },
    // 插件
    plugins: ["import"],
    // 扩展配置
    extends: [
        "plugin:vue/vue3-recommended",
        "plugin:import/recommended",
        "prettier",
    ],
    // 启用规则
    rules: {
        "vue/multi-word-component-names": 0,
    },
    // 全局变量
    globals: {
        h: true,
    },
    // 为指定文件指定处理器
    overrides: [
        {
            files: ["*.vue", "*.jsx"],
            parser: "vue-eslint-parser",
            parserOptions: {
                ecmaVersion: 2018,
            },
            rules: {},
        },
    ],
    settings: {
        "import/resolver": {
            alias: {
                map: [
                    ["@", "./src"],
                    ["@assets", "./src/assets"],
                    ["@components", "./src/components"],
                    ["@views", "./src/views"],
                    ["@views", "./src/views"],
                ],
            },
        },
    },
};
