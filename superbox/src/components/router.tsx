import Home from '../views/home';
import DevelopIndex from "../pages/develop/index"
import {Route, Routes} from "react-router-dom";

const routes = [
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