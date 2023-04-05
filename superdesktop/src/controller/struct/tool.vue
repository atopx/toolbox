<script lang="ts" setup>

import MonacoEditor from 'vue-monaco-editor'
import {ref} from "vue";

const code = ref("");

const options = {
    readOnly: true,
}

module.exports = {
    components: {
        Monaco
    },
    data() {
        return {
            code: '// Type away! \n',
            options: {
                selectOnLineNumbers: false
            }
        };
    },
    methods: {
        onMounted(editor) {
            this.editor = editor;
        },
        onCodeChange(editor) {
            console.log(editor.getValue());
        }
    }
};

</script>


<template>

    <MonacoEditor
            height="600"
            language="typescript"
            :code="code"
            :editorOptions="options"
            @mounted="onMounted"
            @codeChange="onCodeChange"
    >
    </MonacoEditor>

</template>