// @ts-ignore
import * as monaco from 'monaco-editor/esm/vs/editor/editor.api.js'

export default function useMonaco(language = 'json', readOnly = true) {
    let monacoEditor: monaco.editor.IStandaloneCodeEditor | null = null
    const updateVal = async (val: string) => {
        monacoEditor?.setValue(val)
        setTimeout(async () => {
            readOnly && monacoEditor?.updateOptions({readOnly: false})
            await monacoEditor?.getAction('editor.action.formatDocument').run()
            readOnly && monacoEditor?.updateOptions({readOnly: true})
        }, 100)
    }

    const createEditor = (el: HTMLElement | null, editorOption: monaco.editor.IStandaloneEditorConstructionOptions = {}) => {
        if (monacoEditor) {
            return
        }
        monacoEditor = el && monaco.editor.create(el, {
            language,
            minimap: {enabled: false},
            theme: 'vs-light',
            multiCursorModifier: 'ctrlCmd',
            scrollbar: {
                verticalScrollbarSize: 8,
                horizontalScrollbarSize: 8,
            },
            tabSize: 4,
            automaticLayout: true, // 自适应宽高
            readOnly: readOnly,
            acceptSuggestionOnCommitCharacter: true,
            accessibilitySupport: 'on',
            autoClosingBrackets: 'always',
            autoClosingDelete: 'always',
            autoClosingQuotes: 'always',
            folding: true,
            ...editorOption
        })
        return monacoEditor
    }
    const onFormatDoc = () => {
        monacoEditor?.getAction('editor.action.formatDocument').run()
    }
    return {
        updateVal,
        getEditor: () => monacoEditor,
        createEditor,
        onFormatDoc
    }
}
