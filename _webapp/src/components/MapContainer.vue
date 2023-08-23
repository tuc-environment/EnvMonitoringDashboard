<template>
  <div id="container" class="container-fluid" ref="container"></div>
</template>

<script setup lang="ts">
import AMapLoader from '@amap/amap-jsapi-loader'
import { computed, reactive, ref, onMounted, nextTick } from 'vue'
import httpclient, { type Station } from '@/httpclient'
import iconMarker from '@/assets/img/marker.png'
import iconMarkerSelected from '@/assets/img/marker-selected.png'

var map: any | undefined
var aMap: any | undefined
var currentMarkerInfoPopup: any | undefined
var markerSelected: any | undefined
const container = ref()
const stations = reactive<{
  allStations: Station[]
}>({
  allStations: []
})
var markers: any[]

const emit = defineEmits<{
  (e: 'didSelectStation', station: Station | undefined): void
}>()

onMounted(async () => {
  initMap()
})

const initMap = async () => {
  try {
    aMap = await AMapLoader.load({
      key: 'd9997a27b8b492c266a01c2d5e5be64a',
      version: '2.0',
      plugins: ['']
    })
    map = new aMap.Map('container', {
      zoom: 4.1,
      center: [104, 35],
      pitch: 0,
      viewMode: '2D',
      mapStyle: 'amap://styles/grey',
      clickable: true,
      dragEnable: false,
      zoomEnable: false,
      doubleClickZoom: false,
      keyboardEnable: false
    })
    map.on('click', clickMapHandler)
    await addStationMarks()
  } catch (err: any) {
    console.log('[map] load map with error: ', err.toString())
  }
}

const setStations = async (updatedStations: Station[]) => {
  stations.allStations = updatedStations
  await addStationMarks()
}

// add map elements

const addStationMarks = async () => {
  const icon = markerIcon()
  markers = stations.allStations.flatMap((station: Station) => {
    console.log('[map] add marker for station: ', JSON.stringify(station))
    if (station.name && station.lat && station.lng) {
      const marker = new aMap.Marker({
        extData: station,
        position: new aMap.LngLat(station.lng, station.lat),
        title: station.name,
        icon: icon,
        label: markerLabel(station.name),
        offset: new aMap.Pixel(0, 0),
        anchor: 'bottom-center',
        animation: 'AMAP_ANIMATION_BOUNCE',
        clickable: true
      })
      marker.on('click', clickMarkerHandler)
      return marker
    } else {
      console.log('[map] invalid station: ', station.id)
      return null
    }
  })
  map.add(markers)
}

const addStationPopup = (station: Station, position: number[]) => {
  const dom: string = `
        <div class="markerPopupContainer">
            <div class="popLabelRow"> 
                <b>Name</b> 
                <span class="popLabelValue">${station.name}</span>
            </div>
            <div class="popLabelRow"> 
                <b>Latitude</b> 
                <span class="popLabelValue">${station.lat}</span>
            </div>
            <div class="popLabelRow"> 
                <b>Longitude</b> 
                <span class="popLabelValue">${station.lng}</span>
            </div>
            <div class="popLabelRow"> 
                <b>Altitude</b> 
                <span class="popLabelValue">${station.altitude}</span>
            </div>
        </div>
    `

  // 创建 infoWindow 实例
  currentMarkerInfoPopup = new aMap.InfoWindow({
    content: dom,
    anchor: 'bottom-center',
    offset: new aMap.Pixel(0, -60)
  })
  currentMarkerInfoPopup.open(map, position)
}

// handlers

const clickMapHandler = (event: any) => {
  console.log('[map] click map: ', event)
  if (currentMarkerInfoPopup) {
    currentMarkerInfoPopup.close()
  }
  deselectMarker(markerSelected)
}

const clickMarkerHandler = (event: any) => {
  console.log('[map] click mark: ', event)
  selectMarker(event.target)
}

// utils

const selectMarker = (selectedMarker: any) => {
  markerSelected = selectedMarker
  const icon = markerIcon()
  const selectedIcon = markerIconSelected()
  markers.forEach((marker) => {
    if (marker == selectedMarker) {
      marker.setIcon(selectedIcon)
    } else {
      marker.setIcon(icon)
      marker.setLabel(markerLabel(marker.getExtData().name))
    }
  })
  const station: Station = selectedMarker.getExtData()
  console.log('[map] selectedStation: ', JSON.stringify(station))
  addStationPopup(station, selectedMarker.getPosition())
  emit('didSelectStation', station)
}

const deselectMarker = (deselectedMarker: any) => {
  const icon = markerIcon()
  deselectedMarker.setIcon(icon)
  deselectedMarker.setLabel(markerLabel(deselectedMarker.getExtData().name))
  emit('didSelectStation', undefined)
}

const markerLabel = (title: string) => {
  return {
    content: title,
    direction: 'top',
    offset: new aMap.Pixel(0, -5)
  }
}

const markerIcon = () => {
  const icon = new aMap.Icon({
    size: new aMap.Size(20, 20),
    image: iconMarker,
    imageOffset: new aMap.Pixel(0, 0),
    imageSize: new aMap.Size(20, 20)
  })
  return icon
}

const markerIconSelected = () => {
  const icon = new aMap.Icon({
    size: new aMap.Size(30, 30),
    image: iconMarkerSelected,
    imageOffset: new aMap.Pixel(0, 0),
    imageSize: new aMap.Size(30, 30)
  })
  return icon
}

defineExpose({
  setStations
})
</script>

<style scoped>
#container {
  left: 0px;
  right: 0px;
  padding: 0px;
  margin: 0px;
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

.amap-info-content {
  border-radius: 8px;
}

.markerPopupContainer {
  font-size: 12px;
  padding-right: 8px;
  padding-left: 8px;
}

.popLabelValue {
  margin-left: 8px;
  text-align: right;
  flex: 1;
}

.popLabelRow {
  display: flex;
}
</style>
