import {createRouter, createWebHistory, RouteRecordRaw} from "vue-router";

// 配置路由
const routes: Array<RouteRecordRaw> = [
    {
        name: "staru",
        path: "/staru",
        component: () => import("./views/staru.vue"),
    },
    {
        name: "system",
        path: "/system",
        component: () => import("./views/system.vue"),
    },
    {
        name: "network",
        path: "/network",
        component: () => import("./views/network.vue"),
    },
    {
        name: "novel",
        path: "/novel",
        component: () => import("./views/novel.vue"),
    },
    {
      name: "struct",
      path: "/struct",
      component: () => import("./views/struct.vue"),
    }
];

// 返回一个 router 实列，为函数，里面有配置项（对象） history
const router = createRouter({
    history: createWebHistory(),
    routes,
});

// 导出路由，在main.ts注册
export default router