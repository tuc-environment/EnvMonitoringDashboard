<template>
  <div id="container" ref="container"></div>
</template>

<script setup lang="ts">
import AMapLoader from '@amap/amap-jsapi-loader'
import { computed, reactive, ref, onMounted, nextTick } from 'vue'
import httpclient, { type Station } from '@/httpclient'
import iconMarker from "@/assets/img/marker.png"
import iconMarkerSelected from "@/assets/img/marker-selected.png"

const requesting = ref(false)
var map: any | null
var aMap: any | null
const container = ref()
const stations = reactive<{
  allStations: Station[]
}>({
  allStations: []
})
var markers: any[]
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
            clickable: true,
            dragEnable: false,
            zoomEnable: false,
            doubleClickZoom: false,
            keyboardEnable: false,
        })
        await loadStations()
        map.on('click', clickMapHandler);
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
    const icon = markerIcon();
    markers = stations.allStations.flatMap((station: Station) => {
        console.log('[map] add marker for station: ', JSON.stringify(station))
        if (station.name && station.lat && station.lng) {
            const marker = new aMap.Marker({
                extData: station,
                position: new aMap.LngLat(station.lng, station.lat),
                title: station.name,
                icon: icon,
                label: {
                    content: station.name,
                    direction: 'top',
                    offset: new aMap.Pixel(0, -5)
                },
                offset: new aMap.Pixel(0, 0),
                anchor: 'bottom-center',
                animation: 'AMAP_ANIMATION_BOUNCE',
                clickable: true,
            })
            marker.on('click', clickMarkerHandler)
            return marker;
        } else {
            console.log('[map] invalid station: ', station.id)
            return null
        }
    })
    map.add(markers)
}

const clickMapHandler = (event: any) => {
    console.log('[map] click map: ', event)
};

const clickMarkerHandler = (event: any) => {
    const icon = markerIcon();
    const selectedIcon = markerIconSelected()
    markers.forEach((marker) => {
        if (marker == event.target) {
            marker.setIcon(selectedIcon)
        } else {
            marker.setIcon(icon)
        }
    })
    console.log('[map] click mark: ', event)
    const station: Station = event.target.getExtData()
    console.log('[map] selectedStation: ', JSON.stringify(station))
}

const markerIcon = () => {
    const icon = new aMap.Icon({
        size: new aMap.Size(20, 20),
        image: iconMarker,
        imageOffset: new aMap.Pixel(0, 0),
        imageSize: new aMap.Size(20, 20),
    })
    return icon
}

const markerIconSelected = () => {
    const icon = new aMap.Icon({
        size: new aMap.Size(30, 30),
        image: iconMarkerSelected,
        imageOffset: new aMap.Pixel(0, 0),
        imageSize: new aMap.Size(30, 30),
    })
    return icon
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

<style>
.amap-marker-label {
    border: 0;
    font-size: 14px;
    padding: 4px;
    background-color: silver;
    border-radius: 8px;
}
</style>

