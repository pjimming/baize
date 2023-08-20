<template>
<div class="container">
    <h5>包列表 / Package List</h5>
    <el-table :data="formattedPackageList" stripe style="width: 100%">
        <el-table-column prop="packageName" label="包名 / Package Name" />
    </el-table>
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
        packageList() {
            return this.$store.state.projectInfo.packageList;
        },
        formattedPackageList() {
            return this.packageList.map(packageName => ({
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
