import * as monaco from 'monaco-editor';


export default function useMonaco(language = 'json', readOnly = true) {
    let editor: monaco.editor.IStandaloneCodeEditor | null;
    const updateVal = async (val: string) => {
        editor?.setValue(val)
        setTimeout(async () => {
            readOnly && editor?.updateOptions({readOnly: false})
            await editor?.getAction('editor.action.formatDocument')?.run()
            readOnly && editor?.updateOptions({readOnly: true})
        }, 100)
    }

    self.MonacoEnvironment = {
        getWorker: function (workerId, label) {
            const getWorkerModule = (moduleUrl: string, label: string) => {
                // @ts-ignore
                return new Worker(self.MonacoEnvironment.getWorkerUrl(moduleUrl), {
                    name: label,
                    type: 'module'
                });
            };

            switch (label) {
                case 'json':
                    return getWorkerModule('/monaco-editor/esm/vs/language/json/json.worker?worker', label);
                default:
                    return getWorkerModule('/monaco-editor/esm/vs/editor/editor.worker?worker', label);
            }
        }
    };

    const createEditor = (el: HTMLElement | null, editorOption: monaco.editor.IStandaloneEditorConstructionOptions = {}) => {
        if (editor) {
            return
        }

        editor = el && monaco.editor.create(el, {
            language,
            minimap: {enabled: false},
            theme: 'vs-light',
            multiCursorModifier: 'ctrlCmd',
            scrollbar: {
                verticalScrollbarSize: 2,
                horizontalScrollbarSize: 2,
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
        return editor
    }
    const onFormatDoc = () => {
        console.log(editor);
        editor?.getAction('editor.action.formatDocument')?.run()
    }
    return {
        updateVal,
        getEditor: () => editor,
        createEditor,
        onFormatDoc
    }
}
