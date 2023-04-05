import React from 'react';
import {Layout, theme} from 'antd';
import Router from "./components/router";
import Sidebar from "./components/sider";
import Footer from "./components/footer";

const {Content} = Layout;

export default function App() {
    const {token: {colorBgContainer}} = theme.useToken();
    return (
        <Layout style={{minHeight: '100vh'}}>
            <Sidebar/>
            <Layout className="site-layout">
                <Content style={{margin: '0 16px'}}>
                    <div style={{padding: 24, background: colorBgContainer}}>
                        <Router/>
                    </div>
                </Content>
                <Footer/>
            </Layout>
        </Layout>
    );
};
