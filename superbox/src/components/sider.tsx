import React, {useState} from 'react';
import {
    CarryOutOutlined,
    CloudDownloadOutlined,
    CodepenOutlined,
    FileDoneOutlined,
    PieChartOutlined,
    PlayCircleOutlined, SettingOutlined,
} from '@ant-design/icons';
import type {MenuProps} from 'antd';
import {Layout, Menu} from 'antd';
import {To, useNavigate} from "react-router-dom";

type MenuItem = Required<MenuProps>['items'][number];

function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[],
): MenuItem {
    return {
        key,
        icon,
        children,
        label,
    } as MenuItem;
}

const items: MenuItem[] = [
    getItem('首页', '/index', <PieChartOutlined/>),
    getItem('待办', '/todo', <CarryOutOutlined/>),
    getItem('开发', '/develop', <CodepenOutlined/>),
    getItem('笔记', '/note', <FileDoneOutlined/>),
    getItem('下载', '/download', <CloudDownloadOutlined/>),
    getItem('任务', '/task', <PlayCircleOutlined/>),
    getItem('设置', '/setting', <SettingOutlined/>),
];


export default function Sidebar() {
    const {Sider} = Layout;
    const [collapsed, setCollapsed] = useState(false);
    const navigate = useNavigate()
    const onClick = (menu: { key: To; }) => {
        navigate(menu.key, {replace: true})
    }

    return (
        <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
            <Menu theme="dark" defaultSelectedKeys={["/index"]} mode="inline" items={items} onClick={onClick}/>
        </Sider>
    )
}