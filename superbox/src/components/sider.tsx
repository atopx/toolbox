import React, {useState} from 'react';
import type {MenuProps} from 'antd';
import {Layout, Menu} from 'antd';
import {To, useNavigate} from "react-router-dom";
import routes, {IRoute} from "../router";


type MenuItem = Required<MenuProps>['items'][number];


function trans(route: IRoute): MenuItem {
    let key = route.path;
    let icon = route.icon;
    let label = route.label;
    let children = route.children?.map((c) => {
        return trans(c)
    })
    return {key, label, icon, children} as MenuItem;
}

const items: MenuItem[] = routes.map((route) => {
    return trans(route)
})


export default function Sidebar() {
    const {Sider} = Layout;
    const [collapsed, setCollapsed] = useState(false);
    const navigate = useNavigate()
    const onClick = (menu: { key: To; }) => {
        navigate(menu.key, {replace: true})
    }

    return (
        <Sider collapsible width={160} collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
            <Menu theme="dark" defaultSelectedKeys={["/"]} mode="inline" items={items} onClick={onClick}/>
        </Sider>
    )
}