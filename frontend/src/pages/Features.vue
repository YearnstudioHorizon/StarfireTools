<script setup>
import { reactive } from "vue";
import ChoiceGroup from "../components/ChoiceGroup.vue";
import Choice from "../components/Choice.vue";

const settings = reactive({
  perf: {
    enabled: false,
    title: "启用性能调整",
    desc: "将系统视觉效果切换为“调整为最佳性能”。",
  },
  restore: {
    enabled: true,
    title: "禁用系统还原",
    desc: "关闭系统还原策略和配置入口。",
  },
  error: {
    enabled: false,
    title: "禁用错误报告",
    desc: "禁用 Windows Error Reporting。",
  },
});

const handleSwitch = (key) => {
  settings[key].enabled = !settings[key].enabled;
};
</script>

<template>
  <ChoiceGroup title="系统与体验">
    <Choice
      v-for="(val, key) in settings"
      :key="key"
      :title="val.title"
      :desc="val.desc"
      :enabled="val.enabled"
      @switch="handleSwitch(key)"
    />
  </ChoiceGroup>

  <ChoiceGroup title="更多选项">
    <Choice
      v-for="(val, key) in settings"
      :key="key"
      :title="val.title"
      :desc="val.desc"
      :enabled="val.enabled"
      @switch="handleSwitch(key)"
    />
  </ChoiceGroup>
</template>

