<template>
    <ContentBase>
        <PageNotFound />
    </ContentBase>
</template>
    
<script>
import ContentBase from '@/components/ContentBase.vue';
import PageNotFound from '@/components/404NotFound.vue'

export default {
    name: "PageNotFoundView",
    components: {
        ContentBase,
        PageNotFound,
    },
}
</script>
    
<style></style>
    