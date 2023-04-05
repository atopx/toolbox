import {OnChange, OnMount} from "@monaco-editor/react";
import React, {useEffect, useState} from "react";
import {Button, Col, Divider, Row, Select} from "antd";
import {RightCircleOutlined} from "@ant-design/icons";
import {invoke} from '@tauri-apps/api/tauri'
import EditorBox from "../../components/editor";
import Clipboard from "../../hooks/clipboard";

function Transform() {
    const [actions, setActions] = useState<string[]>([]);
    const [action, setAction] = useState("");
    const [inputValue, setInputValue] = useState("");
    const [outputValue, setOutputValue] = useState("");

    useEffect(() => {
        // 初始化data
        invoke("list_develop_action").then((data: any) => {
            setActions(data)
        })
    }, [])

    const onSelectAction = (value: string) => {
        setAction(value)
    }

    const leftOnChange: OnChange = (change, event) => {
        if (typeof change === "string") {
            setInputValue(change)
        }
    }

    const leftOnMount: OnMount = (editor, monaco) => {
        setInputValue(editor.getValue())
    }

    const trans = async () => {
        // call server
        await invoke("exec_develop_action", {action: action, src: inputValue}).then((resp) => {
            let data = resp as string;
            // show
            setOutputValue(data);
            // copy
            Clipboard.Set(data)

        }).catch((err) => {
            setOutputValue(err)
        })
    }

    const trans_string = (v: string | number | null | undefined): string => {
        if (v === null || v === undefined) {
            return ""
        }
        return v.toString()
    }


    return (
        <div>
            <Row>
                <Col span={5}>
                    <Select
                        showSearch
                        style={{width: 190}}
                        placeholder="选择动作"
                        optionFilterProp="children"
                        onSelect={onSelectAction}
                        filterOption={(input, option) =>
                            trans_string(option?.value).toLowerCase().includes(input.toLowerCase())
                        }>
                        {
                            // 初始化select
                            actions.map((action) => {
                                return <Select.Option key={action} value={action} children={action}/>
                            })
                        }
                    </Select>
                </Col>
                <Col span={3}>
                    <Button type="primary" icon={<RightCircleOutlined/>} onClick={trans}>运行</Button>
                </Col>
            </Row>
            <Divider/>
            <Row>
                <Col span={12}><EditorBox readOnly={false} onChange={leftOnChange} onMount={leftOnMount}
                                          defaultValue={"{}"} scrollbarSize={1}/></Col>
                <Col span={12}><EditorBox readOnly={true} scrollbarSize={0} value={outputValue}/></Col>
            </Row>
        </div>
    );
}

export default Transform;