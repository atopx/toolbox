export default {
    path: "/computer",
    name: "computer",
    meta: {
        icon: "computer",
        title: "主机管理",
        rank: 3
    },
    children: [
        {
            path: "/computer/list",
            name: "computerList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "设备管理"
            }
        },
        {
            path: "/port/list",
            name: "portList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "端口管理"
            }
        }
    ]
} as RouteConfigsTable;
