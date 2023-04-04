<template>
    <div class="editor-area" :class="isFull ? 'full' : ''" :style="{ width, height }">
        <div class="tools">
            <ElTooltip placement="right" :content="isFull ? '缩小' : '放大'">
                <div class="expand" @click="isFull = !isFull">
                    <i :class="isFull ? 'el-icon-close' : 'el-icon-full-screen'"></i>
                </div>
            </ElTooltip>
            <ElTooltip placement="right" content="格式化">
                <div class="expand" @click="onFormatDoc">
                    <i class="el-icon-finished"></i>
                </div>
            </ElTooltip>
        </div>
    </div>
</template>

<script lang='ts'>
import useMonaco from './monaco'
import {defineComponent, ref} from 'vue'
import {ElTooltip} from 'element-plus'

export default defineComponent({
    components: {
        ElTooltip
    },
    props: {
        width: {
            type: String,
            default: '100%'
        },
        height: {
            type: String,
            default: '90vh'
        },
        language: {
            type: String,
            default: 'json'
        },
        readonly: {
            type: Boolean,
            default: true
        },
        preComment: {
            type: String,
            default: ''
        },
        modelValue: {
            type: String,
            default: ''
        },
        editorOptions: {
            type: Object,
            default: () => ({})
        },
    },
    watch: {
        modelValue(val) {
            val !== this.getEditor()?.getValue() && this.updateMonacoVal(val)
        }
    },
    setup(props) {
        const {updateVal, getEditor, createEditor, onFormatDoc} = useMonaco(props.language, props.readonly)
        const isFull = ref(false)
        return {
            isFull,
            updateVal,
            getEditor,
            createEditor,
            onFormatDoc
        }
    },
    methods: {
        updateMonacoVal(_val?: string) {
            const {modelValue, preComment} = this.$props
            const val = preComment ? `${preComment}\n${_val || modelValue}` : (_val || modelValue)
            this.updateVal(val)
        }
    },
    mounted() {
        if (this.$el) {
            const monacoEditor = this.createEditor(this.$el, this.$props.editorOptions)
            this.updateMonacoVal()
            monacoEditor!.onDidChangeModelContent(() => {
                this.$emit('update:modelValue', monacoEditor!.getValue())
            })
            monacoEditor!.onDidBlurEditorText(() => {
                this.$emit('blur')
            })
        }
    }
})
</script>

<style lang="css" scoped>

.editor-area {
    position: relative;
    border: 1px solid #ddd;
    border-radius: 4px;
    overflow: hidden;
    padding: 5px 5px 5px 0;
    background-color: #fff;
    box-sizing: border-box;
}
</style>
