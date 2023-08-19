<template>
    <div id="container" ref="container"></div>
</template>

<script setup lang="ts">

import AMapLoader from '@amap/amap-jsapi-loader'
import { computed, reactive, ref, onMounted } from 'vue'
import httpclient, { type Station } from '@/httpclient'

const requesting = ref(false)
var map: any | null
var aMap: any | null
const container = ref()
const stations = reactive<{
    allStations: Station[]
}>({
    allStations: []
})

onMounted(async () => {
    initMap()
})

const initMap = async () => {
    try {
        aMap = await AMapLoader.load({
            key: "d9997a27b8b492c266a01c2d5e5be64a",
            version: "2.0",
            plugins: [''],
        })
        map = new aMap.Map('container', {
            zoom: 4.1,
            center: [110, 35],
            pitch: 0,
            viewMode: '2D',
            mapStyle: 'amap://styles/grey',
        })
        await loadStations()
    } catch (err: any) {
        console.log('[map] load map with error: ', err.toString())
    }
}

const loadStations = async () => {
    requesting.value = true
    const resp = await httpclient.getStations()
    if (resp?.code == 200) {
        stations.allStations = resp.payload
    } else {
    }
    requesting.value = false
    await addStationMarks()
}

const addStationMarks = async () => {
    stations.allStations.forEach((station: Station) => {
        console.log('[map] add marker for station: ', JSON.stringify(station))
        if (station.name && station.lat && station.lng) {
            const marker = new aMap.Marker({
                position: new aMap.LngLat(station.lng, station.lat),
                title: station.name,
                offset: new aMap.Pixel(-13, -30)
            })
            map.add(marker)
        } else {
            console.log('[map] invalid station: ', station.id)
        }
    })
}



</script>

<style scoped>
#container {
    position: absolute;
    left: 0px;
    right: 0px;
    padding: 0px;
    margin: 0px;
    width: 100%;
    height: calc(100vh);
}
</style>