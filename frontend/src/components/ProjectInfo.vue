<template>
    <div class="container">
        <h5>项目信息 / Project Info</h5>
        <el-descriptions direction="vertical" :column="3" :size="size" border>
            <el-descriptions-item label="模块名称 / Module Name">{{ projectInfo.moduleName }}</el-descriptions-item>
            <el-descriptions-item label="模块版本 / Module Version">{{ projectInfo.moduleVersion }}</el-descriptions-item>
            <el-descriptions-item label="Go文件数量 / Go File Count">{{ projectInfo.goFileCount }}</el-descriptions-item>
            <el-descriptions-item label="项目包数量 / Project Packages Count">
                {{ projectInfo.projectPkgCount }}
            </el-descriptions-item>
            <el-descriptions-item label="第三方包数量 / Other Packages Count">
                {{ projectInfo.otherPkgCount }}
            </el-descriptions-item>
        </el-descriptions>
    </div>
</template>

<script>
export default {
    name: 'ProjectInfo',
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
