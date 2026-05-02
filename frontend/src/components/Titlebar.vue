<script setup>
import { onMounted } from "vue";
import logo from "../assets/images/appicon.png";
import { ref } from "vue";
import { EventsOn } from "../../wailsjs/runtime/runtime";

defineEmits(["min", "max", "close"]);
const isMaximise = ref(null);
onMounted(async () => {
  isMaximise.value = await window.runtime.WindowIsMaximised();
});
</script>

<template>
  <div
    data-wails-drag
    style="
      height: 60px;
      width: 100%;
      background: var(--bg-color);
      user-select: none;
      flex-shrink: 0;
      color: var(--font-color);
      display: flex;
      align-items: center;
    "
  >
    <img :src="logo" alt="logo" style="height: 55%; margin-left: 20px" />
    <div style="margin-left: 5px; font-size: large">
      星火工具箱&nbsp;Starfire&nbsp;Tools
    </div>
    <div style="margin-left: auto">
      <div class="window-controls">
        <button
          class="window-btn btn-minimize"
          aria-label="最小化"
          @click="$emit('min')"
        >
          <svg viewBox="0 0 12 12">
            <line x1="1" y1="6" x2="11" y2="6" />
          </svg>
        </button>

        <button
          class="window-btn btn-maximize"
          aria-label="最大化"
          @click="$emit('max')"
        >
          <!-- <svg viewBox="0 0 12 12" v-if="isMaximise">
            <path d="M3 1H11V9H9V3H3V1Z" />
            <path d="M1 3H9V11H1V3Z" />
          </svg> -->

          <svg viewBox="0 0 12 12">
            <rect x="1" y="1" width="10" height="10" rx="1" />
          </svg>
        </button>

        <button
          class="window-btn btn-close"
          aria-label="关闭"
          @click="$emit('close')"
        >
          <svg viewBox="0 0 12 12">
            <line x1="1" y1="1" x2="11" y2="11" />
            <line x1="11" y1="1" x2="1" y2="11" />
          </svg>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.window-controls {
  display: flex;
  align-items: center;
  background-color: var(--bg-color);
  padding: 8px 16px;
  height: 40px;
  box-sizing: border-box;
  font-family:
    -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
}

/* 按钮基础样式 */
.window-btn {
  width: 32px;
  height: 32px;
  background: transparent;
  border: none;
  outline: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: 8px;
  transition:
    background-color 0.2s,
    color 0.2s;
  border-radius: 4px;
}

/* 图标基础样式 */
.window-btn svg {
  width: 12px;
  height: 12px;
  stroke: var(--font-color);
  stroke-width: 1.5;
  fill: none;
}

/* 最小化按钮鼠标效果 */
.btn-minimize:hover {
  background-color: color-mix(in srgb, var(--font-color), transparent 80%);
}

/* 最大化按钮悬浮效果 */
.btn-maximize:hover {
  background-color: color-mix(in srgb, var(--font-color), transparent 80%);
}

/* 关闭按钮特殊悬浮效果 */
.btn-close:hover {
  background-color: #e81123;
  color: var(--font-color);
}
.btn-close:hover svg {
  stroke: var(--font-color);
}
</style>
