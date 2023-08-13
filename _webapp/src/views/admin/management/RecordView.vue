<template>
    <div>
        <div class="input-group">
            <button class="btn btn-outline-secondary" type="button" id="inputGroupFileAddon04"
                @click="handleTemplateDownload">
                Download Template
            </button>
            <input type="file" class="form-control" id="inputGroupFile04" aria-describedby="inputGroupFileAddon04"
                aria-label="Upload" @change="selectCSVFile">
            <button class="btn btn-outline-primary" type="button"
                id="inputGroupFileAddon04" :disabled="disableSubmitButton" @click="uploadCSVFile">
                Submit
            </button>
        </div>

        <PagedTableComponent />
    </div>
</template>

<script setup lang="ts">

import httpclient from '@/httpclient'
import PagedTableComponent from '@/components/PagedTableComponent.vue'
import { computed, ref } from 'vue';

const uploading = ref(false)
const file = ref<string | null>(null)
const disableSubmitButton = computed((): boolean => {
    return file.value == null
})

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

const selectCSVFile = (event: any) => {
    file.value = event.target.files[0]
};

const uploadCSVFile = async () => {
    if (file.value) {
        const formData = new FormData()
        formData.append('file', file.value)
        try {
            const response = await httpclient.uploadCSVRecords(formData)
            console.log(response?.payload);
        } catch (error) {
            console.error(error);
        }
    }
}

</script>
