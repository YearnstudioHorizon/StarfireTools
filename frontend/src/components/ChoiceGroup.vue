<script setup>
import { ref } from "vue";

const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  expanded: {
    type: Boolean,
    default: true,
  },
});

const isExpanded = ref(props.expanded);
const emit = defineEmits(["item-switch"]);

const toggle = () => {
  isExpanded.value = !isExpanded.value;
};

const handleItemSwitch = (itemId) => {
  emit("item-switch", itemId);
};

const onEnter = (el) => {
  el.style.height = "0px";
  el.style.overflow = "hidden";

  requestAnimationFrame(() => {
    el.style.height = `${el.scrollHeight}px`;
  });
};

const onAfterEnter = (el) => {
  el.style.height = "";
  el.style.overflow = "";
};

const onLeave = (el) => {
  el.style.height = `${el.scrollHeight}px`;
  el.style.overflow = "hidden";

  // force reflow so the browser picks up the starting height
  el.getBoundingClientRect();

  requestAnimationFrame(() => {
    el.style.height = "0px";
  });
};

const onAfterLeave = (el) => {
  el.style.height = "";
  el.style.overflow = "";
};
</script>

<template>
  <div class="setting-group" :class="{ 'is-collapsed': !isExpanded }">
    <div class="group-header" @click="toggle">
      <span class="group-title">{{ props.title }}</span>
      <div class="chevron-icon">
        <svg
          width="12"
          height="12"
          viewBox="0 0 12 12"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M2.5 4.5L6 8L9.5 4.5"
            stroke="currentColor"
            stroke-width="1.2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </div>
    </div>

    <transition
      name="collapse"
      @enter="onEnter"
      @after-enter="onAfterEnter"
      @leave="onLeave"
      @after-leave="onAfterLeave"
    >
      <div v-show="isExpanded" class="group-content">
        <slot></slot>
      </div>
    </transition>
  </div>
</template>

<style scoped>
.setting-group {
  width: 100%;
  max-width: 98%;
  background: var(--bg-color);
  border: 1px solid color-mix(in srgb, var(--font-color), transparent 80%);
  border-radius: 8px;
  margin-left: 10px;
  margin-bottom: 16px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  color: var(--font-color);
}

.setting-group:hover {
  z-index: 5;
}

.setting-group:has(.is-open) {
  z-index: 10;
}

.group-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 20px;
  cursor: pointer;
  user-select: none;
  transition: background 0.1s;
  border-radius: 8px 8px 0 0;
}

.is-collapsed .group-header {
  border-radius: 8px;
}

.group-header:hover {
  background: color-mix(in srgb, var(--bg-color), transparent 80%);
}

.group-title {
  font-size: 14px;
  font-weight: 600;
}

.chevron-icon {
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.is-collapsed .chevron-icon {
  transform: rotate(-90deg);
}

.group-content {
  /* 移除内部的内边距，完全对齐 */
  padding: 0;
  display: flex;
  flex-direction: column;
}

.collapse-enter-active,
.collapse-leave-active {
  transition: height 0.22s cubic-bezier(0.4, 0, 0.2, 1);
  will-change: height;
}

@media (prefers-reduced-motion: reduce) {
  .collapse-enter-active,
  .collapse-leave-active {
    transition-duration: 0ms;
  }
}

:deep(.setting-item) {
  margin: 0 !important;
  box-shadow: none !important;
  background: transparent !important;
  border-radius: 0 !important;
  border-left: none !important;
  border-right: none !important;
  border-bottom: 1px solid var(--group-border);
}

:deep(.setting-item:last-child) {
  border-bottom: none;
  border-radius: 0 0 8px 8px !important;
}

:deep(.setting-item:hover) {
  background: var(--header-hover) !important;
}
</style>
