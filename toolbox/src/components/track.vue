<script setup>
import { ref } from "vue";


const trace_id = ref("");
const range_time = ref([]);

async function search() {
    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Cookie: 'sid=Fe26.2**90468cf1810aefa339e94133fb23b0b2da7b0167c06aa96d792341e73c5d314c*7I4Ye7U4bRf8QYQjViTwCA*wgZIQPqDYswcNGy8U6V2tAuFNlVZMj8dSq2bL0SFgZ6i66_T7mZJbJKxTNX0KhJrTesEeGas4SGTmjcqEleHYEIzvvceMcDwcppFU_M3avnq8-3uoebphH9NLgnXa3cSs0H-pw8y0rhUfjQ6YfbtmyF-um-Y3qZB0tKbBt0ChiaRrG44gtMr7ragWmQc6Ghp0Cj1Qvd0bBx2oJoSqYBTaiW1p02ormLTY139cCMH-3g**34d339e0d8465e6fb3841b2fa13701f9a6696e70433ad2ec3184e35697e58096*hfnK-JejkZtGi9AYmWf3zG0fDprmJvEZj-YJZbtUFdU',
            Origin: 'http://dm-kibana-tx.domob.cn',
            'Proxy-Connection': 'keep-alive',
            Referer: 'http://dm-kibana-tx.domob.cn/s/staru/app/discover',
            'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36',
            'kbn-version': '7.13.2'
        },
        body: '{"batch":[{"request":{"id":"FmRaTzJtYkU3UXR1NUg3N2NkR05PeGcgSk5hRDl3TGhSZ1NlU243aU8zblU5ZzoyNjg2MTM1NzM=","params":{"index":"star*","body":{"size":500,"sort":[{"time":{"order":"desc","unmapped_type":"boolean"}}],"version":true,"fields":[{"field":"*","include_unmapped":"true"},{"field":"@timestamp","format":"strict_date_optional_time"},{"field":"time","format":"strict_date_optional_time"}],"aggs":{"2":{"date_histogram":{"field":"time","fixed_interval":"30s","time_zone":"Asia/Shanghai","min_doc_count":1}}},"script_fields":{},"stored_fields":["*"],"runtime_mappings":{},"_source":false,"query":{"bool":{"must":[],"filter":[{"bool":{"should":[{"match":{"context.trackid":"1676454965686669162"}}],"minimum_should_match":1}},{"range":{"time":{"gte":"2023-02-15T09:41:53.719Z","lte":"2023-02-15T09:56:53.719Z","format":"strict_date_optional_time"}}}],"should":[],"must_not":[]}},"highlight":{"pre_tags":["@kibana-highlighted-field@"],"post_tags":["@/kibana-highlighted-field@"],"fields":{"*":{}},"fragment_size":2147483647}},"preference":1676454978239}},"options":{"sessionId":"78f73966-da2b-4240-ab02-c67381163d2a","isRestore":false,"strategy":"ese","isStored":false}}]}'
    };

    fetch('http://dm-kibana-tx.domob.cn/s/staru/internal/bsearch', options)
        .then(response => response.json())
        .then(response => console.log(response))
        .catch(err => console.error(err));
}


</script>


<template>
    <div style="display:flex;  align-items: center; margin-bottom: 5px;">
        <lay-date-picker v-model="range_time" range :placeholder="['开始日期', '结束日期']"
            style="margin-right: 20px;"></lay-date-picker>
        <lay-input v-model="trace_id" placeholder="trace id" style="margin-right: 20px;"></lay-input>
        <lay-button type="primary" class="track_form">追踪</lay-button>
</div>
</template>


<style></style>