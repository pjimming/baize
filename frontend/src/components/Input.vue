<template>
<div class="input-group input-group-lg">
    <span class="input-group-text" id="inputGroup-sizing-default">请输入项目的绝对路径</span>
    <input type="text" class="form-control" v-model="inputValue" disabled>
    <button class="btn btn-primary" @click="getProjectInfo">Submit</button>
</div>
</template>

<script>
const baizeURL = "http://localhost:8888/baize/v1"

export default {
    name: "InputVue",
    data() {
        return {
            selectType: "text",
            inputValue: "D:\\GoProject\\baize\\backend",
        };
    },
    methods: {
        async getProjectInfo() {
            const queryParams = {
                dir: this.inputValue,
            }
            const queryString = Object.keys(queryParams)
                .map(key => `${encodeURIComponent(key)}=${encodeURIComponent(queryParams[key])}`)
                .join('&');
            const apiURL = baizeURL + "/local/module";

            try {
                const response = await fetch(`${apiURL}?${queryString}`, {
                    method: "GET",
                });

                if (response.ok) {
                    const responseData = await response.json();
                    console.log("GetModuleName Response:", responseData);
                } else {
                    console.error("Request failed with status:", response.status);
                }
            } catch (error) {
                console.error("An error occurred:", error);
            }
        },
    },
}
</script>

<style scoped>

</style>
