<template>
    <div class="container">
        <h5>Go 文件列表 / Go File List</h5>
        <el-table :data="pkgList.goFileList" stripe border max-height="280" style="width: 100%">
            <el-table-column prop="name" label="文件名 / File Name" style="width: 50%;" />
            <el-table-column prop="size" label="文件大小 / File Size" />
        </el-table>
    </div>
    <div class="container">
        <h5>包列表 / Package List</h5>
        <h6>工程包列表 / Project Package List</h6>
        <el-table :data="formattedProjectList" stripe border max-height="280" style="width: 100%; margin-bottom: 10px;">
            <el-table-column prop="packageName" label="包名 / Package Name" />
        </el-table>
        <h6>外部包列表 / Other Package List</h6>
        <el-table :data="pkgList.otherPkgList" stripe border max-height="280" style="width: 100%;">
            <el-table-column prop="name" label="包名 / Package Name" />
            <el-table-column prop="version" label="包版本 / Package Version" />
        </el-table>
    </div>
</template>

<script>
export default {
    name: 'ListInfo',
    data() {
        return {
            size: 'large',
        };
    },
    computed: {
        pkgList() {
            return this.$store.state.projectInfo;
        },
        formattedProjectList() {
            return this.pkgList.projectPkgList.map(packageName => ({
                packageName: packageName
            }));
        },
    },
}
</script>

<style scoped>
.container {
    margin-top: 20px;
}
</style>
