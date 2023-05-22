import { type ConfigEnv, type UserConfigExport, loadEnv } from "vite"
import path, { resolve } from "path"
import vue from "@vitejs/plugin-vue"
import vueJsx from "@vitejs/plugin-vue-jsx"
import { createSvgIconsPlugin } from "vite-plugin-svg-icons"
import svgLoader from "vite-svg-loader"
import UnoCSS from "unocss/vite"
import prismjs from "vite-plugin-prismjs"
import DefineOptions from "unplugin-vue-define-options/vite"

/** 配置项文档：https://cn.vitejs.dev/config */
export default (configEnv: ConfigEnv): UserConfigExport => {
    const viteEnv = loadEnv(configEnv.mode, process.cwd()) as ImportMetaEnv
    const { VITE_PUBLIC_PATH } = viteEnv
    return {
        /** 打包时根据实际情况修改 base */
        base: VITE_PUBLIC_PATH,
        resolve: {
            alias: {
                /** @ 符号指向 src 目录 */
                "@": resolve(__dirname, "./src")
            }
        },
        optimizeDeps: {
            include: ["@kangc/v-md-editor/lib/theme/vuepress.js"]
        },
        server: {
            /** 是否开启 HTTPS */
            https: false,
            /** 设置 host: true 才可以使用 Network 的形式，以 IP 访问项目 */
            host: true, // host: "0.0.0.0"
            /** 端口号 */
            port: 18002,
            /** 是否自动打开浏览器 */
            open: false,
            /** 跨域设置允许 */
            cors: true,
            /** 端口被占用时，是否直接退出 */
            strictPort: true,
            /** 接口代理 */
            proxy: {
                "/api/v1": {
                    target: "http://127.0.0.1:18000/api",
                    ws: true,
                    /** 是否允许跨域 */
                    changeOrigin: true
                }
            }
        },
        build: {
            /** 消除打包大小超过 500kb 警告 */
            chunkSizeWarningLimit: 2000,
            /** Vite 2.6.x 以上需要配置 minify: "terser", terserOptions 才能生效 */
            minify: "terser",
            /** 在打包代码时移除 console.log、debugger 和 注释 */
            terserOptions: {
                compress: {
                    drop_console: true,
                    drop_debugger: true,
                    pure_funcs: ["console.log"]
                },
                format: {
                    /** 删除注释 */
                    comments: true
                }
            },
            /** 打包后静态资源目录 */
            assetsDir: "static"
        },
        /** Vite 插件 */
        plugins: [
            vue(),
            vueJsx(),
            /** 将 SVG 静态图转化为 Vue 组件 */
            svgLoader(),
            /** SVG */
            createSvgIconsPlugin({
                iconDirs: [path.resolve(process.cwd(), "src/icons/svg")],
                symbolId: "icon-[dir]-[name]"
            }),
            /** UnoCSS */
            UnoCSS(),
            /** DefineOptions 可以更简单的注册组件名称 */
            DefineOptions(),
            prismjs({
                languages: [
                    "json",
                    "go",
                    "rust",
                    "python",
                    "typescript",
                    "java",
                    "javascript",
                    "sql",
                    "bash",
                    "protobuf"
                ]
            })
        ]
    }
}
