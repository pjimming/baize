<template>
<div class="container">
    <h5>项目信息 / Project Info</h5>
    <el-descriptions direction="vertical" :column="3" :size="size" border>
        <el-descriptions-item label="模块路径 / Module Path">{{ projectInfo.modulePath }}</el-descriptions-item>
        <el-descriptions-item label="模块名称 / Module Name">{{ projectInfo.moduleName }}</el-descriptions-item>
        <el-descriptions-item label="包数量 / Packag Count">{{ projectInfo.packageCount }}</el-descriptions-item>
    </el-descriptions>
</div>
</template>

<script>
export default {
    data() {
        return {
            size: 'large',
        };
    },
    computed: {
        projectInfo() {
            return this.$store.state.projectInfo;
        },
    },
}
</script>

<style scoped>
.container {
    margin-top: 20px;
}
</style>
