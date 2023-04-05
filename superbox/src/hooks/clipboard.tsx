import {writeText, readText} from "@tauri-apps/api/clipboard";


async function Get() {
    return await readText() as string;
}

async function Set(object: string) {
    await writeText(object)
}


export default {Get, Set}