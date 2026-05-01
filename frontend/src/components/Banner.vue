<script setup>
import { computed } from "vue";

const props = defineProps({
  // 主标题
  title: {
    type: String,
    required: true,
  },
  // 副标题或描述内容
  desc: {
    type: String,
    required: false,
    default: "",
  },
  // 主操作按钮文字，不传则不显示按钮
  buttonText: {
    type: String,
    required: false,
    default: "",
  },
});

const emit = defineEmits(["click", "action"]);

const hasDesc = computed(() => props.desc !== "");
const hasButton = computed(() => props.buttonText !== "");
</script>

<template>
  <div class="win11-card win11-banner" @click="emit('click')">
    <div class="banner-body">
      <div class="banner-title">{{ title }}</div>
      <div class="banner-desc" v-if="hasDesc">
        {{ desc }}
      </div>
    </div>

    <div class="banner-action" v-if="hasButton || $slots.action">
      <slot name="action">
        <button class="banner-btn" @click.stop="emit('action')">
          {{ buttonText }}
        </button>
      </slot>
    </div>
  </div>
</template>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.win11-card {
  background-color: var(--backgourd-color);
  backdrop-filter: blur(16px);
  -webkit-backdrop-filter: blur(16px);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 24px 32px;
  width: 95%;
  box-shadow:
    0 2px 2px color-mix(in srgb, var(--font-color), transparent 80%),
    0 2px 2px color-mix(in srgb, var(--font-color), transparent 70%);
  color: var(--font-color);
  transition:
    transform 0.2s cubic-bezier(0.1, 0.9, 0.2, 1),
    background-color 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.win11-card:hover {
  /* background-color: rgba(45, 45, 45, 0.9); */
  transform: translateY(-1px);
}

/* Banner样式 */
.win11-banner {
  padding: 32px 40px;
}

.banner-body {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-width: 75%;
}

.banner-title {
  font-size: 24px;
  font-weight: 600;
  line-height: 1.4;
  color: var(--font-color);
}

.banner-desc {
  font-size: 14px;
  color: #8a99ad;
  line-height: 1.5;
}

.banner-action {
  display: flex;
  align-items: center;
}

.banner-btn {
  background-color: var(--backgourd-color);
  border: 1px solid var(--font-color);
  color: var(--font-color);
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition:
    background-color 0.2s,
    border-color 0.2s;
}

.banner-btn:hover {
  background-color: var(--backgourd-color);
  color: var(--font-color);
  /* border-color: rgba(255, 255, 255, 0.2); */
}

.banner-btn:active {
  transform: scale(0.98);
}
</style>
