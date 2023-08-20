<template>
<ContentBase>
    <h3>golang项目包依赖可视化工具</h3>
    <hr>
    <InputVue />
</ContentBase>
</template>

<script>
import ContentBase from '@/components/ContentBase.vue';
import InputVue from '@/components/Input.vue';

export default {
    name: "HomeView",
    components: {
        ContentBase,
        InputVue,
    },
}
</script>

<style>

</style>
