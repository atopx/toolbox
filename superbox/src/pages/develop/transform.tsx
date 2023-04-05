import Editor, {BeforeMount, OnChange, OnMount} from "@monaco-editor/react";
import {useState} from "react";
import {Button, Col, Divider, Row, Select, Space} from "antd";
import {RightCircleOutlined} from "@ant-design/icons";

const selectOptions = [
    {
        value: "0",
        label: "json to yaml",
    },
    {
        value: "1",
        label: "json to go",
    },
    {
        value: "2",
        label: "json to rust",
    },
    {
        value: "3",
        label: "json to protobuf",
    },
    {
        value: "3",
        label: "json to mysql",
    },
    {
        value: "4",
        label: "go to json",
    },
    {
        value: "5",
        label: "rust to json",
    },
    {
        value: "6",
        label: "protobuf to json",
    },
    {
        value: "7",
        label: "yaml to json",
    },
    {
        value: "3",
        label: "mysql to json",
    },
]

function LeftEditor() {
    const options = {
        fontSize: 14,
        readOnly: false,
        codeLens: false,
        wordWrap: 'wordWrapColumn',
        wordWrapColumn: 80,
        accessibilitySupport: 'on',
        fontFamily: "Menlo, Consolas, monospace, sans-serif",
        minimap: {enabled: false},
        quickSuggestions: false,
        lineNumbers: "on",
        renderValidationDecorations: "off",
        renderLineHighlight: "none",
        scrollbar: {
            horizontal: 'hidden',
            vertical: 'hidden',
            useShadows: false,
            verticalHasArrows: false,
            verticalSliderSize: 0,
            verticalScrollbarSize: 1
        }
    };

    const [inputValue, setInputValue] = useState("{}");

    const beforeMount: BeforeMount = (monaco) => {
        monaco.languages.typescript.javascriptDefaults.setEagerModelSync(true);
    }

    const onMount: OnMount = (editor, monaco) => {
        console.log("editor", editor);
    }

    const onChange: OnChange = (change, event) => {
        typeof change === "string" ? setInputValue(change) : undefined;
    }

    return (
        <Editor
            height="70vh"
            language="json"
            value={inputValue}
            defaultLanguage={"json"}
            options={options}
            beforeMount={beforeMount}
            onMount={onMount}
            onChange={onChange}
        />
    );
}

function RightEditor() {
    const options = {
        fontSize: 14,
        readOnly: true,
        codeLens: false,
        wordWrap: 'wordWrapColumn',
        wordWrapColumn: 80,
        accessibilitySupport: 'on',
        fontFamily: "Menlo, Consolas, monospace, sans-serif",
        minimap: {enabled: false},
        quickSuggestions: false,
        lineNumbers: "on",
        renderValidationDecorations: "off",
        renderLineHighlight: "none",
        scrollbar: {
            horizontal: 'hidden',
            vertical: 'hidden',
            useShadows: false,
            verticalHasArrows: false,
            verticalSliderSize: 0,
            verticalScrollbarSize: 0
        }
    };

    const beforeMount: BeforeMount = (monaco) => {
        monaco.languages.typescript.javascriptDefaults.setEagerModelSync(true);
    }

    return (
        <Editor
            height="70vh"
            language="json"
            value={"test"}
            defaultLanguage={"json"}
            options={options}
            beforeMount={beforeMount}
        />
    );
}

function Transform() {
    return (
        <div>
            <Row>
                <Col span={5}>
                    <Select
                        showSearch
                        style={{width: 190}}
                        placeholder="选择动作"
                        optionFilterProp="children"
                        filterOption={(input, option) =>
                            (option?.label ?? '').toLocaleString().toLowerCase().includes(input.toLowerCase())
                        }
                        options={selectOptions}
                    />
                </Col>
                <Col span={3}>
                    <Button type="primary" icon={<RightCircleOutlined/>}>运行</Button>
                </Col>
            </Row>
            <Divider/>
            <Row>
                <Col span={12}><LeftEditor/></Col>
                <Col span={12}><RightEditor/></Col>
            </Row>
        </div>
    );
}

export default Transform;