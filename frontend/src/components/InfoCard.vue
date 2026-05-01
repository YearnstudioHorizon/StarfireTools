<script setup>
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
});
defineEmits(["click"]);
</script>

<template>
  <div class="win11-card" @click="$emit('click')">
    <div class="card-content">{{ title }}</div>
    <div class="card-footer" v-if="desc != ''">
      <span class="card-title">{{ desc }}</span>
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
  width: 100%;
  flex: 1;
  box-shadow:
    0 2px 4px color-mix(in srgb, var(--font-color), transparent 80%),
    0 2px 4px color-mix(in srgb, var(--font-color), transparent 70%);
  color: var(--font-color);
  transition:
    transform 0.2s cubic-bezier(0.1, 0.9, 0.2, 1),
    background-color 0.2s ease;
}

.win11-card:hover {
  /* background-color: rgba(45, 45, 45, 0.9); */
  transform: translateY(-1px);
}

.card-content {
  font-size: 20px;
  line-height: 1.6;
  margin-bottom: 16px;
  letter-spacing: 0.5px;
  font-weight: 400;
}

.card-footer {
  display: flex;
  align-items: center;
  font-size: 13px;
  color: #a0a0a0;
}

.card-title {
  color: #8a99ad;
}
</style>
