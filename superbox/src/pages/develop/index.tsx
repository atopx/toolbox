import React from 'react';
import {Tabs} from "antd";
import Transform from "./transform";
import Network from "./network";

const tabs = [
    {
        label: "编码",
        key: "transform",
        Component: Transform,
    },
    {
        label: "网络",
        key: "network",
        Component: Network,
    },
]


export default function DevelopIndex() {
    return <Tabs
        defaultActiveKey={tabs[0].key}
        centered
        items={tabs.map((tab, i) => {
            return {label: tab.label, key: tab.key, children: <tab.Component/>};
        })}
    />
};