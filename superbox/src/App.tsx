import React from 'react';
import {Layout, theme} from 'antd';
import routes, {IRoute} from "./router";
import {Route, Routes} from "react-router-dom";
import Sidebar from "./components/sider";
import Footer from "./components/footer";

const {Content} = Layout;

function router(route: IRoute) {
    return <Route path={route.path} key={route.path} element={route.component}>{route.children?.map((r) => {
        return <Route path={r.path} key={r.path} element={r.component}/>
    })}</Route>
}


export default function App() {
    const {token: {colorBgContainer}} = theme.useToken();
    return (
        <Layout style={{minHeight: '100vh'}}>
            <Sidebar/>
            <Layout className="site-layout">
                <Content style={{margin: '0 16px'}}>
                    <div style={{padding: 20, background: colorBgContainer}}>
                        <Routes>{routes.map((r) => {
                            return router(r)
                        })} </Routes>
                    </div>
                </Content>
                <Footer/>
            </Layout>
        </Layout>
    );
};
