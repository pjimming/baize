<template>
<ContentBase>
    <h3>golang项目包依赖可视化工具🛠</h3>
    <hr>
    <InputVue />
    <ProjectInfo />
    <PackageList />
</ContentBase>
</template>

<script>
import ContentBase from '@/components/ContentBase.vue';
import InputVue from '@/components/Input.vue';
import ProjectInfo from '@/components/ProjectInfo.vue';
import PackageList from '@/components/PackageList.vue';

export default {
    name: "HomeView",
    components: {
        ContentBase,
        InputVue,
        ProjectInfo,
        PackageList,
    },
}
</script>

<style>

</style>
