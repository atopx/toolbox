import Home from '../views/home';
import ToolIndex from "../pages/tool/index"
import {Route, Routes} from "react-router-dom";

const routes = [
    {
        path: "/home",
        component: Home,
    },
    {
        path: "/tool",
        component: ToolIndex,
    }
];

export default function Router() {
    return (
        <Routes>
            {
                routes.map((r) => {
                    console.log(r.path)
                    return <Route path={r.path} key={r.path} element={<r.component/>}/>
                })
            }
        </Routes>
    );
};