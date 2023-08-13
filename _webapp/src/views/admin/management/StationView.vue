<template>
    <table class="table table-striped table-hover">
        <thead>
            <th scope="col">#</th>
            <th scope="col">Name</th>
            <th scope="col">Latitude</th>
            <th scope="col">Longitude</th>
            <th scope="col">Altitude</th>
        </thead>
        <tbody>
            <tr v-for="station in stations.allStations">
                <th scope="row">{{ station.id }}</th>
                <td>{{ station.name }}</td>
                <td>{{ station.lat }}</td>
                <td>{{ station.lng }}</td>
                <td>{{ station.altitude }}</td>
            </tr>
        </tbody>
    </table>
</template>
  
<script setup lang="ts">
import { RouterLink, useRouter } from 'vue-router'
import httpclient, { type Station } from '@/httpclient'
import { reactive, ref, onMounted } from 'vue'

const stations = reactive<{
    allStations: Station[];
}>({
    allStations: [],
})
const requesting = ref(false)
const router = useRouter()

onMounted(async () => {
    requesting.value = true
    const resp = await httpclient.getStations()
    if (resp?.code == 200) {
        stations.allStations = resp.payload
    } else {

    }
    requesting.value = false
})

</script>
  
<style scoped></style>
  