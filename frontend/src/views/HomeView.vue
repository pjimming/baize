<template>
    <ContentBase>
        <h3>🛠golang项目包依赖可视化工具😼</h3>
        <hr>
        <InputVue />
        <ProjectInfo />
        <ListInfo />
        <BackToTop />
    </ContentBase>
</template>

<script>
import ContentBase from '@/components/ContentBase.vue';
import InputVue from '@/components/Input.vue';
import ProjectInfo from '@/components/ProjectInfo.vue';
import ListInfo from '@/components/ListInfo.vue';
import BackToTop from '@/components/BackToTop.vue'

export default {
    name: "HomeView",
    components: {
        ContentBase,
        InputVue,
        ProjectInfo,
        ListInfo,
        BackToTop,
    },
}
</script>

<style></style>
