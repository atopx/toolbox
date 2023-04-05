import Home from '../views/home';
import DevelopIndex from "../pages/develop/index"
import {Route, Routes} from "react-router-dom";

const routes = [
    {
        path: "/",
        component: Home,
    },
    {
        path: "/index",
        component: Home,
    },
    {
        path: "/develop",
        component: DevelopIndex,
    },
    {
        path: "/todo",
        component: Home,
    },
    {
        path: "/note",
        component: Home,
    },
    {
        path: "/task",
        component: Home,
    },
    {
        path: "/download",
        component: Home,
    },
    {
        path: "/setting",
        component: Home,
    },
];

export default function Router() {
    return (
        <Routes>
            {
                routes.map((r) => {
                    return <Route path={r.path} key={r.path} element={<r.component/>}/>
                })
            }
        </Routes>
    );
};