import { type RouteRecordRaw, createRouter, createWebHashHistory, createWebHistory } from "vue-router"

const Layout = () => import("@/layout/index.vue")

/** 常驻路由 */
export const constantRoutes: RouteRecordRaw[] = [
    {
        path: "/redirect",
        component: Layout,
        meta: {
            hidden: true
        },
        children: [
            {
                path: "/redirect/:path(.*)",
                component: () => import("@/views/redirect/index.vue")
            }
        ]
    },
    {
        path: "/403",
        component: () => import("@/views/error-page/403.vue"),
        meta: {
            hidden: true
        }
    },
    {
        path: "/404",
        component: () => import("@/views/error-page/404.vue"),
        meta: {
            hidden: true
        },
        alias: "/:pathMatch(.*)*"
    },
    {
        path: "/login",
        component: () => import("@/views/login/index.vue"),
        meta: {
            hidden: true
        }
    },
    {
        path: "/",
        component: Layout,
        redirect: "/dashboard",
        children: [
            {
                path: "dashboard",
                component: () => import("@/views/dashboard/index.vue"),
                name: "Dashboard",
                meta: {
                    title: "首页",
                    svgIcon: "dashboard",
                    affix: true
                }
            }
        ]
    },
    {
        path: "/note",
        component: Layout,
        redirect: "/notes/list",
        name: "Note",
        meta: {
            title: "文档管理",
            elIcon: "Notebook",
            alwaysShow: true // 将始终显示根菜单
        },
        children: [
            {
                path: "list",
                component:  () => import("@/views/note/list.vue"),
                name: "NoteList",
                meta: {
                    title: "笔记列表"
                }
            },
            {
                path: "editor",
                component:  () => import("@/views/note/editor.vue"),
                name: "NoteEditor",
                meta: {
                    title: "编辑笔记",
                    hidden: true,
                }
            },
        ]
    },
    {
        path: "/permission",
        component: Layout,
        redirect: "/permission/roles",
        name: "Permission",
        meta: {
            title: "权限管理",
            svgIcon: "lock",
            alwaysShow: true // 将始终显示根菜单
        },
        children: [
            {
                path: "roles",
                component: () => import("@/views/permission/role.vue"),
                name: "RoleManage",
                meta: {
                    title: "角色管理"
                }
            },
            {
                path: "directive",
                component: () => import("@/views/permission/directive.vue"),
                name: "DirectivePermission",
                meta: {
                    title: "指令权限" // 如果未设置角色，则表示：该页面不需要权限，但会继承根路由的角色
                }
            }
        ]
    },
    {
        path: "/:pathMatch(.*)*", // Must put the 'ErrorPage' route at the end, 必须将 'ErrorPage' 路由放在最后
        redirect: "/404",
        name: "ErrorPage",
        meta: {
            hidden: true
        }
    }
]
export const asyncRoutes: RouteRecordRaw[] = []

const router = createRouter({
    history:
        import.meta.env.VITE_ROUTER_HISTORY === "hash"
            ? createWebHashHistory(import.meta.env.VITE_PUBLIC_PATH)
            : createWebHistory(import.meta.env.VITE_PUBLIC_PATH),
    routes: constantRoutes
})

/** 重置路由 */
export function resetRouter() {
    // 注意：所有动态路由路由必须带有 Name 属性，否则可能会不能完全重置干净
    try {
        router.getRoutes().forEach((route) => {
            const { name, meta } = route
            if (name && meta.roles?.length) {
                router.hasRoute(name) && router.removeRoute(name)
            }
        })
    } catch (error) {
        // 强制刷新浏览器也行，只是交互体验不是很好
        window.location.reload()
    }
}

export default router
