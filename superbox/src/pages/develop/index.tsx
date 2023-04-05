import React from 'react';
import {Tabs} from "antd";
import Transform from "./transform";
import TimeTool from "./times";

const tabs = [
    {
        label: "编码",
        key: "transform",
        content: Transform,
    },
    {
        label: "时间",
        key: "/tool/json",
        content: TimeTool,
    }
]


export default function DevelopIndex() {
    return <Tabs
        defaultActiveKey={tabs[0].key}
        centered
        items={tabs.map((tab, i) => {
            console.log(tab.key);
            return {label: tab.label, key: tab.key, children: <tab.content/>};
        })}
    />
};