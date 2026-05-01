<script setup>
import { onMounted, ref } from "vue";
import ChoiceGroup from "../components/ChoiceGroup.vue";
import Choice from "../components/Choice.vue";
import ChoiceSelect from "../components/ChoiceSelect.vue";

const isDarkMode = ref(false);

onMounted(() => {
  // 判断是否为暗黑模式
  const rootStyle = getComputedStyle(document.documentElement);
  const mainColor = rootStyle.getPropertyValue("--bg-color").trim();
  if (mainColor == "#202020") {
    isDarkMode.value = true;
  } else {
    isDarkMode.value = false;
  }
});

const currentLanguage = ref("zh-CN");
const languageOptions = [
  { label: "简体中文", value: "zh-CN" },
  { label: "English", value: "en-US" },
];

const handleThemeSwitch = () => {
  // console.log("[Debug]点击了暗黑模式切换");
  const root = document.documentElement;
  if (isDarkMode.value) {
    // 切换到亮色模式
    console.log("[Debug]正在切换到亮色模式");
    root.style.setProperty("--group-bg", "#f9f9f9");
    root.style.setProperty("--bg-color", "#f9f9f9");
    root.style.setProperty("--font-color", "#202020");
    window.runtime.WindowSetLightTheme();
  } else {
    // 切换到暗色模式
    console.log("[Debug]正在切换到暗色模式");
    root.style.setProperty("--group-bg", "#202020");
    root.style.setProperty("--bg-color", "#202020");
    root.style.setProperty("--font-color", "#f9f9f9");
    window.runtime.WindowSetDarkTheme();
  }
  isDarkMode.value = !isDarkMode.value;
};
</script>

<template>
  <div class="page-div">
    <div class="header-section">
      <h1 class="page-title">设置</h1>
      <p class="page-subtitle">配置应用程序的显示与默认行为</p>
    </div>

    <ChoiceGroup title="个性化与外观">
      <Choice
        title="深色模式"
        desc="开启后应用将切换为暗色主题风格。"
        :enabled="isDarkMode"
        @switch="handleThemeSwitch"
      />

      <!-- 演示你刚刚需要的新组件 -->
      <ChoiceSelect
        title="显示语言"
        desc="更改应用程序界面所使用的语言。"
        :options="languageOptions"
        v-model="currentLanguage"
      />
    </ChoiceGroup>
  </div>
</template>

<style scoped>
.page-div {
  margin: 20px 0;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.header-section {
  width: 95%;
  margin-bottom: 24px;
}

.page-title {
  color: var(--font-color);
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 6px;
}

.page-subtitle {
  font-size: 14px;
  color: #606060;
}

/* 适配深色模式标题栏 */
@media (prefers-color-scheme: dark) {
  .page-subtitle {
    color: #a6a6a6;
  }
}
</style>
