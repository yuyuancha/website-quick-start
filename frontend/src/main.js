import { createApp } from "vue";
import App from "./App.vue";
import ViewUIPlus from "view-ui-plus";
import "view-ui-plus/dist/styles/viewuiplus.css";
import { createRouter, createWebHistory } from "vue-router";

const routes = [
    {
        path: "/",
        name: "App",
        component: App,
    },
];

const router = createRouter({
    history: createWebHistory(),
    base: process.env.BASE_URL,
    routes,
});

createApp(App).use(ViewUIPlus).use(router).mount("#app");
