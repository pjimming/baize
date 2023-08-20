<template>
<div class="input-group input-group-lg">
    <span class="input-group-text" id="inputGroup-sizing-default">请输入项目的绝对路径</span>
    <input type="text" class="form-control" v-model="inputValue" disabled>
    <button class="btn btn-primary" @click="getModulePath">Submit</button>
</div>
</template>

<script>
import { fetchModulePath, fetchProjectInfo } from '@/services/api';
import store from '@/store'

export default {
    name: "InputVue",
    data() {
        return {
            inputValue: "D:\\GoProject\\baize\\backend",
        };
    },
    methods: {
        async getModulePath() {
            console.log(this.$store.state);
            const queryParams = {
                dir: this.inputValue,
            }

            store.commit('setDir', this.inputValue);
            
            try {
                const modulePath = await fetchModulePath(queryParams);
                console.log("GetModulePath Response:", modulePath);

                await fetchProjectInfo({
                    dir: this.inputValue,
                    modulePath: this.$store.state.modulePath,
                })
            } catch (error) {
                console.error(error.message);
            }
        },
    },
}
</script>

<style scoped>

</style>
