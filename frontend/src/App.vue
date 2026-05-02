<script setup>
import { RouterView } from "vue-router";
import Titlebar from "./components/Titlebar.vue";
import Sidebar from "./components/Sidebar.vue";

function minWindow() {
  window.runtime.WindowMinimise();
}

function maxWindow() {
  window.runtime.WindowToggleMaximise();
}

function exit() {
  window.runtime.Quit();
}
</script>

<template>
  <div class="app-shell">
    <Titlebar @min="minWindow" @max="maxWindow" @close="exit" />
    <div class="app-body">
      <Sidebar />
      <main class="app-main" data-wails-no-drag>
        <!-- 注意这里提取了 route，并绑定了 key -->
        <RouterView v-slot="{ Component, route }">
          <transition name="slide-up" mode="out-in">
            <component :is="Component" :key="route.fullPath" />
          </transition>
        </RouterView>
      </main>
    </div>
  </div>
</template>

<style scoped>
.app-shell {
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
}

.app-body {
  flex: 1;
  min-height: 0;
  min-width: 0;
  display: flex;
  overflow: hidden;
}

.app-main {
  flex: 1;
  min-width: 0;
  overflow-x: hidden;
  overflow-y: auto;
  padding: 14px 10px 18px;
  box-sizing: border-box;

  /* 隐藏滚动条 */
  -ms-overflow-style: none;
  scrollbar-width: none;
}

.app-main::-webkit-scrollbar {
  width: 0;
  height: 0;
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.06s ease-out;
}
.slide-up-enter-from {
  opacity: 0;
  transform: translateY(20px); /* 新组件从下方 20px 处浮现 */
}
.slide-up-leave-to {
  opacity: 0;
  transform: translateY(-20px); /* 旧组件向上方 20px 处渐隐 */
}
</style>
