import React from "react";
import Editor, {loader, OnChange} from "@monaco-editor/react";

loader.config({
    paths: {
        vs: "/static/monaco/vs",
    }
});

export function processSize(size: string) {
    return !/^\d+$/.test(size) ? size : `${size}px`;
}

interface MonacoProps {
    theme?: string;
    language?: string;
    value?: string;
    width?: number | string;
    height?: number | string;
    options?: any;
    defaultValue?: string;
    onChange: OnChange;
}

export const Monaco: React.FC<MonacoProps> = (
    {
        language,
        value,
        defaultValue,
        height,
        width,
        options,
        onChange
    }) => {
    return (
        <Editor
            defaultLanguage={language}
            defaultValue={defaultValue}
            value={value}
            height={height}
            width={width}
            options={options}
            onChange={onChange}
        />
    );
};

export default Monaco;
