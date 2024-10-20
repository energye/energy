<template>
  <div class="header">
    <div class="title">ENERGY VUE</div>
    <div class="buttons">
      <button title="最小化" @click="minimize()">最小化</button>
      <button title="最大化/还原" @click="maximize()">最大化/还原</button>
      <button title="关闭" @click="close()">关闭</button>
    </div>
  </div>
  <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo"/>
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo"/>
    </a>
    <HelloWorld msg="Vite + Vue"/>
  </div>
</template>

<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
function minimize() {
  console.log("minimize")
  ipc.emit('minimize')
}

function maximize() {
  console.log("maximize")
  ipc.emit('maximize', function () {
  })
}

function close() {
  ipc.emit('close')
}

ipc.on("testasync", function (complete) {
  console.log("testasync",  complete)
  setTimeout(function (){
    complete.callback("data");
  }, 1000)
}, {mode: 1})

</script>

<style scoped>
.header {
  width: 100%;
  height: 35px;
  -webkit-app-region: drag;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: #f0f0f0;
  user-select: none;

.title {
  margin: 10px;
  font-size: 14px;
  font-weight: bold;
  color: #1f5eff;
}

.buttons button {
  border: none;
  background-color: transparent;
  font-size: 14px;
  cursor: pointer;
}

.buttons button:focus {
  border: 0 none;
  outline: none;
}

.buttons button:hover {
  color: #182a85;
  font-weight: bold;
}

}
.main {
  max-height: calc(100vh - 35);
}

.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
