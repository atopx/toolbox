import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
 
// 配置路由
const routes: Array<RouteRecordRaw> = [
  {
    name: "greet",
    path: "/greet",
    component: () => import("./views/greet.vue"),
  },

  {
    name: "novel",
    path: "/novel",
    component: () => import("./views/novel.vue"),
  },

];

// 返回一个 router 实列，为函数，里面有配置项（对象） history
const router = createRouter({
  history: createWebHistory(),
  routes,
});
 
// 导出路由，在main.ts注册
export default router