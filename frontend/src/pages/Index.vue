<script setup>
import InfoCard from "../components/InfoCard.vue";
import {
  GetUserName,
  GetOSName,
  GetKernelVersion,
  GetCPU,
  GetCPUCores,
  GetCPUMHz,
  GetPhyMem,
  GetVirMem,
} from "../../wailsjs/go/main/App";
import { onMounted, ref } from "vue";
import Banner from "../components/Banner.vue";
const username = ref(null);
const os = ref("正在加载");
const version = ref("正在加载");
const cpuname = ref("正在加载");
const cpucore = ref("?");
const cpumhz = ref("?");
const phymem = ref("正在获取");
const virmem = ref("正在获取");
GetUserName().then((data) => {
  username.value = data;
});
GetOSName().then((data) => {
  os.value = data;
  // console("获取到系统: ", os.value);
});
GetKernelVersion().then((data) => {
  version.value = data;
});
GetCPU().then((data) => {
  cpuname.value = data;
});
GetCPUCores().then((data) => {
  cpucore.value = data;
});
GetCPUMHz().then((data) => {
  cpumhz.value = data;
});
GetVirMem().then((data) => {
  virmem.value = Math.trunc(data / 1024 / 1024 / 100) / 10;
});
GetPhyMem().then((data) => {
  phymem.value = Math.trunc(data / 1024 / 1024 / 100) / 10;
});
// 时间问候语
const hours = new Date().getHours();
const nowName = ref("");
if (hours > 6 && hours < 12) {
  nowName.value = "早上好";
} else if (hours >= 12 && hours <= 14) {
  nowName.value = "中午好";
} else if (hours > 14 && hours <= 19) {
  nowName.value = "下午好";
} else if (hours > 19 && hours < 23) {
  nowName.value = "晚上好";
} else {
  nowName.value = "夜深了";
}
// 一言
const content = ref("正在加载一言");
const fromwho = ref("");
function refreshYiyan() {
  fromwho.value = "正在刷新";
  // 请求
  fetch("https://v1.hitokoto.cn/")
    .then((response) => response.json())
    .then((data) => {
      content.value = data.hitokoto;
      console.log(data.from_who, ",,,", data.from);
      if (data.from_who === data.from || data.from_who == null) {
        fromwho.value = data.from;
      } else {
        fromwho.value = data.from_who + ", " + data.from;
      }
    })
    .catch((error) => {
      content.value = "获取失败, 请检查网络";
      fromwho.value = "";
    });
}
onMounted(() => {
  refreshYiyan();
});
function openGithub() {
  window.runtime.BrowserOpenURL(
    "https://github.com/YearnstudioHorizon/StarfireTools",
  );
}
</script>

<template>
  <div class="page-div">
    <h1 style="margin-left: 3%">{{ username + ", " + nowName }}</h1>
    <div
      style="
        width: 95%;
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 15px;
      "
    >
      <Banner
        title="欢迎使用 星火工具箱 Starfire Tools"
        desc="好用的话, 给我们点个Star吧"
        button-text="访问 Github↗"
        @action="openGithub"
      ></Banner>
      <div class="cards">
        <InfoCard
          :title="content"
          :desc="
            fromwho == '' || fromwho == '正在刷新'
              ? fromwho
              : 'From: ' + fromwho
          "
          @click="refreshYiyan"
        ></InfoCard>
        <InfoCard :title="os" :desc="version"></InfoCard>
      </div>
      <div class="cards">
        <InfoCard
          :title="cpuname"
          :desc="cpucore + '核心, ' + cpumhz + 'MHz'"
        ></InfoCard>
        <InfoCard
          :title="'运行内存: ' + phymem + 'GiB'"
          :desc="'虚拟内存: ' + virmem + 'GiB'"
        ></InfoCard>
      </div>
    </div>
  </div>
</template>

<style scoped>
* {
  color: var(--font-color);
  background-color: var(--backgourd-color);
}

.cards {
  display: flex;
  width: 95%;
  gap: 15px;
}
</style>
