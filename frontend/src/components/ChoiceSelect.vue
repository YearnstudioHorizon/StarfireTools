<script setup>
import { ref, computed, onMounted, onUnmounted } from "vue";

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
  options: {
    type: Array,
    required: true,
    // [{ label: '选项 A', value: 'a' }]
  },
  modelValue: {
    type: [String, Number],
    required: false,
  },
});

const emit = defineEmits(["update:modelValue", "change"]);

const isOpen = ref(false);
const selectContainer = ref(null);

const currentLabel = computed(() => {
  const opt = props.options.find((o) => o.value === props.modelValue);
  return opt ? opt.label : "";
});

function toggleOpen() {
  isOpen.value = !isOpen.value;
}

function selectOption(value) {
  emit("update:modelValue", value);
  emit("change", value);
  isOpen.value = false;
}

function handleClickOutside(event) {
  if (selectContainer.value && !selectContainer.value.contains(event.target)) {
    isOpen.value = false;
  }
}

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
});
</script>

<template>
  <div class="setting-item">
    <div class="setting-content">
      <div class="setting-title">{{ props.title }}</div>
      <div class="setting-desc" v-if="props.desc">
        {{ props.desc }}
      </div>
    </div>
    <div class="setting-action" ref="selectContainer">
      <div class="win11-custom-select">
        <div
          class="win11-select-trigger"
          :class="{ 'is-open': isOpen }"
          @click="toggleOpen"
        >
          {{ currentLabel }}
        </div>
        <transition name="fade-slide">
          <div v-if="isOpen" class="win11-select-menu">
            <div
              v-for="opt in props.options"
              :key="opt.value"
              class="win11-select-option"
              :class="{ 'is-selected': opt.value === props.modelValue }"
              @click="selectOption(opt.value)"
            >
              {{ opt.label }}
            </div>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<style scoped>
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
  border: 1px solid color-mix(in srgb, var(--font-color), transparent 80%);
  border-radius: var(--font-color);
  box-shadow: 0 2px 4px rgb(0, 0, 0);
  transition:
    background-color 0.15s ease,
    border-color 0.15s ease;
}

/* .setting-item:hover {
  background-color: var(--win11-bg-hover);
} */

.setting-content {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.setting-title {
  text-align: left;
  font-size: 14px;
  font-weight: 600;
}

.setting-desc {
  font-size: 12px;
}

.setting-action {
  display: flex;
  align-items: center;
  gap: 12px;
  /* border: color-mix(in srgb, var(--font-color), transparent 80%); */
  border: color-mix(in srgb, var(--font-color), transparent 80%);
  border-style: solid;
  border-width: 1px;
  border-radius: 8px;
}

/* Windows 11 自定义下拉框样式 */
.win11-custom-select {
  position: relative;
  min-width: 140px;
}

.win11-select-trigger {
  background-color: color-mix(in srgb, var(--bg-color), transparent 60%);
  /* border: 1px solid var(--win11-border-color); */
  border-radius: 4px;
  padding: 6px 32px 6px 12px;
  font-size: 14px;
  color: var(--font-color);
  cursor: pointer;
  outline: none;
  transition: all 0.2s ease;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3Csvg xmlns='http://www.w3.org/2000/svg' width='12' height='12' viewBox='0 0 12 12' fill='none'%3E%3Cpath d='M2.5 4.5L6 8L9.5 4.5' stroke='%23606060' stroke-width='1.2' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 10px center;
  user-select: none;
}

.win11-select-trigger:hover {
  background-color: color-mix(in srgb, var(--bg-color), transparent 90%);
}

/* .win11-select-trigger.is-open {
  border-bottom: 2px solid var(--win11-accent);
  padding-bottom: 5px; 
} */

/* 展开的悬浮菜单 */
.win11-select-menu {
  position: absolute;
  top: 100%;
  right: 0;
  margin-top: 4px;
  width: 100%;
  background-color: var(--bg-color);
  border-radius: 8px; /* 圆角菜单 */
  padding: 4px;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12); /* 明显悬浮阴影 */
  z-index: 1000;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.win11-select-option {
  padding: 6px 12px;
  font-size: 14px;
  color: var(--font-color);
  cursor: pointer;
  border-radius: 4px; /* 选项本身也带有小圆角 */
  transition: background-color 0.15s ease;
  user-select: none;
}

.win11-select-option:hover {
  background-color: rgba(0, 0, 0, 0.04);
}

.win11-select-option.is-selected {
  background-color: rgba(0, 0, 0, 0.03);
  font-weight: 600;
}

/* 下拉动画 */
.fade-slide-enter-active,
.fade-slide-leave-active {
  transition:
    opacity 0.2s ease,
    transform 0.2s ease;
}

.fade-slide-enter-from,
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-5px);
}
</style>
