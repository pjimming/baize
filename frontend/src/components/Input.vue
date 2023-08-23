<template>
    <div class="input-group input-group-lg">
        <span class="input-group-text" id="inputGroup-sizing-default">请输入项目的绝对路径</span>
        <input type="text" class="form-control" placeholder="例如：/Users/panjiangming/Project/baize/backend"
            v-model="inputValue">
        <button class="btn btn-primary" @click="getModulePath">Submit</button>
    </div>
</template>

<script>
import { fetchModuleInfo } from '@/services/api';
import store from '@/store'

export default {
    name: "InputVue",
    data() {
        return {
            inputValue: "D:\\GoProject\\baize\\backend",
            // inputValue: "/Users/panjiangming/Project/baize/backend",
        };
    },
    methods: {
        async getModulePath() {
            const queryParams = {
                dir: this.inputValue,
            }

            store.commit('setDir', this.inputValue);

            try {
                fetchModuleInfo(queryParams);
            } catch (error) {
                console.error(error.message);
            }
        },
    },
}
</script>

<style scoped></style>
