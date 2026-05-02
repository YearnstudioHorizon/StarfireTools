<script setup>
import { GetStatus } from "../../wailsjs/go/main/App";
const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  desc: {
    type: String,
    required: false,
    default: "",
  },
  disabled: {
    type: Boolean,
    required: false,
    default: true,
  },
  enabled: {
    type: Boolean,
    required: false,
    default: false,
  },
  group: {
    type: String,
    required: true,
  },
});
GetStatus(props.group, props.title)
  .then((data, err) => {
    console.log('[Debug]组件"' + props.title + '"获取到状态' + data);
  })
  .catch((err) => {
    console.log('[Debug]组件"' + props.title + '"获取状态出错: ' + err);
  });
const emit = defineEmits(["switch", "update:enabled"]);

function switchChoice() {
  emit("switch");
}
</script>

<template>
  <div class="setting-item">
    <div class="setting-content">
      <div class="setting-title">{{ props.title }}</div>
      <div class="setting-desc">
        {{ props.desc }}
      </div>
    </div>
    <div class="setting-action">
      <label class="switch">
        <input
          type="checkbox"
          @change="switchChoice"
          :checked="props.enabled"
        />
        <span class="slider"></span>
      </label>
      <span class="switch-label">{{ props.enabled ? "开" : "关" }}</span>
    </div>
  </div>
</template>

<style lang="css" scoped>
.setting-item {
  --win11-radius: 8px;
  --win11-border-color: rgba(0, 0, 0, 0.08);
  --win11-bg: #ffffff;
  --win11-bg-hover: rgba(255, 255, 255, 0.85);
  --win11-text-primary: #191919;
  --win11-text-secondary: #606060;
  --win11-accent: #0067c0;
  --win11-switch-bg-off: rgba(0, 0, 0, 0.25);
  --win11-switch-thumb-off: #ffffff;
}

/* 适配深色模式 */
@media (prefers-color-scheme: dark) {
  .setting-item {
    --win11-border-color: rgba(255, 255, 255, 0.08);
    --win11-bg: #1c1c1c;
    --win11-bg-hover: rgba(0, 0, 0, 0.85);
    --win11-text-primary: #ffffff;
    --win11-text-secondary: #a6a6a6;
    --win11-accent: #60cdff;
    --win11-switch-bg-off: rgba(255, 255, 255, 0.4);
    --win11-switch-thumb-off: #ffffff;
  }
}

* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-family:
    -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial,
    sans-serif;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  max-width: 100%;
  margin: 8px 0;
  padding: 16px 20px;
  background-color: var(--bg-color);
  border: 1px solid var(--bg-color);
  border-radius: var(--bg-color);
  box-shadow: 0 2px 4px rgb(0, 0, 0);
  transition:
    background-color 0.15s ease,
    border-color 0.15s ease;
}

.setting-item:hover {
  background-color: var(--bg-color);
}

.setting-content {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.setting-title {
  text-align: left;
  font-size: 14px;
  font-weight: 600;
  color: var(--font-color);
}

.setting-desc {
  font-size: 12px;
  color: var(--font-color);
}

.setting-action {
  display: flex;
  align-items: center;
  gap: 12px;
}

/* Switch 样式 */
.switch {
  position: relative;
  display: inline-block;
  width: 40px;
  height: 20px;
  /* border: var(--font-color); */
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
  border: var(--font-color);
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: color-mix(in srgb, var(--font-color), transparent 80%);
  transition: 0.2s cubic-bezier(0.1, 0.9, 0.2, 1);
  border-color: black;
  border-radius: 20px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 12px;
  width: 12px;
  left: 4px;
  bottom: 4px;
  background-color: color-mix(in srgb, var(--bg-color), transparent 80%);
  transition: 0.2s cubic-bezier(0.1, 0.9, 0.2, 1);
  border-radius: 50%;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.15);
}

input:checked + .slider {
  background-color: #0067c0;
}

input:checked + .slider:before {
  transform: translateX(20px);
  background-color: #ffffff;
}

.switch-label {
  font-size: 13px;
  width: 20px;
}
</style>
