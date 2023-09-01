<template>
  <div id="container" class="mapContainer w-100" ref="container"></div>
</template>

<script setup lang="ts">
import AMapLoader from '@amap/amap-jsapi-loader'
import { reactive, ref, onMounted, createApp } from 'vue'
import { type Station } from '@/httpclient'
import iconMarker from '@/assets/img/marker.png'
import iconMarkerSelected from '@/assets/img/marker-selected.png'
import iconPredictionMarker from '@/assets/img/prediction_marker.png'
import MapPredictionPopup from '@/components/map/MapPredictionPopup.vue'

const mapCenter = [104, 35]
var map: any | undefined
var aMap: any | undefined
var currentMarkerInfoPopup: any | undefined
var currentPredictionPopup: any | undefined
var markerSelected: any | undefined
var predictionMarker: any | undefined
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
      center: mapCenter,
      pitch: 0,
      viewMode: '2D',
      mapStyle: 'amap://styles/grey',
      clickable: true,
      dragEnable: false,
      zoomEnable: false,
      doubleClickZoom: false,
      keyboardEnable: false
    })
    aMap.plugin('AMap.moveAnimation', function () {})
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
  if (!aMap) return
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

const addStationPopup = (station: Station, position: any) => {
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

const addPredictionPopup = (lng: number, lat: number) => {
  const dom = document.createElement('div')
  const app = createApp({
    components: {
      // eslint-disable-next-line vue/no-unused-components
      MapPredictionPopup
    },
    template: `<MapPredictionPopup :lng="lng" :lat="lat" @on-position-changed="onPositionChanged"/>`,
    data: () => {
      return {
        lng,
        lat
      }
    },
    methods: {
      onPositionChanged(lng: number, lat: number) {
        console.log(`[map] on prediction marker changed lat: ${lat}, lng: ${lng}`)
        if (predictionMarker) {
          const updatedPosition = new aMap.LngLat(lng, lat)
          predictionMarker.setPosition(updatedPosition)
          currentPredictionPopup.open(map, updatedPosition)
        }
      }
    }
  })
  app.mount(dom)

  currentPredictionPopup = new aMap.InfoWindow({
    content: dom,
    anchor: 'bottom-center',
    offset: new aMap.Pixel(0, -60)
  })
  currentPredictionPopup.open(map, new aMap.LngLat(lng, lat))
}

const addStationPrediction = (lng: number, lat: number) => {
  console.log('[map] add prediction to lnglat: ', lng, lat)
  const icon = predictionMarkerIcon()
  predictionMarker = new aMap.Marker({
    position: new aMap.LngLat(lng, lat),
    title: '预测点',
    icon: icon,
    label: predictionMarkerLabel(),
    offset: new aMap.Pixel(0, 0),
    anchor: 'bottom-center',
    animation: 'AMAP_ANIMATION_BOUNCE',
    clickable: true
  })
  map.add(predictionMarker)
}

// handlers

const clickMapHandler = (event: any) => {
  console.log('[map] click map: ', event)
  if (currentMarkerInfoPopup) {
    currentMarkerInfoPopup.close()
  }
  if (markerSelected) {
    deselectMarker(markerSelected)
  }
  if (predictionMarker) {
    removePredictionMarker()
  } else if (event.lnglat.pos) {
    console.log('[map] click lnglat: ', event.lnglat)
    addStationPrediction(event.lnglat.lng, event.lnglat.lat)
    addPredictionPopup(event.lnglat.lng, event.lnglat.lat)
  }
}

const clickMarkerHandler = (event: any) => {
  console.log('[map] click mark: ', event)
  selectMarker(event.target)
}

// utils

const selectMarker = (selectedMarker: any) => {
  removePredictionMarker()
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

const removePredictionMarker = () => {
  if (predictionMarker) {
    map.remove(predictionMarker)
    predictionMarker = undefined
  }
  if (currentPredictionPopup) {
    currentPredictionPopup.close()
  }
  map.setCenter(mapCenter)
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

const predictionMarkerIcon = () => {
  const icon = new aMap.Icon({
    size: new aMap.Size(30, 30),
    image: iconPredictionMarker,
    imageOffset: new aMap.Pixel(0, 0),
    imageSize: new aMap.Size(30, 30)
  })
  return icon
}

const predictionMarkerLabel = () => {
  return {
    content: '预测建站点',
    direction: 'top',
    offset: new aMap.Pixel(0, -5)
  }
}

defineExpose({
  setStations
})
</script>

<style>
.amap-marker-label {
  border: 0;
  font-size: 14px;
  padding: 4px;
  background-color: white;
  border-radius: 8px;
  color: black;
}

.amap-info-content {
  border-radius: 8px;
  color: black;
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

.mapContainer {
  aspect-ratio: 16/9;
}
</style>
