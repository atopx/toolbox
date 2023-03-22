const Layout = () => import("@/layout/index.vue");

export default {
    path: "/document",
    name: "document",
    component: Layout,
    redirect: "/document",
    meta: {
        icon: "document",
        title: "知识管理",
        rank: 2
    },
    children: [
        {
            path: "/document/noteList",
            name: "noteList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "笔记管理"
            }
        },
        {
            path: "/document/blogList",
            name: "blogList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "博客管理"
            }
        },
        {
            path: "/document/novelList",
            name: "novelList",
            component: () => import("@/views/welcome/index.vue"),
            meta: {
                title: "图书管理"
            }
        }
    ]
} as RouteConfigsTable;
