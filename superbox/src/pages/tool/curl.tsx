import Editor, {OnChange} from "@monaco-editor/react";
import {useState} from "react";

function CurlTool() {
    const options = {

        fontSize: 14,
        readOnly: false,
        codeLens: false,
        fontFamily: "Menlo, Consolas, monospace, sans-serif",
        minimap: {enabled: false},
        quickSuggestions: false,
        lineNumbers: "on",
        renderValidationDecorations: "off"
    };

    const [value, setValue] = useState("{}");

    const onChange = (value: string, event: any)  =>  {
        console.log(value, event)
    }

    return (
        <Editor
            height="80vh"
            language="json"
            value={value}
            options={options}
            // onChange={onChange}
        />
    );
}

export default CurlTool;