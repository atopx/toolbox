const Layout = () => import("@/layout/index.vue");

export default {
    path: "/cloud",
    name: "cloud",
    component: Layout,
    redirect: "/cloud",
    meta: {
        icon: "cloud",
        title: "云盘管理",
        rank: 2
    },
    children: [
        {
            path: "/cloud/fileList",
            name: "fileList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "文件管理"
            }
        },
        {
            path: "/cloud/mediaList",
            name: "mediaList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "多媒体"
            }
        },
        {
            path: "/cloud/photoList",
            name: "photoList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "照片墙"
            }
        },
        {
            path: "/cloud/taskList",
            name: "taskList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "离线任务"
            }
        }
    ]
} as RouteConfigsTable;
