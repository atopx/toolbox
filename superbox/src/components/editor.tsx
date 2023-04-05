import Editor, {BeforeMount, OnChange, OnMount} from "@monaco-editor/react";
import React from "react";

interface IProps {
    value?: string
    defaultValue?: string,
    onChange?: OnChange,
    onMount?: OnMount,
    readOnly: boolean,
    scrollbarSize: number,
}

const EditorBox: React.FC<IProps> = (props) => {

    const option = {
        fontSize: 14,
        readOnly: props.readOnly,
        codeLens: false,
        wordWrap: 'wordWrapColumn',
        wordWrapColumn: 57,
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
            verticalScrollbarSize: props.scrollbarSize,
        }
    }
    const beforeMount: BeforeMount = (monaco) => {
        monaco.languages.typescript.javascriptDefaults.setEagerModelSync(true);
    }

    return (
        <Editor
            height="70vh"
            language="json"
            value={props.value}
            defaultValue={props.defaultValue}
            defaultLanguage={"json"}
            options={option}
            beforeMount={beforeMount}
            onMount={props.onMount}
            onChange={props.onChange}
        />
    );
}

export default EditorBox;