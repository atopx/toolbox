import React, {useState} from 'react';
import {DesktopOutlined, CodepenOutlined, PieChartOutlined, TeamOutlined, UserOutlined,} from '@ant-design/icons';
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
    getItem('Option 2', '/home', <DesktopOutlined/>),
    getItem('User', 'sub1', <UserOutlined/>),
    getItem('Team', 'sub2', <TeamOutlined/>),
    getItem('工具', '/tool', <CodepenOutlined />),
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
            <Menu theme="dark" defaultSelectedKeys={["/", "/index"]} mode="inline" items={items} onClick={onClick}/>
        </Sider>
    )
}