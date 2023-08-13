<template>
    <div>
        <button @click="handleTemplateDownload">
            Download Template
        </button>
        <PagedTableComponent />
    </div>
</template>

<script setup lang="ts">

import httpclient from '@/httpclient'
import PagedTableComponent from '@/components/PagedTableComponent.vue'

const handleTemplateDownload = async () => {
    const csv = await httpclient.downloadTemplate()
    if (csv) {
        const anchor = document.createElement('a')
        anchor.href = 'data:text/csv;charset=utf-8,' + encodeURIComponent(csv)
        anchor.target = '_blank'
        anchor.download = 'template.csv'
        anchor.click()
    }
}

</script>
