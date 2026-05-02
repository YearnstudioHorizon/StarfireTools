<script setup>
import { reactive, ref } from "vue";
import ChoiceGroup from "../components/ChoiceGroup.vue";
import Choice from "../components/Choice.vue";
import { GetItemGroups, Do, Cancel } from "../../wailsjs/go/main/App";
const itemGroups = ref([]);
const items = ref({});
GetItemGroups().then((data) => {
  data.forEach((element, k) => {
    console.log("[Debug]组名: " + element.name);
    itemGroups.value.push(element.name);
    let itemList = [];
    element.items.forEach((v, key) => {
      const temp = {
        name: v,
        desc: element.items_descs[key],
        enabled: false,
        disabled: false,
      };
      console.log("  [Debug]优化项名称: " + v);
      itemList.push(temp);
    });
    items.value[element.name] = itemList;
    // console.log(itemList);
  });
});

const handleSwitch = (key, title) => {
  items.value[title].forEach((v) => {
    if (v.name == key) {
      // 找到了组件对象
      console.log(
        key +
          "进行了状态切换, 来自组" +
          title +
          ", 该组件切换前启用状态: " +
          v.enabled,
      );
      v.enabled = !v.enabled; // 切换启用状态
      v.disabled = true; // 禁用组件 防抖
      if (v.enabled) {
        // 启用
        Do(title, key).then(() => {
          items.value[title].disabled = false;
        });
      } else {
        // 禁用
        Cancel(title, key).then(() => {
          items.value[title].disabled = false;
        });
      }
    }
  });
};
</script>

<template>
  <div>
    <ChoiceGroup v-for="(v, k) in itemGroups" :key="k" :title="v">
      <Choice
        v-for="(val, key) in items[v]"
        :key="key"
        :title="val.name"
        :desc="val.desc"
        :enabled="val.enabled"
        :disabled="val.disabled"
        :group="v"
        @switch="handleSwitch(val.name, v)"
      />
    </ChoiceGroup>
  </div>
</template>
