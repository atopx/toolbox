import Home from './views/home';
import DevelopIndex from "./pages/develop"
import MusicIndex from "./pages/music"


import {
    AuditOutlined,
    CarryOutOutlined,
    CloudDownloadOutlined,
    CodepenOutlined,
    FileDoneOutlined,
    PieChartOutlined,
    PlayCircleOutlined,
    SettingOutlined,
} from '@ant-design/icons';

export interface IRoute {
    path: string,
    label: string,
    icon?: React.ReactNode,
    component?: any,
    children?: Array<IRoute>,
}

const routes: Array<IRoute> = [
    {
        path: "/",
        component: <Home/>,
        label: "首页",
        icon: <PieChartOutlined/>,
    },
    {
        path: "/music",
        label: "音乐",
        icon: <AuditOutlined/>,
        children: [
            {
                path: "/music/index",
                component: <MusicIndex/>,
                label: "首页",
            }
        ]
    },
    {
        path: "/todo",
        component: <Home/>,
        label: "待办",
        icon: <CarryOutOutlined/>,
    },
    {
        path: "/develop",
        component: <DevelopIndex/>,
        label: "开发",
        icon: <CodepenOutlined/>,
    },
    {
        path: "/note",
        component: <Home/>,
        label: "笔记",
        icon: <FileDoneOutlined/>,
    },
    {
        path: "/download",
        component: <Home/>,
        label: "下载",
        icon: <CloudDownloadOutlined/>,
    },
    {
        path: "/task",
        component: <Home/>,
        label: "任务",
        icon: <PlayCircleOutlined/>,
    },
    {
        path: "/setting",
        component: <Home/>,
        label: "设置",
        icon: <SettingOutlined/>,
    },
];

export default routes;
