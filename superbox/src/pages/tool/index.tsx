import React from 'react';
import {Tabs} from "antd";
import Editor from "./json";
import CurlTool from "./curl";

const tabs = [
    {
        label: "CURL",
        key: "/tool/curl",
        content: Editor,
    },
    {
        label: "JSON",
        key: "/tool/json",
        content: CurlTool,
    }
]


export default function ToolIndex() {
    return <Tabs
        defaultActiveKey={tabs[0].key}
        centered
        items={tabs.map((tab, i) => {
            console.log(tab.key);
            return {label: tab.label, key: tab.key, children: <tab.content/>};
        })}
    />
};