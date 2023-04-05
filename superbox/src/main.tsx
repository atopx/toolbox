import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./App.css";
import {BrowserRouter} from 'react-router-dom';
import {loader} from "@monaco-editor/react";

loader.config({
    paths: {
        vs: "/monaco/vs",
    }
});


ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <React.StrictMode>
        <BrowserRouter>
            <App/>
        </BrowserRouter>
    </React.StrictMode>
);
